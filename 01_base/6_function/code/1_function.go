package main

import "fmt"

func main() {
	// 1.返回值类型带变量名， 3 <nil>
	fmt.Println(add2(1, 2))
	// 2.可变参数的函数，10 <nil>
	fmt.Println(add3(1, 2, 3, 4))

	// 3.函数作为变量
	funcVar := add3
	// 12 (0x0,0x0)
	println(funcVar(2, 4, 6))

	// 返回是函数的调用
	calc("+")()
	// 9
	fmt.Println(calc2(add4, 4, 5))

	// 4.参数传递匿名函数
	funcIsParam()

	// 5.闭包
	fmt.Println("=================================")
	increment := autoIncrement()
	// 1
	fmt.Println(increment())
	// 2
	fmt.Println(increment())
}

/**
* 1.函数的定义
 */
func add1(a, b int) (sum int, err error) {
	sum = a + b
	return sum, err
}

/*
* 2.如果返回值定义了变量名，那么返回可以省略
 */
func add2(a, b int) (sum int, err error) {
	sum = a + b
	// 可省略sum和err， 但是不能省略return关键字
	return
}

/*
* 3.可变参数的函数
 */
func add3(items ...int) (sum int, err error) {
	for _, item := range items {
		sum += item
	}
	return sum, err
}

func add4(items ...int) int {
	sum := 0
	for _, item := range items {
		sum += item
	}
	return sum
}

// 返回值是函数
func calc(op string) func() {
	switch op {
	case "+":
		return func() {
			fmt.Println("这是加法")
		}
	case "-":
		return func() {
			fmt.Println("这是减法")
		}
	default:
		return func() {
			fmt.Println("这既不是加法也不是减法！")
		}
	}
}

func funcIsParam() {
	// 4.参数传递匿名函数
	// 4.1.使用变量接收匿名函数
	localFunc := func() {
		fmt.Println("这个是匿名函数")
	}
	localFunc()
	// 4.2.调用函数时，参数传递匿名函数
	sum := calc2(func(items ...int) int {
		sum := 0
		for _, value := range items {
			sum += value
		}
		return sum
	}, 1, 2, 3)
	// 6
	fmt.Println(sum)
}

// 参数是函数
func calc2(myFunc func(items ...int) int, items ...int) int {
	return myFunc(items...)
}

// 闭包
func autoIncrement() func() int {
	local := 0
	// 定义一个匿名函数
	return func() int {
		local++
		return local
	}
}
