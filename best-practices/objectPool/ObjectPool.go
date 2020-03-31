package objectPool

import (
	"fmt"
	"time"
)

type MySQL struct {
}

type MySQLPool struct {
	MaxConn    int
	ObjectPool chan MySQL
}

func (pool *MySQLPool) New(mySQL MySQL) {
	pool.ObjectPool = make(chan MySQL, 10)

	select {
	case pool.ObjectPool <- mySQL:
		fmt.Println("加入对象池成功")
	case <-time.After(time.Second * 1):
		fmt.Println("加入对象池超时")
	}
}

func (pool *MySQLPool) Get() MySQL {
	return <-pool.ObjectPool
}

func (pool *MySQLPool) Release(mySQL MySQL) {
	select {
	case pool.ObjectPool <- mySQL:
		fmt.Println("释放MySQL对象成功")
	case <-time.After(time.Second * 1):
		fmt.Println("释放MySQL对象超时")
	}
}
