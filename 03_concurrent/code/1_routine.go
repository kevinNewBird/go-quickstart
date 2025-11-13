package main

import (
	"fmt"
	"time"
)

func main() {
	// 1.协程的使用，关键字go
	// 主死随从
	go asyncPrint()

	time.Sleep(time.Second)
}

func asyncPrint() {
	fmt.Println("async print")
}
