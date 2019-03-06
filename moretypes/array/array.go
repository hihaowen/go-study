package array

import "fmt"

func showArray() bool {
	// 形式1
	var arr1 [3]int
	arr1[0] = 1
	arr1[1] = 2
	arr1[2] = 3

	// 形式2
	var arr2 = [3]string{"s1", "s2", "s3"}

	// 形式3 省略数量，让编译器来统计
	var arr3 = [...]float32{1, 2.1, 3.2}

	fmt.Println(arr1, arr2, arr3)

	return true
}
