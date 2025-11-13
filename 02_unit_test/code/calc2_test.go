package calc

import (
	"strconv"
	"strings"
	"testing"
)

// 2.1.性能测试： Benchmark开头
func BenchmarkAdd(bb *testing.B) {
	var a, b, c int = 123, 456, 579

	// 对数据做多少轮的测试
	for i := 0; i < bb.N; i++ {
		if actual := add(a, b); actual != c {
			bb.Errorf("%d + %d, expect: %d, actual: %d", a, b, c, actual)
		}
	}
}

const numbers = 10000

// -short是因为功能单元测试有错误的用例，不指定会影响benchmark的执行
// go test -bench=".*" -short
func BenchmarkStringJoiner(bb *testing.B) {
	bb.ResetTimer()

	for i := 0; i < bb.N; i++ {
		var str string
		for j := 0; j < numbers; j++ {
			str = str + strconv.Itoa(j)
		}
	}
	bb.StopTimer()
}

func BenchmarkStringBuilder(bb *testing.B) {
	bb.ResetTimer()

	for i := 0; i < bb.N; i++ {
		var builder strings.Builder
		for j := 0; j < numbers; j++ {
			builder.WriteString(strconv.Itoa(j))
		}
		_ = builder.String()
	}
	bb.StopTimer()
}
