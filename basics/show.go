package basics

import (
	"fmt"
)

const C1 = 1

func Show()  {
	const C2 = 2

	fmt.Println(C1)
	fmt.Println(C2)

	// int
	var i int
	i = (1<<64) / 2 - 1

	fmt.Println(i)

	// float
	var f float64
	f = -9999999999999999

	fmt.Println(f)

	// 原生变量类型初始化值不为nil
	var n1 int // 0
	var n2 float32 // 0
	var n3 string // 空字符串
	var n4 bool // false

	// 自定义变量类型初始化之后的值为nil
	var n5 []int // nil
	var n6 map[string]int // nil

	if n5 == nil && n6 == nil {
		fmt.Println("n5、n6 为 nil")
	}

	fmt.Println(n1, n2, n3, n4, n5, n6)
}


