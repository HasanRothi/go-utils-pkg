package arrayutils

import (
	"fmt"
	"testing"
)

func TestSliceContains(t *testing.T) {
	result := SliceContains([]string{"a", "b", "c"}, "b")
	expected := true
	if result != expected {
		t.Errorf("SliceContains() = %v; want %v", result, expected)
	}
}

func TestDynamicArray(t *testing.T) {
	fmt.Println("ok")
}
