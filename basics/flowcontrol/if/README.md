# if
最基本的条件控制语句，有以下特点：
* 没有`(`和`)`; 这点跟其它如Java、C、PHP完全不一样，显得更简洁，当然你要是硬加也没问题
* 可以在条件判断语句前加一个简单操作，但改操作的结果并不影响条件流程以外的部分

## 作用
没啥好说的

## 表现形式

### 普通

```$xslt
if 条件表达语句 {
    // 成立则执行
}
```

### 特别

```$xslt
if 简单语句; 条件表达语句 {
    // 成立则执行
}
```
