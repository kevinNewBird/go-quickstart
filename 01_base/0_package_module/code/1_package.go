package main

import (
	"fmt"
	"go-quickstart/01_base/0_package_module/code/test"          // 导入包路径
	course2 "go-quickstart/01_base/0_package_module/code/test2" // 导入包路径并指定别名：可以有效解决不同包路径下的package命名冲突
	. "go-quickstart/01_base/0_package_module/code/test3"       // 导入包路径下所有
	_ "go-quickstart/01_base/0_package_module/code/test4"       // 导入匿名包路径：初始化
)

func main() {
	// 1.调用：使用package名--course
	c := course.Course{Name: "tommy"}

	// test1: tommy
	fmt.Println(course.GetCourse(c))

	// 2.调用：使用package的别名
	c2 := course2.Course{Name: "tommy2"}
	// test2: tommy2
	fmt.Println(course2.GetCourse(c2))

	// 3.调用：导入包路径下所有，直接使用
	a1 := Animal{Name: "doggy"}
	fmt.Println(a1.Name)

	// 4.导入匿名包路径：用于初始化
	// test4 init
}
