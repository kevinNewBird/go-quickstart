package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// 1.协程的使用，关键字go
	// 主死随从
	go asyncPrint()

	time.Sleep(time.Second)

	// 2.waitGroup机制
	waitGroupDemo1()
	waitGroupDemo2()
}

func asyncPrint() {
	fmt.Println("async print")
}

func waitGroupDemo1() {
	// 2.waitGroup机制
	var wg sync.WaitGroup
	// 两种方式：1.一次性指定；2.递增的方式
	// 2.1.方式1: 一次性指定监控100个goroutine执行结束
	wg.Add(100)
	for i := 0; i < 100; i++ {
		// 2.2.方式2：递增的方式，每一个子goroutine+1
		//wg.Add(1)

		go func(i int) {
			fmt.Println(i, "demo1")
			wg.Done() // 必须要有，表明当前子goroutine执行结束
		}(i)
	}

	// 等到所有子goroutine执行结束
	wg.Wait()
	fmt.Println("demo1 all done")
}

func waitGroupDemo2() {
	// 2.waitGroup机制
	var wg sync.WaitGroup
	// 两种方式：1.一次性指定；2.递增的方式
	// 2.1.方式1: 一次性指定监控100个goroutine执行结束
	//wg.Add(100)
	for i := 0; i < 100; i++ {
		// 2.2.方式2：递增的方式，每一个子goroutine+1
		wg.Add(1)

		go func(i int) {
			defer wg.Done() // defer的常用方式
			fmt.Println(i, "-demo2")
		}(i)
	}

	// 等到所有子goroutine执行结束
	wg.Wait()
	fmt.Println("demo2 all done")
}
