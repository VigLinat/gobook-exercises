// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"math/cmplx"
	"os"
	"strings"
)

const (
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
	width, height          = 1024, 1024
)

func main() {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	if len(os.Args) > 1 && strings.HasSuffix(os.Args[1], ".png") {
		w, err := os.Create(os.Args[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
			return
		}
		png.Encode(w, img)
		defer w.Close()
	} else {
		png.Encode(os.Stdout, img) // NOTE: ignoring errors
	}
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			red := uint8(math.Abs(real(v)) / width * 255)
			green := uint8(math.Abs(imag(v)) / height * 255)
			// blue := 255 - contrast*n
			return color.RGBA{red, green, 0, 255}
			// return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
