package main

import "fmt"

func main() {
	// 常量， 定义的时候就指定值，不能修改
	// 隐式定义
	const NAME1 = "A"
	// 显式定义
	const NAME2 string = "B"

	// 以组的方式定义
	const (
		UNKNOWN = 1
		FEMALE  = 2
		MALE    = 3
	)

	// 以组的方式定义的规则
	const (
		X int = 16
		Y
		S = "abc"
		Z
	)
	// 16 16 abc abc
	fmt.Println(X, Y, S, Z)
}
