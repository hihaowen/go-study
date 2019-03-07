# map映射
个人觉得Map就很想数组，只不过数组要求必须是int类型的key，而map则可以是字符串了，这对于有些特定的场景比如获取相关配置信息的功能就很好用

## 作用
kay => value 值映射

## 表现形式

初始化
`可通过make创建 @todo`

```$xslt
var mysqlConf map[string]MysqlConf

or

var mysqlConf = make(map[string]MysqlConf)
```

赋值
```$xslt
mysqlConf = map[string]MysqlConf {
    "local": {"127.0.0.1", "root", "root", 3306},
    "test": {"10.88.88.1", "test", "test123", 3306},
    "product": MysqlConf{"10.88.88.2", "prod", "prod123", 8888},
}
```
