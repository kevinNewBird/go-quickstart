package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	// 如果我们想要往goroutine中传递一些信息，可以通过channel传递（也就是参数的方式）
	// 而go语言给我们提供了一种更加优雅的方式，即context
	// context包提供了三种函数：WithCancel， WithTimeout， WithValue
	// 如果你的goroutine，函数中希望被控制：超时、传值，但是我不希望影响我原来的接口信息的时候，函数参数重第一个参数就尽量要加上一个ctx
	wg5.Add(1)
	// timeout 主动超时
	ctx1, _ := context.WithTimeout(context.Background(), 6*time.Second) // 父
	// value 传值
	ctx2 := context.WithValue(ctx1, "traceid", "x1sfs") // 子
	// deadline 在时间点超时
	go cpuInfo5(ctx2) // 传递子
	wg5.Wait()
	fmt.Println("监控完成 ")
}

// 方式2: 使用select
var wg5 sync.WaitGroup

func cpuInfo5(ctx context.Context) {
	// 这里能拿到一个请求的id
	fmt.Printf("traceid: %s\r\n", ctx.Value("traceid"))
	defer wg5.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("退出cpu监控")
			return
		default: // 不阻塞走的方法
			// 每2秒钟获取cpu的信息
			time.Sleep(1 * time.Second)
			fmt.Println("cpu的信息")
		}
	}
}
