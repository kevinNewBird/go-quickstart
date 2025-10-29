package main

import "fmt"

// 函数a返回值 int ， bool
func a() (int, bool) {
	return 0, false
}

func main() {
	// 匿名变量：定义一个变量但是不使用它，即使用下划线
	var _ int
	// 调用a函数，此时如果只想使用一个就很麻烦了（定义了必须要使用）
	// 这个时候，就可以使用匿名变量
	r1, _ := a()

	// 0
	fmt.Println(r1)
}
