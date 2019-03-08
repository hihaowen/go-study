# 方法

首先声明：方法和函数的区别，函数是面向过程的func，而方法在遇到Go之前我的理解是面向对象（也就是类）中的func、而Go并没有明确的面向对象的概念，
但是面向对象的思想是有的，比如通常类中有属性和方法两项、再结合interface或abstract进行进一步抽象化。这些其实在Go中也存在的，Go里有类型（type）的概念，
我说的类型指的是关键字`type`，而不是原生类型int、string等。类型我认为是包括三部分：struct、interface、带有`type`关键字的原生类型(int, string)，而方法
可以在上面除`interface`类型的定义下进行创建，这样通过组合的形式，达到了面向对象的要求，这可能来源于Go作者对技术的深刻理解吧

* 三种类型：interface、struct、带有关键字`type`的原生类型(int、string .etc)

* 注意：这里为什么不能为interface实现方法，因为Go里面的interface的特点是：不勉强、不拒绝、但是很负责，这种松耦合的用法，使得代码更灵活

* 我们知道面向对象里面的属性在整个类中是引用的关系，在一个类中随时随地可以动态的改变该属性的值，而在Go中同样能办到，它是通过引用指针来改变属性的值的，
Go中指针是通过`&`符号来获得值的指针，然后通过`*`来操作指针从而改变指针指向的值，举个例子:
```text
intCal := IntCal{1}
myIntCal := intCal // 其实这里的的`intCal`应该是`&intCal`，Go隐式的帮我们加上了
myIntCal.plus(2)
myIntCal.plus(2)
fmt.Printf("now myInt = %d", myIntCal.myInt)

type IntCal struct {
	myInt int
}
// 通过这里的`*IntCal`，Go会自动的帮我们拿到指针，也就是上面的`myIntCal := intCal`，`&`被隐式的添加
// 注意：这里跟`函数`完全不同，函数的指针获取必须是显式的，不然会报错
func (i *IntCal) plus(otherInt int) { 
	i.myInt += otherInt
}
```

## 作用
1、丰富代码行为 (类似实现了面向对象)

## 用法

### 基本写法
```
func (实现除`interface`类型之一的变量名称 实现除`interface`类型之一的变量类型) 方法名(参数 ...) 返回值类型(可以没有返回值) {
    // 具体实现
}
```
### 类型方法
* struct类型方法
```text
// 先定义类型
type ST struct{
    field1, field2 int
}

// 创建ST类型方法
func (st ST) StFunc(arg1 int, arg2 string) {
    // 具体实现
}
```
* 带有关键字`type`的原生类型方法

这里有个疑问，原生类型方法到底有啥用？@todo

```text
// 先定义类型
type MyFuncInt int

// 创建ST类型的方法
func (myInt MyFuncInt) StFunc(arg1 int, arg2 string) {
    // 具体实现
}
```
