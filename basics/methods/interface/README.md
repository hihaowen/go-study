# 接口
Go的接口在方法一节已经说过，三原则：不主动、不拒绝、很负责
不主动：你愿意用就用，不用也不勉强
不拒绝：你要用的当然也可以
很负责：你用了我，就要遵循我的一些原则，一些必要的方法一定要有，而且要求的方法参数、返回值类型等必须按我的要求来

几点：
* 接口是存在空接口的，即`interface{}`，它类似于PHP中的`stdClass`，可以随便定义值如
```
var I interface{}
I = "i am str1"
fmt.Println(I)

// 获取I的值还可以
str1, isSet := I.(string)
fmt.Println(str1, isSet)
```

* 接口、底层类型值的零值是nil
```text
var I2 interface{}
if I2 == nil {
    fmt.Println("interface is nil")
}
```
* 如果按照接口的规范，指针获取的隐式添加将不存在，必须手动添加上`&`(具体看代码)

## 作用
* 对行为进行抽象

## 用法

```text
type 接口名称 interface{
	方法名称(参数 ...) 返回值类型
}

// 具体见代码

```