package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Animal struct {
	name string
	age  int
}

// 结构体方法： 值传递方式
func (s Animal) toString() string {
	builder := strings.Builder{}
	builder.WriteString("Animal{name=\"" + s.name + "\", age=" + strconv.Itoa(s.age) + "}")
	s.name = "doggy2"
	return builder.String()
}

// 结构体方法：指针
func (s *Animal) setName(name string) {
	s.name = name
}

func main() {
	p := Animal{
		"doggy", 3,
	}
	// Animal{name="doggy", age=3}
	fmt.Println(p.toString())
	// {doggy 3}
	fmt.Println(p)

	p.setName("cat")
	// Animal{name="cat", age=3}
	fmt.Println(p.toString())
}
