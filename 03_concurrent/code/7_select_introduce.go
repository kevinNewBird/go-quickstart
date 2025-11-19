package main

import (
	"fmt"
	"time"
)

func main() {
	// 案例：监听多个channel
	g1Chan := make(chan struct{})
	g2Chan := make(chan struct{})
	go g11(g1Chan)
	go g22(g2Chan)

	// 1.某一个分支就绪了就执行分支 2.如果两个都就绪了，先执行哪个--随机的：目的是为了防止饥饿
	// 应用场景，处理timeout的情况
	timer := time.NewTimer(time.Second)
	for {
		select {
		case <-g1Chan:
			fmt.Println("g1 done")
			return
		case <-g2Chan:
			fmt.Println("g2 done")
			return
		case <-timer.C:
			fmt.Println("timeout")
			return
			//default: // 这种方式不优雅，比较好的方式是使用timer
			//	time.Sleep(1 * time.Second)
			//	fmt.Println("timeout")
			//	return
		}

	}
}

func g11(ch chan struct{}) {
	time.Sleep(1 * time.Second)

	ch <- struct{}{}
}

func g22(ch chan struct{}) {
	time.Sleep(1 * time.Second)

	ch <- struct{}{}
}
