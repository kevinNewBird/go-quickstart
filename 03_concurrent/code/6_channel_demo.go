package main

import (
	"fmt"
	"time"
)

/**
使用两个goroutine交誉打印序到，一个goroutine打印数字，另外一个goroutine打印字母，最终效果如下:
12ABB4CD56EF78GH910IJ1112KL1314MN15160P17180R1920ST2122UV2324WX2526YZ2728
*/

func main() {
	go printNum()
	go printLetter()
	number <- true

	time.Sleep(10 * time.Millisecond)
}

var number, letter = make(chan bool), make(chan bool)

func printNum() {
	i := 1
	for {
		<-number
		fmt.Printf("%d%d", i, i+1)
		i += 2
		letter <- true
	}
}

func printLetter() {
	i := 0
	str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for {
		<-letter
		if i >= len(str) {
			return
		}
		fmt.Print(str[i : i+2])
		i += 2
		number <- true
	}
}
