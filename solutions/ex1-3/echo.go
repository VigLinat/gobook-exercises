package echo

import (
	"strings"
	"os"
)

func Slow(args []string) string {
	s, sep := "", ""
	for _, arg := range os.Args {
		s += sep + arg
		sep = " "
	}
	return s
}

func Fast(args []string) string {
	return strings.Join(os.Args, " ")
}
