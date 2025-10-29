package main

import "fmt"

func main() {
	// go基本数据类型，三大类：布尔、数值、字符/字符串
	var a int8
	var b int16
	var c int32
	var d int64
	var ua uint8
	var ub uint16
	var uc uint32
	var ud uint64
	// 0 0 0 0 0 0 0 0
	fmt.Println(a, b, c, d, ua, ub, uc, ud)

	var e int // 动态类型： 根据操作系统来定义的，系统是64位，那么int就表示64位；反之，系统是32位，那么int就表示32位
	// 0
	fmt.Println(e)

	var f1 float32
	var f2 float64
	// 0 0
	fmt.Println(f1, f2)
	f1 = 3.14
	f2 = 3.14

	var bc byte // 主要用于存放字符
	bc = 'a'
	// 97
	fmt.Println(bc)
	// character=a
	fmt.Printf("character=%c\n", bc)

	var ra rune = 98 // 也是字符（范围更大，比如中文/韩文等等）
	// 98
	fmt.Println(ra)
}
