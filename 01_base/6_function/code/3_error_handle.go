package main

import (
	"errors"
	"fmt"
)

func main() {
	// 1.error
	_, err := errorHandle()
	if err != nil {
		fmt.Println(err)
	}

	// 2.panic会导致程序的退出，平时开发中不要随便用。一般在服务启动的过程中，必须有些依赖服务准备好，日志文件存在、mysql能连接通等，
	// 这个时候才能启动服务，如果服务启动检查中发现了这些人和一个不满足那就调用panic
	//panicHandle()

	// 3.recover
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover err:", r)
		}
	}()

	panicHandle() // 只是为了触发recover
}

func errorHandle() (int, error) {
	return 0, errors.New("this is an error")
}

func panicHandle() {
	// panic是个内置函数，其会导致程序退出
	panic("this is a panic")
	fmt.Println("this is a func")
}
