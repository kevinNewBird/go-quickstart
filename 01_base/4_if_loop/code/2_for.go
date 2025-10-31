package main

import (
	"fmt"
	"time"
)

func main() {
	// 1.for循环的基础操作
	baseOp()

	// 2.for range
	forRange()

	// 3.遍历打印含有中文的字符串
	forChineseStr()

	// 4.goto语句
	gotoDemo()
}

func baseOp() {
	// for循环
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	var i int
	for ; i < 3; i++ {
		fmt.Println(i)
	}

	i = 0
	for i < 3 {
		fmt.Println(i)
		i++
	}

	// 死循环
	var j int
	for {
		if j > 5 {
			break
		}
		fmt.Println(j)
		j++
		time.Sleep(20 * time.Millisecond)
	}
}

func forRange() {
	// for 遍历的另外一种方式：for range
	str := "hello world"
	// 0, h
	// 1, e
	// 2, l
	// 3, l
	// 4, o
	// 5,
	// 6, w
	// 7, o
	// 8, r
	// 9, l
	// 10, d
	for idx, char := range str {
		fmt.Printf("%d, %c\n", idx, char)
	}

	// 返回的是索引
	//0
	//1
	//2
	//3
	//4
	//5
	//6
	//7
	//8
	//9
	//10
	for idx := range str {
		fmt.Println(idx)
	}
}

func forChineseStr() {
	name := "体系课-go工程师"
	namebs := []rune(name) // 切片方式

	// 遍历含有中文字符的字符串
	/**
	体
	系
	课
	-
	g
	o
	工
	程
	师
	*/
	for i := 0; i < len(namebs); i++ {
		fmt.Printf("%c\n", namebs[i])
	}
}

func gotoDemo() {
	// goto语句的示例， 在嵌套循环中，如何跳出整个循环
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if j == 1 { // 期望是打印 0 0。 实际使用break无法跳过外循环，会打印0 0/ 1 0 / 2 0
				//break // 仅会跳出j代表的整个循环，而无法跳出外循环，此时就需要使用到goto循环了
				goto over
			}
			fmt.Println(i, j)
		}
	}

	// 定义goto语句块
over:
	fmt.Println("over")
}
