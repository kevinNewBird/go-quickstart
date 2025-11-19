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
### 2.5.1.语法格式
channel底层是一个环形数组。
```properties
var channel_name chan channel_type

说明：
channel_name: 定义的channel的变量名
channel_type: 指定channel管道交换的数据类型
```
### 2.5.2.channel的放值和取值
```go
     // 1.channel的初始化, 初始化大小如果为0在放值的时候会阻塞，从而报错deadlock
	var message chan string = make(chan string, 1)

	// 2.放值到channel管道中
	message <- "hello"

	// 3.从channel管道中取值赋给data变量
	data := <-message
```
### 2.5.3.channel的有缓冲和无缓冲
有缓冲是指在初始化时，长度大于0；而无缓冲则是指初始化时，长度等于0。<br/>
无缓冲区：必须要开goroutine，否则会有deadlock（go语言的happen-before机制保证了其的正确性，即内存屏障机制）。<br/>
两者的适用场景说明：
- 无缓冲区channel适用于通知，比如B要第一时间知道A是否已经完成；
- 有缓冲区channel适用于生产者和消费者之间的通讯
```go
// 1.有缓冲的channel
var cache chan string = make(chan string, 1)

// 2.无缓冲的channel 
var nocache chan string = make(chan string, 0)
```
go 中channel的应用场景：
- 1.消息传递、消息过滤
- 2.信号广播
- 3.事件订阅和广播
- 4.任务分发
- 5.结果汇总
- 6.并发控制
- 7.同步和异步等等
### 2.5.4.channel的forrange遍历获取
当channel中不停的塞入数据时（生产者-消费者），可以通过forrange去进行遍历获取。同时挡不需要再接收数据了，也可以通过close方法通知到goroutine
退出channel(已经关闭的channel，不能再放值了，但是可以再取值)。<br/>
```go
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
```
### 2.5.5.单向channel
默认情况下，channel时双向的，即可以发送也可以接收。 但是，我们经常把一个channel作为参数传递，希望对方是单向传递。<br/>
```go
//  单向channel：只能写入float64的数据
var wCh chan<- float64
// 单向channel：只能读取int的数据
var rCh <-chan int
```
### 2.5.6.select语句
类似于switch...case语句，写法上很相似。不同的是，select语句更多是用于goroutine中，其功能和linux中的selec、poll、epoll大体是差不多的。<br/>
select作用于多个channel，其会在多个channel中选择一个目前已经就绪的channel。<br/>
实现代码片段：
```go
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
		}
	}
}
```
### 2.5.7.上下文context
用于goroutine的信息传递。context包提供了三种函数：WithCancel， WithTimeout， WithValue。