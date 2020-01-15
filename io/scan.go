package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	fmt.Printf("Please type in something:\n")

	for input.Scan() {
		line := input.Text()

		if line == "bye" {
			break
		}

		counts[line]++
	}

	for line, n := range counts {
		fmt.Printf("%d : %s\n", n, line)
	}
}
