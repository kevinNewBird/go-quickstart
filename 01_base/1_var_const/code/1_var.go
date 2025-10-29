package main

import "fmt"

// 全局变量
// 1.单个定义
var name = ""
var age = 0

// 2.多个同时定义
var (
	name2 = ""
	age2  = 0
)

func main() {
	// go是静态语言，静态语言和动态语言的变量差异很大
	// 1.便利那个必需先定义后使用
	// 2.变量必须有类型
	// 3.类型确定后不能改变
	var name string
	name = "A"
	var age = 1
	height := 160

	// 注意：在go语言中局部变量定义了不使用是不行的
	fmt.Printf("username:%s, age: %d, height:%d\n", name, age, height)

	// 2.多变量定义
	var user1, user2, user3 string = "a", "b", "c"
	fmt.Println(user1, user2, user3)
}
