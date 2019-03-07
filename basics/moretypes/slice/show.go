package slice

import "fmt"

func ShowSlice() {

	// 普通创建
	var s1 = []int{1, 2, 3}

	fmt.Println(s1)

	// 带有长度、容量的创建，均有初始值
	var s2 = make([]int, 3, 5)

	fmt.Println(s2)

	// 切片的零值为nil
	var s3 []string
	if s3 == nil {
		fmt.Printf("切片的零值为: nil\n")
	}

	s4 := []int{1}
	s5 := make([]int, len(s4), cap(s4) * 2)

	copy(s5, s4)

	// 按照容量填充
	s5 = s5[:cap(s5)]
	fmt.Println(cap(s5))

	// 这里比较神奇的地方是，本身容量6个就可以，结果由4个扩容到了8个，难道扩容是按照倍数增加的？
	s5 = append(s5, 1, 2)

	fmt.Println(s5)
	fmt.Println(cap(s5))

	// append 第二个参数如果也是切片，则成了追加操作，需要使用 ... 语法
	s5 = append(s5, s4 ...)
	fmt.Println(s5)
	fmt.Println(cap(s5))

	s6 := s5[3:5]
	printSplice(s6)

	// s6还有3个空档，可以填充下
	s6 = s6[:cap(s6)]
	printSplice(s6)

	// 对数组也是同样的操作的
	arr1 := [...]int{1, 2, 3}
	s7 := arr1[1:]
	printSplice(s7)

	// 更改原始数组值
	s7[0] = 9527
	fmt.Println(arr1)
}

func printSplice(s []int) {
	fmt.Printf("s len: %d, cap: %d, val: %v\n", len(s), cap(s), s)
}
