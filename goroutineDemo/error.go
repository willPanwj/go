package main

import (
	"fmt"
	"time"
)

func addele(a []int, i int) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("add ele fail")
		}
	}()
	a[i] = i
	fmt.Println(a)
}

func main() {
	Array := make([]int, 4)
	for i := 0; i < 10; i++ {
		go addele(Array, i)
	}
	time.Sleep(time.Second * 2)
}