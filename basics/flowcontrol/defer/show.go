package _defer

import "fmt"

func Show()  {
	defer gc()

	for i := 1; i < 6; i ++ {
		defer fmt.Printf("我是第 %d 个被插入的\n", i)
	}
}

func gc()  {
	fmt.Println("我是最后一步,GC ing ...")
}
