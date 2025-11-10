package main

import "fmt"

func main() {
	// 1.指针的声明和赋值
	pointerDeclare()

	// 2.指针的初始化
	pointerInitial()

	// 3.交换值
	pointerSwap()
}

func pointerDeclare() {
	i := 1
	var po *int = &i

	// 0x1400009a020
	// 在go语言中，指针存在以下特性：
	// 1. 指针是不能直接参与运算的（c/c++可以执行指针的加减改变指向地址）
	// 2. 指针可以通过unsafe包下的方法实现指针的运算
	fmt.Println(po)
	// 1
	fmt.Println(*po)

	*po = 10
	// 0x14000114008
	fmt.Println(po)
	// 10
	fmt.Println(*po)
}

type Grade struct {
	name string
}

func pointerInitial() {
	// 1.初始化方式1
	pi1 := &Grade{name: "math"}
	// &{math}
	fmt.Println(pi1)

	// 2.初始化方式2
	var grade = Grade{name: "english"}
	pi2 := &grade
	// &{english}
	fmt.Println(pi2)

	// 3.初始化方式3
	pi3 := new(Grade)
	// &{}
	fmt.Println(pi3)
	pi4 := new(int)
	// 0x1400010e028
	fmt.Println(pi4)
}

func swap(a, b *int) {
	*a, *b = *b, *a
}

func pointerSwap() {
	a, b := 1, 2
	swap(&a, &b)
	fmt.Println(a, b)
}
