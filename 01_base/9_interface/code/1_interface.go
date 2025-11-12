package main

import "fmt"

// 接口的声明， 万能的接口interface{} 等价于java中的Object
type Duck interface {
	Gaga() (a string)
	Swimming()
}

// 定义需要实现Duck接口的结构体
type PskDuck struct {
	legs int
}

// 绑定PskDuck到接口Duck： 即实现Duck接口的所有方法，并使用被实现结构体PskDuck的指针接收器
func (pd *PskDuck) Gaga() (a string) {
	fmt.Println("Gaga...")
	return ""
}

func (pd *PskDuck) Swimming() {
	fmt.Println("Swimming...")
}

func main() {
	// 多态的方式： 绑定的是指针接收器，故采用如下的初始化方式
	var duck Duck = &PskDuck{4}
	// Swimming...
	duck.Swimming()
	// Gaga...
	duck.Gaga()
}
