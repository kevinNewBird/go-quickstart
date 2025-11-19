package main

import (
	"fmt"
	"sync"
	"time"
)

// 引入context的概念的案例
func main() {
	// 案例：有一个goroutine监控cpu的信息
	// 1.实现方式1
	//demoTest1()

	// 2.实现方式2
	demoTest2()
}

func demoTest1() {
	wg1.Add(1)
	go cpuInfo()
	time.Sleep(time.Second * 6)
	stop = true
	wg1.Wait()
	fmt.Println("监控完成 ")
}

func demoTest2() {
	// 如果我们想要往goroutine中传递一些信息，可以通过channel传递（也就是参数的方式）
	// 而go语言给我们提供了一种更加优雅的方式，即context
	stopCh := make(chan struct{})
	wg2.Add(1)
	go cpuInfo2(stopCh)
	time.Sleep(time.Second * 6)
	stopCh <- struct{}{}
	wg2.Wait()
	fmt.Println("监控完成 ")
}

var wg1 sync.WaitGroup
var stop bool

// 方式1:不优雅
func cpuInfo() {
	// 主动退出监控程序
	defer wg1.Done()
	for {
		if stop {
			break
		}
		// 每2秒钟获取cpu的信息
		time.Sleep(1 * time.Second)
		fmt.Println("cpu的信息")
	}
}

// 方式2: 使用select
var wg2 sync.WaitGroup

func cpuInfo2(stopCh chan struct{}) {
	defer wg2.Done()
	for {
		select {
		case <-stopCh:
			fmt.Println("退出cpu监控")
			return
		default: // 不阻塞走的方法
			// 每2秒钟获取cpu的信息
			time.Sleep(1 * time.Second)
			fmt.Println("cpu的信息")
		}
	}
}
