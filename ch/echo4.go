package main

import (
	// "flag"
	"fmt"
	// "strings"
	// "time"
)

// var n = flag.Bool("n", falsd, "omit trailing newline")
// var sep = flag.String("s", " ", "separator")

func main() {
	// flag.Parse()
	// fmt.Print(strings.Join(flag.Args(), *sep))
	// if !*n {
	// 	fmt.Println()
	// }
	// seconds := 1
	// micSeconds := time.Duration(seconds) * 1000
	// fmt.Print(micSeconds.Seconds())
	attended := map[string]int{
		"a": 1,
	}
	a, ok := attended["a"]
	if ok {
		fmt.Print(a)
	} else {
		fmt.Printf("b not existed.")
	}
}
