package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "web" {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "image/svg+xml")
			if err := r.ParseForm(); err != nil {
				fmt.Fprintf(os.Stderr, "Error while parsing form: %v", err)
				writeSvg(w, map[string]string{})
				return
			}
			params := make(map[string]string)
			for k, v := range r.Form {
				params[k] = v[0]
			}
			writeSvg(w, params)
		})
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
	writeSvg(os.Stdout, map[string]string{})
}

func writeSvg(out io.Writer, params map[string]string) {
	height := 320
	width := 600
	stroke := "grey"
	fill := "white"
	for k, v := range params {
		switch k {
		case "height":
			if ivalue, err := strconv.Atoi(v); err == nil {
				height = ivalue
			}
		case "width":
			if ivalue, err := strconv.Atoi(v); err == nil {
				width = ivalue
			}
		case "stroke":
			stroke = v
		case "fill":
			fill = v
		}
	}
	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: %s; fill: %s; stroke-width: 0.7' "+
		"width='%d' height='%d'>", stroke, fill, width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			fmt.Fprintf(out, "<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintln(out, "</svg>")

}

func corner(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}
