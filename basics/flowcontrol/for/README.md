# for
for循环是Go里面仅有的内置循环函数，它不像其它语言那样还有其它的比如while、do .. while、foreach之类的用法，但是Go里面的for完全可以胜任这些工作

* 三部分组成： 初始化语句（第一次循环时运行）、条件表达式：在每次循环前检查、后置语句：在每次循环完执行
* 中断循环用break

## 作用
没啥好说的

## 表现形式

### for
```$xslt
    // for
	for i := 0; i < 6; i ++ {
		fmt.Println(i)
	}
```

or

```$xslt
    i = 0
	for i < 5 {
		i++
		fmt.Println(i)
	}
```

### while
```
	// while 循环
	i := 0
	for {
		if i > 4 {
			break
		}

		i ++

		fmt.Println(i)
	}
```

### foreach
```
	// foreach 有k
	arr1 := []int{1, 2, 3}
	for k := range arr1 {
		fmt.Println(k)
	}

	arr2 := map[string]int{
		"k1" : 1,
		"k2" : 2,
	}

	// foreach 有k有v
	for k, v := range arr2 {
		fmt.Printf("k = %s, v = %d\n", k, v)
	}

	// foreach 有v
	for _, v := range arr2 {
		fmt.Printf("v = %d\n", v)
	}
```
