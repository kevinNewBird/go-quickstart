package main

import (
	"fmt"
	"time"
)

func main() {
	// 1.无缓冲channel测试demo
	channelCacheTest()

	// 2.无缓冲channel测试demo
	channelNoCacheTest()

	// 3.channel的forrange和close机制
	channelLoopAndClose()

	// 4.单向channel
	channelSimplex()

	// 5.单向channel的应用：生产者+消费者
	channelProducerAndConsumer()
}

// 1.有缓冲channel测试demo
func channelCacheTest() {
	// 1.channel的初始化, 初始化大小如果为0在放值的时候会阻塞，从而报错deadlock
	var message chan string = make(chan string, 1)

	// 2.放值到channel管道中
	message <- "hello"

	// 3.从channel管道中取值赋给data变量
	data := <-message

	fmt.Println(data)
}

// 2.无缓冲channel测试demo
func channelNoCacheTest() {
	// 1.channel的初始化, 初始化大小如果为0在放值的时候会阻塞，从而报错deadlock
	var message chan string = make(chan string, 0)

	go func(message chan string) {
		// 2.从channel管道中取值赋给data变量
		data := <-message
		fmt.Println(data)
	}(message)

	// 3.放值到channel管道中
	message <- "hello nocache"

	time.Sleep(time.Second * 2)
}

// 3.channel的forrange和close机制
func channelLoopAndClose() {
	var message chan int = make(chan int, 2)

	go func(message chan int) {
		// 3.1.遍历获取channel缓存的值
		for data := range message {
			fmt.Println(data)
		}

		// 3.2.退出channel，打印相关信息
		fmt.Println("all done")
	}(message)

	// 3.3.往channel中放值
	message <- 1
	message <- 2

	// 3.4.通知channel退出
	close(message)

	// 注意：已经关闭的channel不能再次放值，但是可以取值
	//message <- 3 // 关闭后，不允许放值
	data := <-message
	fmt.Println(data)

	time.Sleep(time.Second * 2)
}

// 4.单向通道
func channelSimplex() {
	ch := make(chan int, 3)
	// ch虽然是双向的，但是send只能是写入数据（send-only）
	var send chan<- int = ch
	// recv-only： 只能读取数据
	var recv <-chan int = ch

	send <- 999
	data := <-recv
	fmt.Println(data)
}

func channelProducerAndConsumer() {
	c := make(chan int, 3)
	// c虽然是双向的，参数传递时会自动转换
	go producer(c)
	go consumer(c)
	time.Sleep(time.Second * 5)
}

func producer(out chan<- int) {
	for i := 0; i < 10; i++ {
		out <- i * i
	}
	close(out)
}

func consumer(in <-chan int) {
	for data := range in {
		fmt.Println("num:", data)
	}
}
