# 函数
Go里面的函数有以下特点：
* 1、当值又当参数; 
*    1、函数闭包可以直接当作值赋给变量，变量可以直接进行调用;
*    2、函数可以作为参数类型
* 2、闭包函数可以直接操作闭包函数体以外的变量 `应对复杂的场景?`

## 作用
* 1、闭包简化操作; 简单的函数就不用再单独找个地方写了
* 2、应对复杂的场景; 上面的特点`2`可以应对复杂场景
* 3、将问题抽象化; 类似于面向对象接口的概念

## 表现形式

### 当值又当参数

* 值 (闭包)

```$xslt
getStrArrLen := func(strArr []string) int {
    return len(strArr)
}
```

* 值 (函数返回)

```$xslt
func plus() func(int, int) int {
    return func(a, b int) int {
        return a + b
    }
}
```

* 值 (函数返回、调用内部变量,函数)

```
func incr() func() int {
    base := 0
    return func() int {
        base ++
        return base
    }
}
```

* 参数

```$xslt
func sendSMS(gateway func(int, string) bool) {
    gateway(13888888888, "sms content")
}
```