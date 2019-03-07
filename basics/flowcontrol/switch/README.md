# switch
跟传统的switch一样，不过还是有些区别：
* case之后自动增加了break
* 条件语句前面可以增加一条简单的执行，这一点跟`if`一样
* case条件语句可以是常量，也可以是值运算

## 作用
没啥可说的

## 表现形式

### 常规的
```$xslt
    switch status {
	case 1:
		statusDesc = "check"
	case 2:
		statusDesc = "done"
	case 3:
		statusDesc = "delete"
	default:
	    statusDesc = "no known"
	}
```

### case为值运算
```$xslt
    switch status {
	case getStatusDesc("check"):
		statusDesc = "check"
	default:
		statusDesc = "no known"
	}
```
