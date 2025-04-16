package math

import (
	"testing"
)

// 测试文件必须以 _test.go 结尾
// 需要 import "testing" 包
// 测试函数格式：func TestXxx(*testing.T)
// 使用 go test 命令执行测试
func TestAdd(t *testing.T) {
	result := Add(2, 3)
	if result != 5 {
		t.Errorf("Expected 5, got %d", result)
	}
}

func TestDivide(t *testing.T) {
	t.Run("normal division", func(t *testing.T) {
		result, err := Divide(6, 2)
		if err != nil {
			t.Fatal("unexpected error:", err)
		}
		if result != 3 {
			t.Error("Expected 3, got", result)
		}
	})

	t.Run("divide by zero", func(t *testing.T) {
		_, err := Divide(5, 0)
		if err == nil {
			t.Fatal("expected error but got nil")
		}
	})
}
