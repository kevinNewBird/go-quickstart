package main

import "fmt"

func main() {
	switchDemo1()
	switchDemo2()
	switchDemo3()
}

func switchDemo1() {
	// 案例：中文的星期几， 输出对应的英文
	// 输出结果: Wednesday
	day := "星期三"
	switch day {
	case "星期一":
		fmt.Println("Monday")
	case "星期二":
		fmt.Println("Tuesday")
	case "星期三":
		fmt.Println("Wednesday")
	case "星期四":
		fmt.Println("Thursday")
	case "星期五":
		fmt.Println("Friday")
	default:
		fmt.Println("Unknown")
	}
}

func switchDemo2() {
	// 案例：中文/英文的星期几， 输出对应的数值
	// 输出结果: 3
	day := "星期三"
	switch day {
	case "星期一", "Monday":
		fmt.Println(1)
	case "星期二", "Tuesday":
		fmt.Println(2)
	case "星期三", "Wednesday":
		fmt.Println(3)
	case "星期四", "Thursday":
		fmt.Println(4)
	case "星期五", "Friday":
		fmt.Println(5)
	default:
		fmt.Println(0)
	}
}

func switchDemo3() {
	// 案例: 分数值转换为等级
	// 输出结果: D
	score := 60
	switch {
	case score < 60:
		fmt.Println("E")
	case score >= 60 && score < 70:
		fmt.Println("D")
	case score >= 70 && score < 80:
		fmt.Println("C")
	case score >= 80 && score < 90:
		fmt.Println("B")
	case score >= 90:
		fmt.Println("A")
	default:
		fmt.Println("Unknown")
	}
}
