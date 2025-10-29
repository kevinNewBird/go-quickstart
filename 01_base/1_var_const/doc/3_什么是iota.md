# 1.概念
iota：特殊常量，可以认为是一个可以被编译器修改的常量

# iota常量的定义方式
格式：const 常量名 [type] = iota [算法]
```go
package main

import (
	"fmt"
)

func main() {
	// 1.ioto的定义方式，使用这种方式后值会有编译器进行自动的递增
	const (
		INFO1 = iota
		INFO2
		INFO3 = "ha" // iota内部仍然会增加计数器
		INFO4        // iota + 1
		INFO5 = iota
	)

	// 0 1 ha 3 4
	fmt.Println(INFO1, INFO2, INFO3, INFO4, INFO5)
}
```

# 3.注意事项
- a.如果中断了iota那么必须显式的恢复, 后续会自动递增
- b.自增类型默认是int类型
- c.iota能简化const类型的定义
- d.每次出现const关键字的时候，iota初始化为零