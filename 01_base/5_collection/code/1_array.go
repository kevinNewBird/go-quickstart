package main

import (
	"fmt"
)

func main() {
	// 1.数组的类型验证
	arrayType()

	// 2.数组的遍历
	arrayTraversal()

	// 3.数组的初始化
	arrayInitial()

	// 4.数组的比较: 比较的前提是类型必须相同
	arrayCompare()

	// 5.多维数组
	arrayRagged()
}

func arrayType() {
	var courses1 [3]string // courses的类型： 只有3个元素的数组类型
	var courses2 [4]string
	// [3]string
	fmt.Printf("%T\n", courses1)
	// [4]string
	fmt.Printf("%T\n", courses2)

	// 注意：[]string是切片，而不是数组。也就是说[3]string和[]string是不同的类型。
}

// 数组的遍历
func arrayTraversal() {
	var courses1 [3]string

	courses1[0] = "go"
	courses1[1] = "gin"
	courses1[2] = "grpc"
	// [go gin grpc]
	fmt.Println(courses1)

	for _, course := range courses1 {
		fmt.Println(course)
	}
}

func arrayInitial() {
	// 数组的初始化
	// 1.方式1
	var courses1 [3]string = [3]string{"java", "c++", "python"}
	// [java c++ python]
	fmt.Println(courses1)

	// 2.方式2: 指定第三个元素为java，其余为默认值
	courses2 := [3]string{2: "java"}
	// [  java] 3
	fmt.Println(courses2, len(courses2))

	// 3.方式3: 变长数组
	courses3 := [...]string{2: "java"}
	// [  java] 3
	fmt.Println(courses3, len(courses3))
}

func arrayCompare() {
	courses1 := [...]string{"go", "java", "grpc"}
	courses2 := [...]string{"go", "java", "python"}
	fmt.Println("courses1 == courses2 ?", courses1 == courses2)
}

func arrayRagged() {
	// 多维数组
	// 3行4列
	var courses1 [3][4]string
	// 第一行
	courses1[0] = [4]string{"go", "1h", "tommy", "18"}
	courses1[1] = [4]string{"c", "1.5h", "kimmy", "28"}
	courses1[2] = [4]string{"groovy", "0.5h", "amy", "26"}

	// [[go 1h tommy 18] [c 1.5h kimmy 28] [groovy 0.5h amy 26]] 3
	fmt.Println(courses1, len(courses1))

	// 遍历
	// go 1h tommy 18
	// c 1.5h kimmy 28
	// groovy 0.5h amy 26
	for i := 0; i < len(courses1); i++ {
		for j := 0; j < len(courses1[i]); j++ {
			fmt.Print(courses1[i][j], " ")
		}
		fmt.Println()
	}

	// 行： [go 1h tommy 18]
	// go 1h tommy 18
	// 行： [c 1.5h kimmy 28]
	// c 1.5h kimmy 28
	// 行： [groovy 0.5h amy 26]
	// groovy 0.5h amy 26
	for _, row := range courses1 {
		fmt.Println("行：", row)
		for _, column := range row {
			fmt.Print(column, " ")
		}
		fmt.Println()
	}
}
