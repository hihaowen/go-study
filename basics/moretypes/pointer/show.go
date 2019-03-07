package pointer

import "fmt"

func Show()  {
	r1 := 1

	r2 := &r1 // 拿到了r1对应值的内存地址
	fmt.Println("r1的内存地址为:", r2)

	*r2 = 2 // 现在r1的值就变成了2
	fmt.Println("r1现在的值为:", r1)
	fmt.Println("r2现在的值为:", *r2)
}
