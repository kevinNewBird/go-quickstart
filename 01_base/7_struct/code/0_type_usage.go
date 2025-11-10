package main

import (
	"fmt"
	"strconv"
)

func main() {
	/*
		type关键字的用途：
		- 1.定义结构体
		- 2.定义接口
		- 3.定义类型别名：为了更好的理解代码
		- 4.类型定义
		- 5.类型判断
	*/

	// 1.定义类型的别名
	typeAlias()

	// 2.类型定义
	typeDefinition()

	// 3.类型判断
	typeValidate()
}

func typeAlias() {
	type myInt = int
	var i myInt = 5
	var j int = 6
	// int: 5
	fmt.Printf("%T: %v\r\n", i, i)
	fmt.Println(i + j) // 在编译的时候，类型别名会被直接替换为int
}

// 自定义类型，并为其增加扩展方法
type MyInt32 int32

func (mi MyInt32) string() string {
	return strconv.Itoa(int(mi))
}

func typeDefinition() {
	// 自定义类型，区别于类型别名
	type MyInt int

	var i MyInt = 5

	// main.MyInt: 5
	fmt.Printf("%T: %v\r\n", i, i)
	// // 类型定义的myInt和int是不同的，无法直接和int进行运算
	var j int = 6
	fmt.Printf("%T: %v\r\n", j, j) // int: 6
	//fmt.Println(i + 6) // 无法通过编译
	// 要实现自定义类型和原类型的运算，需要进行转换
	// 11
	fmt.Println(i + MyInt(j)) // fmt.Println(int(i) + j)

	// 为什么要有自定义类型呢？比如int，假设想要其具有额外的方法，比如i.string(), 使用type实现自定义类型，
	//就可以为基础类型int添加string()等自定义方法
	var m MyInt32 = 10
	// main.MyInt32: 10, 10
	fmt.Printf("%T: %v, %s\r\n", m, m, m.string())
}

func typeValidate() {
	// 类型判断
	var a interface{} = "abc"
	// 1.获取类型
	switch a.(type) {
	case string:
		fmt.Println("a string")
	case int:
		fmt.Println("a int")
	default:
		fmt.Println("unkown type")
	}
	// 2.类型的断言
	// 把interface直接转换为string
	m := a.(string)
	fmt.Println(m)
}
