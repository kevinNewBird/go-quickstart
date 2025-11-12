package main

import "fmt"

func main() {
	var p1 Writer = &Printer{&FileWriter{file: "a.txt"}}
	// file writer write: hello
	p1.write("hello")

	var p2 Writer = &Printer{&DbWriter{host: "192.168.1.1", port: 5432}}
	// db writer write: hello
	p2.write("hello")
}

// 多接口实现
type Writer interface {
	write(string) error
}

type Closer interface {
	close() error
}

// 实现打印机的写入和关闭功能接口
type Printer struct {
	Writer // 可以传入不同的写入实现，达到不同的功能，比如数据库写入器和文件写入器
}

func (w *Printer) write(s string) error {
	w.Writer.write(s)
	return nil
}

func (w *Printer) close() error {
	fmt.Println("close printer")
	return nil
}

// 文件写入器
type FileWriter struct {
	file string
}

func (f *FileWriter) write(s string) error {
	fmt.Println("file writer write:", s)
	return nil
}

type DbWriter struct {
	host string
	port int
}

func (d *DbWriter) write(s string) error {
	fmt.Println("db writer write:", s)
	return nil
}
