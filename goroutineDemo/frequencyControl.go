package main

import (
  "time"
  "fmt"
)

func main(){
  requests := make(chan int, 5)
  for i := 1; i < 5; i++ {
    requests <- i
  }
  close(requests)
  limiter := time.Tick(time.Second * 1)
  for req := range requests {
    <- limiter
    fmt.Println("requets", req, time.Now()) // 执行到这里，需要隔 1 秒才继续往下执行，time.Tick(timer) 上面已定义
  }
}