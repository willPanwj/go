package main

import (
  "fmt"
)

// var readOnlyChan <- chan int   	只读 channel
// var writeOnlyChan chan <- int 		只写 channel
// var mychan  chan int         		读写 channel

func main() {
  mychannel := make(chan int, 10)
  for i := 0; i < 10; i++ {
    mychannel <- i
  }
  close(mychannel)  // 关闭管道
  fmt.Println("data lenght: ", len(mychannel))
  for v := range mychannel {  // 循环管道
    fmt.Println(v)
  }
  fmt.Printf("data lenght:  %d", len(mychannel))
}