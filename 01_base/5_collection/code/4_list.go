package main

import (
	"container/list"
	"fmt"
)

func main() {
	// 1.声明（定义）list
	listDefine()

	// 2.遍历list
	listTraversal()

	// 3.插入元素
	listInsert()

	// 4.删除元素
	listDelete()
}

func listDefine() {
	var myList list.List
	myList.PushBack("go")
	myList.PushBack("python")

	// {{0x1400011e120 0x1400011e150 <nil> <nil>} 2}
	fmt.Println(myList)
}

// 遍历
func listTraversal() {
	var myList list.List
	// 尾插法
	myList.PushBack("go engineering")
	myList.PushBack("java engineering")
	// 	头插法
	myList.PushFront("python engineering")

	// 正向遍历
	for i := myList.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}

	fmt.Println("====================================")
	// 反向遍历
	for i := myList.Back(); i != nil; i = i.Prev() {
		fmt.Println(i.Value)
	}

	fmt.Println("************************************")
}

func listInsert() {
	// 插入需要先找到要插入位置的元素（遍历）
	var myList list.List
	// 尾插法
	myList.PushBack("a")
	myList.PushBack("b")
	myList.PushBack("d")
	myList.PushBack("e")

	i := myList.Front()
	for ; i != nil; i = i.Next() {
		if i.Value.(string) == "d" {
			break
		}
	}
	// 在元素d前面插入c
	myList.InsertBefore("c", i)
	// a b c d e
	for i := myList.Front(); i != nil; i = i.Next() {
		fmt.Print(i.Value)
		fmt.Print(" ")
	}
	fmt.Println()
}

func listDelete() {
	// 插入需要先找到要插入位置的元素（遍历）
	var myList list.List
	// 尾插法
	myList.PushBack("a")
	myList.PushBack("b")
	myList.PushBack("c")
	myList.PushBack("d")
	myList.PushBack("e")

	i := myList.Front()
	for ; i != nil; i = i.Next() {
		if i.Value.(string) == "c" {
			break
		}
	}

	// 删除元素
	myList.Remove(i)
	// a b d e
	for i := myList.Front(); i != nil; i = i.Next() {
		fmt.Print(i.Value)
		fmt.Print(" ")
	}
	fmt.Println()
}
