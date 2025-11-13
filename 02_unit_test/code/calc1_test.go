package calc

import "testing"

// 1.1.功能测试： Test开头
func TestAdd(t *testing.T) {
	expected := 10
	ret := add(2, 8)
	if ret != expected {
		t.Errorf("expected: %d, actual: %d", expected, ret)
	}
}

// 1.2.功能测试： 跳过耗时长的单元测试
// go test -short .
func TestAdd2(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	expected := 7
	ret := add(2, 8)
	if ret != expected {
		t.Errorf("expected: %d, actual: %d", expected, ret)
	}
}

// 1.3.功能测试：基于表格的单元测试
func TestAdd3(t *testing.T) {
	// 定义测试数据表格
	var dataset = []struct {
		a   int
		b   int
		out int
	}{
		{1, 1, 2},
		{12, 12, 24},
		{0, 0, 0},
	}

	for _, data := range dataset {
		re := add(data.a, data.b)
		if re != data.out {
			t.Errorf("expected: %d, actual: %d", data.out, re)
		}
	}
}
