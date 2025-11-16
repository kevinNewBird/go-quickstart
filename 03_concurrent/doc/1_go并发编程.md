# 1.引言
在python/java/php中，并发编程主要为多线程编程或者多进程编程，两者存在的问题主要是耗费内存和上下文切换。<br/>
在web2.0，内存、线程切换问题成为其主要矛盾，所以提出了用户级线程，也叫做绿程、轻量级线程、协程。在python中，提供了asyncio协程方案；php中为
swoole；java中为netty<br/>

Go语言由于是web2.0之后出现的语言，其诞生之后就只有协程可用--goroutine, 非常方便。<br/>
# 2.协程
协程的特点：内存占用小、切换快。<br/>
## 2.1.使用
非常简单，使用go关键字调用方法：
```properties
go method_name
```
## 2.2.原理
也就是GMP。G：gorouine--协程，go语言级别；M：thread线程，操作系统级别；P：处理器，用于调度goroutine给线程执行，go语言级别。

## 2.3.waitgroup
waitgroup包是go语言提供的一种通知机制：子goroutine通知主goroutine自己的状态。主要用于goroutine的执行等待，Add方法要和Done方法配套。
```go
// 2.waitGroup机制
	var wg sync.WaitGroup
	// 两种方式：1.一次性指定；2.递增的方式
	// 2.1.方式1: 一次性指定监控100个goroutine执行结束
	wg.Add(100)
	for i := 0; i < 100; i++ {
		// 2.2.方式2：递增的方式，每一个子goroutine+1
		//wg.Add(1)

		go func(i int) {
			fmt.Println(i)
			wg.Done() // 必须要有，表明当前子goroutine执行结束
		}(i)
	}

	// 等到所有子goroutine执行结束
	wg.Wait()
	fmt.Println("all done")
```
## 2.4.锁和原子操作
### 2.4.1.锁
锁：解决共享资源竞争，属于sync包。锁不能进行复制，一旦复制就失去了效果（锁本质是不断修改状态等信息，复制会把信息带过去，失去了锁的意义）。<br/>
锁的分类：
- 互斥锁mutex：其主要目的是确保同一时刻只有一个goroutine可以访问共享资源或执行特定代码块
- 读写锁RWMutex: 读协程之间并发，读和写之间串行（读写之间串行的原因：商品修改前为100，此时要修改为120，如果这个过程中被读取到了，那么就有问题了，所以读写间要串行）
### 2.4.2.原子Atomic
原子操作：属于atomic。
 
## 2.5.goroutine之间通信
go语言之间通信采用的是channel。类似于其他语言中的消息队列（生产者和消费者模型），再加上Go语言的语法糖，提供给使用者更好的使用体验。<br/>
