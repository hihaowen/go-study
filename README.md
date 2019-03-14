# Go语言知识点整理

## [基本数据结构](basics)

### 1、int、string、bool、复数、float 

### 2、[指针](basics/moretypes/pointer)

### 3、[结构体](basics/moretypes/struct)

## 条件控制

### 1、[if](basics/flowcontrol/if)
### 2、[switch](basics/flowcontrol/switch)
### 3、[defer](basics/flowcontrol/defer)

## [for循环](basics/flowcontrol/for)

## 集合
### 1、[数组](basics/moretypes/array)
### 2、[切片](basics/moretypes/slice)
### 3、[map](basics/moretypes/map)

## [函数](basics/moretypes/function)
### 1、当值又当参数
### 2、闭包

## 方法和接口
### 1、[方法](methods/method)
### 2、[接口](methods/interface)

## 并发(goroutine)
### 1、[go程](concurrency/goroutine)
### 2、[信道](concurrency/channel)
### 3、[select](concurrency/select)

## 问题`@todo`:
```text
1、原生类型方法到底有啥用？
2、复数是个啥
3、switch的用法用错了
4、Go不支持指针运算
    type MyFloat float64
    MyFloat += 1
5、同类型不同级别之间不能进行隐式转换
    int8 = int64
```

## 参考链接
* 图说并发 [:notebo:](https://medium.com/@trevor4e/learning-gos-concurrency-through-illustrations-8c4aff603b3)
