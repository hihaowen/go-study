package use_redis

import (
	"github.com/bmizerany/assert"
	"github.com/gomodule/redigo/redis"
	"log"
	"testing"
	"time"
)

func TestRedis(t *testing.T) {
	// connect
	Connect(RedisConf{"127.0.0.1:6379", ""})

	key := "testKey1"
	value := "testVal1"

	// set
	Set(key, value)

	// get
	getValue, err := Get(key)
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, value, string(getValue))
}

type RedisConf struct {
	Addr string
	Pwd  string
}

var redisConf RedisConf

var pool *redis.Pool

func Connect(conf RedisConf) {
	redisConf = conf
	NewRedis()
}

func Pool() *redis.Pool {
	return pool
}

func NewRedis() {
	pool = &redis.Pool{
		MaxIdle:     50,
		MaxActive:   500,
		IdleTimeout: 300 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", redisConf.Addr)
			if err != nil {
				return nil, err
			}
			if redisConf.Pwd != "" {
				if _, err := c.Do("AUTH", redisConf.Pwd); err != nil {
					c.Close()
					log.Fatal("Redis Auth Error: " + err.Error())
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

func SetEx(key string, seconds, data interface{}) (reply interface{}, err error) {
	conn := Pool().Get()
	defer conn.Close()
	return conn.Do("SETEX", key, seconds, data)
}

func Get(key string) ([]byte, error) {
	conn := Pool().Get()
	defer conn.Close()
	return redis.Bytes(conn.Do("GET", key))
}

func GetInt(key string) (int, error) {
	conn := Pool().Get()
	defer conn.Close()
	return redis.Int(conn.Do("GET", key))
}

func Set(key string, value interface{}) (reply interface{}, err error) {
	conn := Pool().Get()
	defer conn.Close()
	return conn.Do("SET", key, value)
}

func Delete(key string) error {
	conn := Pool().Get()
	defer conn.Close()
	_, err := conn.Do("DEL", key)
	return err
}
