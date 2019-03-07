package _struct

import "fmt"

type st1 struct {
	f1 int
}

func Show()  {
	type st2 struct {
		f2 int
		f3 string
	}

	fmt.Println(st1{1}, st2{2, "st2-f3"})

	st := st2{3, "st2-f3-1"}

	fmt.Println(st.f3)

	st.f3 = "st2-f3-2"

	fmt.Println(st.f3)

	stp := &st

	// 注意，这里其实隐藏了`*`，其实还可以写为`(*stp).f3 = "st2-f3-3"`
	stp.f3 = "st2-f3-3"

	fmt.Println(st.f3)
}