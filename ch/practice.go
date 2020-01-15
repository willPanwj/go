package main

import (
	"fmt"
	"os"
)

func main() {
	// TODO: 性能评估测试
	fmt.Println("commandName:", os.Args[0])
	for i := 1; i < len(os.Args); i++ {
		fmt.Println("index:", i)
		fmt.Println("param:", os.Args[i])
	}
}