package _if

import "fmt"

func Show() {
	var a = 1

	if a == 1 {
		fmt.Println("a == 1")
	}

	if a := a + 1; a > 1 {
		fmt.Println("a > 1")
	}

	// a 还是 1
	fmt.Println(a)
}
