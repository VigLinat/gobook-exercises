package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	sol1() // discard result
	fmt.Println(sol2())
}

// v1
func sol1() string {
	s, sep := "", ""
	for _, arg := range os.Args[0:] {
		s += sep + arg
		sep = " "
	}
	return s
}

// v2
func sol2() string {
	return strings.Join(os.Args[0:], " ")
}
