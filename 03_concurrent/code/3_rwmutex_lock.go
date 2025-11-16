package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	var rw sync.RWMutex // 定义读写锁

	wg.Add(3)
	var sharedData int

	go func() {
		defer wg.Done()
		// 加写锁
		rw.Lock()
		sharedData = 1
		defer rw.Unlock()
		fmt.Println("get write lock")
	}()

	go func() {
		defer wg.Done()
		// 加读锁
		rw.RLock()
		defer rw.RUnlock()
		fmt.Println("get read1 lock", sharedData)
	}()

	go func() {
		defer wg.Done()
		// 加读锁
		rw.RLock()
		defer rw.RUnlock()
		fmt.Println("get read2 lock", sharedData)
	}()

	wg.Wait()
}
