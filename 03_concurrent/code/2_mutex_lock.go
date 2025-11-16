package main

import (
	"fmt"
	"sync"
)

func main() {
	// 1.互斥锁Mutex的使用
	wg.Add(2)
	go add()
	go sub()
	wg.Wait()

	fmt.Println(total)
}

var wg = sync.WaitGroup{}
var lock sync.Mutex
var total int

func add() {
	defer wg.Done()
	for i := 0; i < 10000; i++ {
		lock.Lock()
		total += 1
		lock.Unlock()
	}
}

func sub() {
	defer wg.Done()
	for i := 0; i < 10000; i++ {
		lock.Lock()
		total -= 1
		lock.Unlock()
	}
}
