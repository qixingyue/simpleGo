package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
)

func E_L(e error, m string, n bool) {

}

// 重写生成连接池方法
func newPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:   30,
		MaxActive: 2 * 16, // max number of connections
		Dial: func() (redis.Conn, error) {
			c, err := redis.DialTimeout("tcp", "127.0.0.1:6379", 0, 0, 0)
			if err != nil {
				E_L(err, "create pool failed ...", true)
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

func checkRedisCon(c redis.Conn) error {
	_, err := c.Do("PING")
	return err
}

// 生成连接池
var pool = newPool()

func main() {
	c := pool.Get()
	defer c.Close()
	subName := "pubsub"
	data := "this is new world ...."
	c.Do("PUBLISH", subName, data)
	fmt.Printf("OK ....")
}
