package main

import "fmt"

func main() {
	// 算术运算符
	mathOperators()

	// 逻辑运算符
	logicalOperators()

	// 位运算符
	bitOperators()

	// 其他运算符
	otherOperators()
}

// 算术运算符 + - * / % ++ --
func mathOperators() {
	fmt.Println(">>> mathOperators: ")
	var a, b = 1, 2
	var aStr, bStr = "abc", "def"
	// 3
	fmt.Println(a + b)
	// abcdef
	fmt.Println(aStr + bStr)
	// 1
	fmt.Println(a % b)
	a++
	// 2
	fmt.Println(a)
	fmt.Println("=============END============")
	fmt.Println()
}

// 逻辑运算符 && || !
func logicalOperators() {
	fmt.Println(">>> logicalOperators: ")
	var aBool, bBool = true, false
	fmt.Println(aBool && bBool)
	fmt.Println(aBool || bBool)
	fmt.Println(!aBool)

	fmt.Println("=============END============")
	fmt.Println()
}

// 位运算符
func bitOperators() {
	fmt.Println(">>> bitOperators: ")
	// 60:  0011 1100
	// 13:  0000 1101
	var a, b = 60, 13
	// 12
	// 61
	// 49
	fmt.Println(a & b)
	fmt.Println(a | b)
	fmt.Println(a ^ b)
	// 240 : 1111 0000
	fmt.Println(a << 2)
	// 3 :  0000 0011
	fmt.Println(b >> 2)
	fmt.Println("=============END============")
	fmt.Println()
}

// 其他运算符 & *
func otherOperators() {
	fmt.Println(">>> otherOperators: ")
	var a = 10
	// 0x1400010e020
	fmt.Println(&a) // 取地址
	var b *int = &a // 指针，就是地址
	// 0x1400010e020
	fmt.Println(b)
	fmt.Println("=============END============")
	fmt.Println()
}
