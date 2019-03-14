package _switch

import (
	"fmt"
	"testing"
)

func TestSwitch(t *testing.T) {
	var sw = 3

	switch sw {
	case 1, 2:
		fmt.Println("是1或者2")
	case len([]int{1, 2, 3}):
		fmt.Println("是3位数")
	default:
		fmt.Println("不知道了")
	}

	// 或者
	switch sw -= 1; sw {
	case 1, 2:
		fmt.Println("是1或者2")
	}

	// 又或者 可以当if...else用
	switch {
	case sw == 2:
		fmt.Println("sw = 2")
	case sw > 2:
		fmt.Println("sw > 2")
	default:
		fmt.Println("sw < 2")
	}
}
