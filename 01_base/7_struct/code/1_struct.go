package main

import "fmt"

type Person struct {
	name   string
	age    int
	height float64
}

func main() {
	// 1.结构体的初始化
	structDefineInitial()

	// 2.结构体的赋值取值
	structSetGet()

	// 3.匿名结构体
	structAnonymous()
}

func structDefineInitial() {
	// 1.初始化的两种方式
	// 1.1.不指定属性名的方式（必须初始化所有字段，否则将会编译报错）
	p1 := Person{"tommy", 18, 1.78}
	// 1.2.执行属性名的方式（选择性初始化字段）
	p2 := Person{name: "tommy"}
	fmt.Println(p1, p2)

	// 2.结构体的切片的两种方式
	// 2.1.定义变量然后使用append添加
	var pSlice1 []Person
	pSlice1 = append(pSlice1, p1)
	pSlice1 = append(pSlice1, Person{name: "kitty", age: 18})

	// 2.2.初始化时直接添加
	var pSlice2 []Person = []Person{
		{"mama", 28, 1.78},
		{name: "daddy"},
	}
	fmt.Println(pSlice1, pSlice2)
}

func structSetGet() {
	// 1.结构的赋值
	var p Person
	p.name = "GG"
	p.age = 19
	fmt.Println("赋值：", p)

	// 2.结构体的取值
	fmt.Println("取值：", p.name)
}

// 匿名结构体
func structAnonymous() {
	address := struct {
		province   string
		city       string
		town       string
		population int
	}{
		"四川省", "成都市", "成华区", 100,
	}

	fmt.Println(address.province, address.city, address.town, address.population)
}
