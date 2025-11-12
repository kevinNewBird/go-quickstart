package main

import "fmt"

func add(a, b interface{}) interface{} {
	switch a.(type) { // 接口的判断
	case int:
		// 接口的断言
		return a.(int) + b.(int)
	case int32:
		return a.(int32) + b.(int32)
	case int64:
		return a.(int64) + b.(int64)
	case float32:
		return a.(float32) + b.(float32)
	case float64:
		return a.(float64) + b.(float64)
	default:
		panic("Unsupported type!")
	}
}

func main() {
	// 案例：计算器实现一个通用的加法。传统的做法是不同的类型实现一个函数，但是这样是会有很多不优雅的代码。
	// 使用interface{}可以接收任意类型的数据，故通用的做法是通过interface{} + 断言
	a := 1.0
	b := 2.0
	ret := add(a, b)
	// 3
	fmt.Println(ret)
}
