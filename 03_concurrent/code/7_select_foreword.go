package main

import (
	"fmt"
	"time"
)

// select语句，其功能和linux中的selec、poll、epoll大体是差不多的
// select作用于多个channel，其会在多个channel中选择一个目前已经就绪的channel
func main() {
	// 案例：现在有两个goroutine都在执行，但是在主goroutine中，当某一个执行完成后，这个时候我会立马知道
	go g1()
	go g2()

	// 阻塞。只要任何一个goroutine执行成功，主线程就会从done中获取到结果，从而执行后面的代码
	// ？？？如果每个goroutine都有自己的channel，即监听多个channel的完成情况，此时这种阻塞的方式就不太好了
	<-done
	fmt.Println("done")
}

// 空结构体能节省空间（和使用bool是同一个效果）
var done = make(chan struct{}) // channel是多线程安全的

func g1() {
	time.Sleep(1 * time.Second)

	done <- struct{}{}
}

func g2() {
	time.Sleep(2 * time.Second)

	done <- struct{}{}
}
