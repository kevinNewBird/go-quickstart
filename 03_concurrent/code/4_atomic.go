package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var waitG sync.WaitGroup
var count atomic.Int32

func add2() {
	defer waitG.Done()
	for i := 0; i < 10000; i++ {
		count.Add(1)
	}
}

func sub2() {
	defer waitG.Done()
	for i := 0; i < 10000; i++ {
		count.Add(-1)
	}
}

func main() {
	waitG.Add(2)
	go add2()
	go sub2()

	waitG.Wait()
	fmt.Println(count.Load())

	// atomic的另外一种使用方法
	var count2 int32
	atomic.AddInt32(&count2, 1)
	fmt.Println(count2)
}
