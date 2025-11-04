package main

import (
	"fmt"
	"sync"
)

func main() {
	// 1.defer的定义
	deferDefine()
	// 2.一个语句块中定义多个defer的执行顺序
	deferSort()

	// 3.defer控制函数
	ret := deferReturn()
	// 11
	fmt.Printf("ret = %d\r\n", ret)
}

func deferDefine() {
	// 连接数据库、打开文件、开始锁这些操作，无论如何最后都要记得去释放资源
	var mu sync.Mutex
	mu.Lock()
	defer mu.Unlock() // go语言做了优化，定义defer后，defer后面的代码会放在函数return之前执行。
}

func deferSort() {
	// defer的执行顺序: 先进后出
	// A3 A2 A1
	defer fmt.Println("A1")
	defer fmt.Println("A2")
	fmt.Println("A3")
}

func deferReturn() (ret int) {
	defer func() {
		ret++
	}()
	// 返回之前，defer会先被执行，最终返回的值将会是11
	return 10
}
