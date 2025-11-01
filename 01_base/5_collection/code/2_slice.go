package main

import "fmt"

func main() {
	// 1.切片的定义
	sliceDefine()

	// 2.切片追加元素append
	sliceAppend()

	// 3.切片的初始化
	sliceInitial()

	// 4.访问切片的元素
	sliceAccess()

	// 5.删除切片的元素
	sliceDelete()

	// 6.拷贝
	sliceCopy()
}

func sliceDefine() {
	// 切片的定义
	// 1.方式1
	var courses1 []string
	// 2.方式2
	var courses2 [][4]string
	// 3.方式3
	var courses3 [][]string
	fmt.Printf("%T\n", courses1) // []string
	fmt.Printf("%T\n", courses2) //[][4]string
	fmt.Printf("%T\n", courses3) //[][]string
}

func sliceAppend() {
	// 1.添加元素
	var courses1 []string
	courses1 = append(courses1, "go", "java")
	// [go java], []string
	fmt.Printf("%v, %T\n", courses1, courses1)

	var courses2 [][4]string
	//elems := [4]string{"1", "2", "3", "4"}
	courses2 = append(courses2, [4]string{"1", "2", "3", "4"}, [4]string{"5", "6", "7", "8"})
	// [[1 2 3 4] [5 6 7 8]], [][4]string
	fmt.Printf("%v, %T\n", courses2, courses2)

	// 2.合并两个切片
	cousers3 := []string{"python", "c", "c++"}
	// 这里省略号的目的是表示把切片打散，由go底层来完成
	cousers3 = append(cousers3, courses1...)
	// [python c c++ go java], []string
	fmt.Printf("%v, %T\n", cousers3, cousers3)
	cousers4 := []string{"python", "c", "c++"}
	cousers4 = append(cousers4, courses1[1:]...)
	// [python c c++ java], []string
	fmt.Printf("%v, %T\n", cousers4, cousers4)
}

func sliceInitial() {
	// 1.从数组直接创建
	var courses1 [3]string = [3]string{"a", "b", "c"} // 数组
	// 左闭右开，[0,2)
	courses1Slice := courses1[0:2]
	// [a b], []string
	fmt.Printf("%v, %T\n", courses1Slice, courses1Slice)

	// 2.使用[]type{..}, 比如[]string{"a","b","c"}
	var courses2 []string = []string{"a", "b", "c"} // 切片
	courses2Slice := courses2[0:1]
	// [a], []string
	fmt.Printf("%v, %T\n", courses2Slice, courses2Slice)

	// 3.make函数(预先分配3个元素空间，性能较高)
	courses3 := make([]string, 3)
	// 下面的方式不会自动扩容，需要使用append才会扩容
	courses3[0] = "x"
	courses3[1] = "y"
	courses3[2] = "z"
	//courses3[3] = "d"
	// [x y z], []string // 会报错
	fmt.Printf("%v, %T\n", courses3, courses3)
	courses3 = append(courses3, "a")
	// [x y z a], []string
	fmt.Printf("%v, %T\n", courses3, courses3)

	// 分片错误的赋值方式
	// 初始化一个空间为0的分片courses4
	var courses4 []string
	//courses4[0] = "a" // 报错，因为空间为0。正确的方式为append
	fmt.Printf("%v, %T\n", courses4, courses4)
}

func sliceAccess() {
	// 访问切片的元素
	// 1.访问单个
	var courses1 []string = []string{"a", "b", "c"}
	// a, []string
	fmt.Printf("%s, %T\n", courses1[0], courses1)

	// 2.访问多个
	// [a b c]
	fmt.Println(courses1[:])
	// [b c]
	fmt.Println(courses1[1:])
	// [a b]
	fmt.Println(courses1[:2])
}

func sliceDelete() {
	var courses []string = []string{"a", "b", "c", "d", "e", "f"}
	// 删除c元素
	coursesSlice := append(courses[:2], courses[3:]...)
	// [a b d e f], []string
	fmt.Printf("%v, %T\n", coursesSlice, coursesSlice)
}

func sliceCopy() {
	var courses []string = []string{"a", "b", "c", "d", "e", "f"}
	// 1.浅拷贝
	coursesCopy1 := courses
	coursesCopy2 := courses[:]

	// 2.深拷贝
	var coursesDeepCopy = make([]string, len(courses))
	copy(coursesDeepCopy, courses)

	courses[0] = "java"
	// [java b c d e f], []string
	fmt.Printf("%v, %T\n", coursesCopy1, coursesCopy1)
	// [java b c d e f], []string
	fmt.Printf("%v, %T\n", coursesCopy2, coursesCopy2)
	// [a b c d e f], []string
	fmt.Printf("%v, %T\n", coursesDeepCopy, coursesDeepCopy)
}
