package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// 1.转义字符
	escapeChar()

	// 2.格式化输出
	formatPrint()

	// 3.高性能字符串拼接
	hpJoin()

	// 4.字符串的比较
	compareStr()

	// 5.strings模块的常用方法
	stringsCommonOp()
}

// 1.转义字符
func escapeChar() {
	str := "go语言\"转义符\""
	str2 := `go语言"转义符2"` // ``表示里面的内容随意输入，都会语言输出。
	fmt.Println(str)
	fmt.Println(str2)
}

// 2.格式化输出
func formatPrint() {
	a := 65
	// A
	fmt.Printf("%c\n", a)
	// 'A'
	fmt.Printf("%q\n", a)
	// U+0041
	fmt.Printf("%U\n", a)
	// U+0041 'A'
	fmt.Printf("%#U\n", a)

	// 0x41
	n, err := fmt.Printf("%#x\n", a) // 输出到控制台
	// Printf返回字节数：5, 是否失败：<nil>
	fmt.Printf("Printf返回字节数：%d, 是否失败：%v\n", n, err)
	b := fmt.Sprintf("%U\n", a) // 不会输出到控制台
	// Sprintf格式化转换值： U+0041
	fmt.Println("Sprintf格式化转换值：", b)
}

// 3.高性能字符串拼接
func hpJoin() {
	// 初始化变量
	username := "high performance string join"
	age := 18
	address := "chengdu"
	mobile := "12345678"

	var ages []int = []int{1, 2, 3}
	fmt.Printf("用户名：%s，年龄：%d，地址：%s，电话：%s\r\n", username, ages, address, mobile)             // 性能一般
	userMsg := fmt.Sprintf("用户名：%T，年龄：%T，地址：%s，电话：%s\r\n", username, ages, address, mobile) // 性能一般
	fmt.Println(userMsg)

	// 【高性能】通过string的builder进行字符串拼接
	var builder strings.Builder
	builder.WriteString("用户名：")
	builder.WriteString(username)
	builder.WriteString("，年龄：")
	builder.WriteString(strconv.Itoa(age))
	builder.WriteString("，地址：")
	builder.WriteString(address)
	builder.WriteString("，电话：")
	builder.WriteString(mobile)

	re := builder.String()
	// 用户名：high performance string join，年龄：18，地址：chengdu，电话：12345678
	fmt.Println(re)
}

// 4.字符串的比较
func compareStr() {
	// 4.1.字符串的比较
	a := "hello"
	b := "hello"
	fmt.Println(a == b)

	// 4.2.字符串的大小比较
	fmt.Println(a > b)
	fmt.Println(a <= b)
}

// 5.strings模块的常用方法
func stringsCommonOp() {
	name := "体系课-go工程师"
	namebs := []rune(name) // 切片方式
	// 字符串长度： 21
	fmt.Println("字符串长度：", len(name))
	// 字符串长度： 9
	fmt.Println("字符串长度：", len(namebs))
	// 5.1.是否包含
	isC := strings.Contains(name, "go")
	fmt.Println("是否包含：", isC)

	// 5.2.出现次数
	count := strings.Count(name, "go")
	fmt.Println("出现次数", count)

	// 5.3.分割
	sre := strings.Split(name, "-")
	// 分割结果： [体系课 go工程师]
	fmt.Println("分割结果：", sre)

	// 5.4.是否包含前缀/后缀
	// 是否包含前缀： true
	//  是否包含后缀： true
	fmt.Println("是否包含前缀：", strings.HasPrefix(name, "体系"))
	fmt.Println("是否包含后缀：", strings.HasSuffix(name, "工程师"))

	// 5.5.子串出现的位置
	// 子串出现的位置： 10(是byte的位置， 中文三个byte)
	fmt.Println("子串出现的位置：", strings.Index(name, "go"))
	// 子串出现的位置： 10(是byte的位置， 中文三个byte)
	fmt.Println("子串出现的位置：", strings.IndexRune(name, []rune(name)[4]))

	// 5.6.子串替换
	// 子串替换： 体系课-GO工程师
	fmt.Println("子串替换：", strings.Replace(name, "go", "GO", 1))

	// 5.7.大小写转换
	fmt.Println("大小写转换:", strings.ToLower("GO"))
	fmt.Println("大小写转换:", strings.ToUpper("go"))

	// 5.8.去掉头尾的特殊字符
	// 去掉特殊字符：[java language]
	fmt.Printf("去掉特殊字符：[%s]\n", strings.Trim("#java language#", "#"))
	fmt.Printf("去掉特殊字符：[%s]\n", strings.Trim("#java language$", "#$"))
	// 去掉特殊字符：[java language$]
	fmt.Printf("去掉特殊字符：[%s]\n", strings.TrimLeft("#java language$", "#$"))
}
