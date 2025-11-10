package main

import "fmt"

type Person1 struct {
	name string
	age  int
}

type Student1 struct {
	p     Person1
	score float32
}

type Person2 struct {
	name string
	age  int
}

type Student2 struct {
	Person2
	score float32
}

func main() {
	stu1 := Student1{
		Person1{
			"tommy", 18,
		},
		95.6,
	}

	fmt.Println(stu1.p.name)

	stu2 := Student2{
		Person2{
			"kitty", 29,
		},
		98.8,
	}
	fmt.Println(stu2.name)
}
