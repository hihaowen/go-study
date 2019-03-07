# 数组
* 跟Java、C没啥区别的常规数组，作用相同，但是还是有点区别的，C语言的数组的赋值变量其实是指向数组的指针，变量本身不存数组，而GO的赋值变量则是完全Copy数组
* key必须为数字


## 作用
存放各种类型的值

## 表现形式

```
    // 形式1
	var arr1 [3]int
	arr1[0] = 1
	arr1[1] = 2
	arr1[2] = 3

	// 形式2
	var arr2 = [3]string{"s1", "s2", "s3"}

	// 形式3 省略数量，让编译器来统计
	var arr3 = [...]float32{1, 2.1, 3.2}

```