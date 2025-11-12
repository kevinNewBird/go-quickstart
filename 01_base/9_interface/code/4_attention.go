package main

import "fmt"

// 使用注意事项：
// 1.interface{}可以接收任何类型
// 2.interface{}可变参数不能只能接收interface{}类型的切片等，不能是[]string等，必须进行转换

func main() {
	var data1 []interface{} = []interface{}{"tommy", 2.3, 3}
	// 1.接收interface{}类型的切片，可正常执行
	cPrint(data1...)

	// 2.[]string类型的切片
	var data2 []string = []string{"c", "b"}
	//cPrint(data2...) // 编译器报错：cannot use data2 (variable of type []string) as []interface{} value in argument to cPrint
	// 2.1.必须要进行转换
	var convert2 []interface{}
	for _, ele := range data2 {
		convert2 = append(convert2, ele)
	}
	cPrint(convert2...)
}

func cPrint(data ...interface{}) {
	for _, ele := range data {
		fmt.Println(ele)
	}
}
