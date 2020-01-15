package main

import (
	"fmt"
)

func main() {
	// if num := 9; num < 0 {
	// 	fmt.Println(num, "is negative")
	// } else if num < 10 {
	// 	fmt.Println(num, "has 1 digit")
	// } else {
	// 	fmt.Println(num, "has multiple digits")
	// }

	// fmt.Printf(stringutil.Reverse("!oG ,olleH"))
	// a := 10
	// switch {
	// case a > 0:
	// 	fmt.Println("a > 0")
	// case a > 5:
	// 	fmt.Println("a > 5")
	// default:
	// 	fmt.Println("default")
	// }

	// x := new([1] int)
	// y := new([1] int)
	// fmt.Println(x)
	// fmt.Println(y)
	// fmt.Println(x == y)

	// x := "hello"
	// for _, x := range x {
	// 	x := x + 'A' - 'a'
	// 	fmt.Printf("%c", x)
	// }
	// fmt.Printf(x)

	// x := "hello"
	// fmt.Printf("%+v", x)
	x := new([10] int)
	y := new([20] int)
	fmt.Println(&x == &y)
}