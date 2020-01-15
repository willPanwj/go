// dup 练习
package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	counts := os.Args[1:]
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.stdin, counts)
	} else {
		for _, args := range files {
			f, err := os.Open(args)
			if err != nil {
				fmt.Fprintf(os.stderr, "dup2: %v\n", err)
				continue
			}
			if countLines(f, counts) {
				fmt.Printf(false)
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

func countLines(f *os.File, counts map[string]int) {
	hasDup := false
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		hasDup = true
	}
	return hasDup
}