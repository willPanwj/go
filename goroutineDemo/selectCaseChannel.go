package main

import (
  "fmt"
)

func send(c chan int)  {
  for i := 1; i < 10 ; i++  {
    c <- i
    fmt.Println("send data : ", i)
  }
}

func main() {
  resch := make(chan int, 20)
  strch := make(chan string, 10)
  go send(resch)
  strch <- "wd"
  select {
  case a := <- resch:
    fmt.Println("get data resch : ", a)
  case b := <- strch:
    fmt.Println("get data resch : ", b)
  default:
    fmt.Println("no channel actvie")
  }
}