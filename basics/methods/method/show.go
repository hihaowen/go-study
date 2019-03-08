package method

import "fmt"

func Show()  {
	// struct类型方法
	person := Person{"George", 28}
	person.describe()

	// 原生类型方法
	age := MyAge(30)
	age.AgeDescribe()

	intCal := IntCal{1}
	myIntCal := intCal
	myIntCal.plus(2)
	myIntCal.plus(2)
	fmt.Printf("now myInt = %d", myIntCal.myInt)
}

type IntCal struct {
	myInt int
}
func (i *IntCal) plus(otherInt int) {
	i.myInt += otherInt
}

// 先定义结构体类型
type Person struct{
	name string
	age int
}

// 创建ST类型方法
func (person Person) describe() {
	fmt.Printf("我叫 %s，我今年 %d 岁\n", person.name, person.age)
}

// 先定义原生类型
type MyAge uint8

// 再定义方法
func (age MyAge) AgeDescribe() {
	fmt.Printf("我今年 %d 岁\n", age)
}
