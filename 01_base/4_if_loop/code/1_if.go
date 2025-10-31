package main

import "fmt"

func main() {
	age := 20
	country := "中国"
	// if条件判断
	if age < 18 && country == "中国" {
		fmt.Println("未成年")
	} else if age > 18 && age < 30 {
		fmt.Println("成年人")
	} else if (age > 30 && age < 60) || (age > 30 && age < 50) {
		fmt.Println("中年人")
	} else {
		fmt.Println("老年人")
	}
}
