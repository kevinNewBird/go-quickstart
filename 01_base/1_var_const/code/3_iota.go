package main

import "fmt"

func main() {
	// iota,特殊常量，可以认为是一个可以被编译器修改的常量
	// 1.普通的定义方式
	const (
		ERR1 = 1
		ERR2 = 2
		ERR3 = 3
		ERR4 = 4
	)

	// 2.ioto的定义方式，使用这种方式后值会有编译器进行自动的递增
	const (
		INFO1 int = iota + 1
		INFO2
		INFO3
		INFO4
	)

	// 1 2 3 4
	fmt.Println(INFO1, INFO2, INFO3, INFO4)

	// 3.ioto的定义方式，使用这种方式后值会有编译器进行自动的递增
	const (
		WARN1 int = iota + 1
		WARN2
		WARN3 = "ha" // iota内部仍然会增加计数器
		WARN4 = iota
	)

	// 1 2 ha 3
	fmt.Println(WARN1, WARN2, WARN3, WARN4)

	// 4.ioto的定义方式，使用这种方式后值会有编译器进行自动的递增
	const (
		DEBUG1 int = iota + 1
		DEBUG2
		DEBUG3 = "ha" // iota内部仍然会增加计数器
		DEBUG4
		DEBUG5
		DEBUG6 = iota
	)
	// 1 2 ha ha ha 5
	fmt.Println(DEBUG1, DEBUG2, DEBUG3, DEBUG4, DEBUG5, DEBUG6)
}
