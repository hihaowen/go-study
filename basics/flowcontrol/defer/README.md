# defer

* 在函数体执行完之后再执行; 类似于PHP中goto，goto语句的特定形式
* 属于栈、遵循后进先出顺序

## 作用
比如说垃圾回收之类的需求

## 表现形式

### 普通 && defer栈
```$xslt
    defer gc()

	for i := 1; i < 6; i ++ {
		defer fmt.Printf("我是第 %d 个被插入的\n", i)
	}
```