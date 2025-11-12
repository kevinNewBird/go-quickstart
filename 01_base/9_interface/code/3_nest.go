package main

import "fmt"

// 接口的嵌套
func main() {
	var rw MyReadWriter = &MyReadWriterImpl{}
	rw.read()
	rw.write("hello world")
	rw.readWrite()
}

// 1.定义基础接口MyReader和MyWriter
type MyReader interface {
	read() string
}

type MyWriter interface {
	write(string)
}

// 2.继承基础接口的MyReadWriter接口：也就是嵌套
type MyReadWriter interface {
	MyReader
	MyWriter

	// 自己的接口
	readWrite()
}

// 3.实现MyReadWriter
type MyReadWriterImpl struct {
}

func (m *MyReadWriterImpl) read() string {
	fmt.Println("read me")
	return ""
}

func (m *MyReadWriterImpl) write(s string) {
	fmt.Println("write me")
}

func (m *MyReadWriterImpl) readWrite() {
	fmt.Println("readWrite me")
}
