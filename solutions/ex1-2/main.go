package main

import (
	"fmt"
	"os"
)

func main() {
	s := ""
	for i, arg := range os.Args {
		s += fmt.Sprintf("%d: %s\n", i, arg)
	}
	fmt.Println(s)
}
