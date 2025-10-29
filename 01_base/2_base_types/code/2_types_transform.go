package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 简单转换
	testBase()

	// parse函数
	testParse()

	// format函数
	testFormat()
}

func testBase() {
	// 浮点型
	a := 5.0

	// 转换为int类型
	b := int(a)
	fmt.Println(b)

	// go允许在底层结构相同的两个类型之间互转。例如：IT类型的底层是int类型
	type IT int   // 定义别名
	var a2 IT = 5 // IT类型，底层是int
	// 将a(IT)转换为int
	b2 := int(a2)

	fmt.Println(b2)
}

func testParse() {
	// 字符串转数字
	var istr = "12"
	myint, err := strconv.Atoi(istr)
	if err != nil {
		fmt.Println("convert error")
	}
	fmt.Println(myint)

	// 数字转字符串
	mystr := strconv.Itoa(myint)
	fmt.Println(mystr)

	// 字符串转换为float
	f32, ok := strconv.ParseFloat("3.1415", 32)
	if ok != nil {
		fmt.Println("convert error")
	}
	fmt.Printf("float 32bit: %f\n", f32)

	// 10进制int64
	i64, ok := strconv.ParseInt("-42", 10, 64)
	ix64, ok := strconv.ParseInt("-42", 8, 64)
	if ok != nil {
		fmt.Println("convert error")
	}
	// int 64bit: -42
	// int 64bit: -34
	fmt.Printf("int 64bit: %d\n", i64)
	fmt.Printf("int 64bit: %d\n", ix64)

	b1, err := strconv.ParseBool("true")
	b2, err := strconv.ParseBool("1")
	if err != nil {
		fmt.Println("convert error")
	}
	// bool string true: true
	// bool numberic true: true
	fmt.Printf("bool string true: %t\n", b1)
	fmt.Printf("bool numberic true: %t\n", b2)
}

func testFormat() {
	// 基础类型转换为字符串
	bStr := strconv.FormatBool(true)
	fmt.Printf("bool string true: %s\n", bStr)

	fStr := strconv.FormatFloat(3.1415, 'E', 5, 32)
	// float string 3.1415: 3.14150E+00
	fmt.Printf("float string 3.1415: %s\n", fStr)

	// int转换为16进制的字符串
	i64Str := strconv.FormatInt(-42, 16)
	// int64 string -42: -2a
	fmt.Printf("int64 string -42: %s\n", i64Str)
}
