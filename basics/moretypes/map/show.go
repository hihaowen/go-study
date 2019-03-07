package _map

import "fmt"

type MysqlConf struct {
	host string
	user string
	pass string
	port int
}

func ShowMap() {

	var mysqlConf map[string]MysqlConf

    // or make
	// var mysqlConf = make(map[string]MysqlConf)

	// 零值为nil
	if mysqlConf == nil {
		fmt.Println("mysqlConf eq nil")
		// return
	}

	// 数据库配置
	mysqlConf = map[string]MysqlConf {
		"local": {"127.0.0.1", "root", "root", 3306},
		"test": {"10.88.88.1", "test", "test123", 3306},
		"product": {"10.88.88.2", "prod", "prod123", 8888},
	}

	// map删除
	delete(mysqlConf, "test")

	// map值修改
	mysqlConf["local"] = MysqlConf{"127.0.0.1", "root", "db_r", 7777}

	// map值的获取
	testConf, isSet := mysqlConf["test"]
	if ! isSet {
		fmt.Println("test conf is not set")
		return
	}

	fmt.Println(testConf)
}
