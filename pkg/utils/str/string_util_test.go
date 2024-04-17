package str

import (
	"testing"
)

// TestContainsString 测试ContainsString
func TestContainsString(t *testing.T) {
	var (
		target   = "a"
		arr      = []string{"a", "b", "c", "d", "e"}
		contains bool
		index    int
	)

	contains, index = ContainsString(arr, target)

	if !contains || index != 0 {
		t.Errorf("ContainsString arr: %s, target: %s failed, contains: %v, index: %d", arr, target, contains, index)
	}
	target = "f"

	contains, index = ContainsString(arr, target)

	if contains || index != -1 {
		t.Errorf("ContainsString arr: %s, target: %s failed, contains: %v, index: %d", arr, target, contains, index)
	}
}
