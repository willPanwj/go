package main

import "fmt"

func test(c chan int) {
  for i := 0; i < 10; i++ {
    fmt.Println("send ", i)
    c <- i
  }
}
func main() {
  ch := make(chan int, 10)
  go test(ch)
  for j := 0; j < 10; j++ {
    fmt.Println("get ", <- ch)
  }
}
