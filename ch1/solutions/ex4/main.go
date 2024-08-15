package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			if countDups := countLines(f, counts); countDups > 1 {
				fmt.Printf("%s contains %d duplicates\n", arg, countDups)
			}
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

// countLines fills the counts map of duplicate lines
// and returns count of duplicates, if any
func countLines(f *os.File, counts map[string]int) int {
	countDups := 1
	input := bufio.NewScanner(f)
	for input.Scan() {
		text := input.Text()
		if counts[text]++; counts[text] > 1 {
			countDups++
		}
	}
	// NOTE: ignoring potential errors from input.Err()
	return countDups
}
