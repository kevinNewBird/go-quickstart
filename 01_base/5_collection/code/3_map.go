package main

import "fmt"

func main() {
	// 1.map入门
	mapBasic()

	// 2.map初始化
	mapInitial()

	// 3.map的遍历
	mapTraversal()

	// 4.map的值获取
	mapGet()

	// 5.map的删除
	mapDelete()
}

func mapBasic() {
	courseMap := map[string]string{
		"go":     "go工程师",
		"python": "python工程师",
		"java":   "java工程师",
	}
	// map[go:go工程师 java:java工程师 python:python工程师]
	fmt.Println(courseMap)
	// 1.获取值。
	// go工程师
	fmt.Println(courseMap["go"])

	// 2.设置值
	courseMap["mysql"] = "mysql数据库"
	// map[go:go工程师 java:java工程师 mysql:mysql数据库 python:python工程师]
	fmt.Println(courseMap)
}

func mapInitial() {
	// 1.定义时初始化
	//var courseMap map[string]string
	courseMap := map[string]string{}
	courseMap["mysql"] = "mysql原理"
	// map[mysql:mysql原理]
	fmt.Println(courseMap)

	courseMap2 := make(map[string]string, 3)
	courseMap2["oracle"] = "oracle原理"
	// map[oracle:oracle原理]
	fmt.Println(courseMap2)
}

// map遍历
func mapTraversal() {
	courseMap := map[string]string{
		"go":     "go工程师",
		"python": "python工程师",
		"java":   "java工程师",
	}

	// 1.使用range遍历
	for k, v := range courseMap {
		fmt.Println(k, v)
	}

	for k := range courseMap {
		fmt.Println(k, courseMap[k])
	}
}

// map的值获取
func mapGet() {
	courseMap := make(map[string]string, 3)
	courseMap["java"] = "java课程"

	// 1.单值返回的获取方式
	mv1 := courseMap["java"]
	fmt.Println(mv1)

	// 2.两个值返回的获取方式
	// 第二个返回值表示值是否存在： true表示存在
	mv2, ok := courseMap["java"]
	if ok {
		fmt.Println("found the java", mv2)
	} else {
		fmt.Println("not exist the java")
	}

	// 上述的if...else可简写，如下：
	if v, ok := courseMap["java"]; !ok {
		fmt.Println("not exist the java")
	} else {
		fmt.Println("found the java", v)
	}
}

func mapDelete() {
	courseMap := map[string]string{
		"go":     "go工程师",
		"python": "python工程师",
		"java":   "java工程师",
	}

	delete(courseMap, "go")
	// map[java:java工程师 python:python工程师]
	fmt.Println(courseMap)
}
