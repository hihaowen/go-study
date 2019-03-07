package _for

import (
	"fmt"
)

func Show() {
	// for
	for i := 0; i < 6; i ++ {
		fmt.Println(i)
	}

	i := 0
	for i < 5 {
		i++
		fmt.Println(i)
	}

	// while 循环
	i = 0
	for {
		if i > 4 {
			break
		}

		i ++

		fmt.Println(i)
	}

	// foreach 有k
	arr1 := []int{1, 2, 3}
	for k := range arr1 {
		fmt.Println(k)
	}

	arr2 := map[string]int{
		"k1" : 1,
		"k2" : 2,
	}

	// foreach 有k有v
	for k, v := range arr2 {
		fmt.Printf("k = %s, v = %d\n", k, v)
	}

	// foreach 有v
	for _, v := range arr2 {
		fmt.Printf("v = %d\n", v)
	}
}
