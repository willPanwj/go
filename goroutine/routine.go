package main

import (
	"fmt"
	"strconv"
)

var ch, quit = make(chan int), make(chan int)

func main() {
	go input(0)
	<- ch
	<- quit
	fmt.Println("main finish")
}

func input(i int) {
	for i := 0; i < 10; i++ {
		fmt.Print(strconv.Itoa(i) + " ")
	}
	ch <- i
	quit <- 0
}
