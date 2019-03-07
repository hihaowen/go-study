# 基础

* 1、基本数据类型
```text
整型
    // 不带负号的、类似于mysql中的unsigned
    uint8(byte) 0 ~ 1<<8 - 1
    uint16 0 ~ 1<<16 - 1
    uint32 0 ~ 1<<32 - 1
    uint64 0 ~ 1<<64 - 1
    
    // 带负号的，其实就是对半分
    int8 - (1<<8) / 2 ~ (1<<8) / 2
    int16 - (1<<16) / 2 ~ (1<<16) / 2
    int32 - (1<<32) / 2 ~ (1<<32) / 2
    int64 - (1<<64) / 2 ~ (1<<64) / 2
    
    # int 根据系统而定，如果是32位的系统，int = uint32，如果是64位，int = uint64
    
浮点 

    float32(单精度)
    float64(双精度) 

    具体可参考 https://zh.wikipedia.org/wiki/IEEE_754
    
布尔
    bool 取值为 true、false
字符串
    string 占用大小不固定
复数
    这个下来可以研究下 `@todo`
    
零型 (nil)
    // 原生变量类型初始化值不为nil
    var n1 int // 0
    var n2 float32 // 0
    var n3 string // 空字符串
    var n4 bool // false

    // 自定义变量类型初始化之后的值为nil
    var n5 []int // nil
    var n6 map[string]int // nil
```
* 2、变量、常量、赋值运算

变量：写法跟其它语言有点区别，类型是放在右边，标准的写法是
```
var v1 int
v1 = 1
```

以上可简写为
```text
v1 := 1
```

当然还可以双赋值
```text
v1, v2 = 1, 2
```

常量：用关键词`const`来标记，取值范围为int、float、string、boolean
特点是：
1、只能用`const`标记，不能用`:=`
2、常量不能修改

```text
const C1 = "c2"
```

* 3、包相关
```text
1、每个文件头部都有 `package` 字样，类似于命名空间的概念，但不同的是，该`package`的值只当前文件所在的目录名，并不是和其它语言如java、PHP等类似的目录结构
```

* 4、函数
```text
1、如果文件内func名首字母大写的话，说明这个函数是对外可调用的
2、编写格式为: 
func 必填的函数名称(非必填参数 参数类型) 非必填的返回值类型 {
    
}

例子：
func func1(arg1 int) string {
    return "123"
}

3、返回值可以是多个
func func1(arg1 int) string, int {
    return "123", 123
}

```