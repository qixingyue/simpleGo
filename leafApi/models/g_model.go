package models

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"net"
	"strings"
	"time"
)

var TTL_MINUTE = 60
var TTL_HOUR = 3600
var TTL_DAY = 86400
var TTL_WEEK = 604800
var TTL_YEAR = 31536000
var TTL_FOREVER = 1892160000 //60YEAR

var ip string
var pool = newPool()

type ItemFilter struct{}

// 重写生成连接池方法
func newPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:   3600,
		MaxActive: 32, // max number of connections
		Dial: func() (redis.Conn, error) {
			c, err := redis.DialTimeout("tcp", "127.0.0.1:6379", 0, 1*time.Second, 1*time.Second)
			if err != nil {
				fmt.Printf("create pool failed.....")
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

func init() {
	addrs, _ := net.InterfaceAddrs()
	for _, addr := range addrs {
		if strings.HasPrefix(addr.String(), "10.29.8") {
			ip = addr.String()[0:10]
		}
	}
}

//如果不存在会把当前元素添加到set当中
func IsAddInSet(setName string, checkValue string, ttl int) bool {
	c := pool.Get()
	defer c.Close()
	v, _ := c.Do("TTL", setName)
	ttl_now, _ := v.(int64)
	fmt.Printf("setName : %s , checkValue :  %s, ttl :  %d , ttl_now : %d \n ", setName, checkValue, ttl, ttl_now)
	if ttl_now < 0 {
		c.Do("SADD", setName, "default_key")
		redis.Bool(c.Do("EXPIRE", setName, ttl))
	}
	if ok, _ := redis.Bool(c.Do("SADD", setName, checkValue)); ok {
		return false
	}
	return true
}

func IsInSet(setName string, checkValue string) bool {
	c := pool.Get()
	defer c.Close()
	if ok, _ := redis.Bool(c.Do("SISMEMBER", setName, checkValue)); ok {
		return true
	} else {
		return false
	}
}

func SetMembers(setName string) []string {
	c := pool.Get()
	defer c.Close()
	d, e := redis.Strings(c.Do("SMEMBERS", setName))
	if nil != e {
		return []string{""}
	} else {
		return d
	}
}

func WriteKV(k string, v string) (bool, error) {
	c := pool.Get()
	defer c.Close()
	return redis.Bool(c.Do("SET", k, v))
}

func HSet(n string, k string, v string) (bool, error) {
	c := pool.Get()
	defer c.Close()
	return redis.Bool(c.Do("HSET", n, k, v))
}

func HGet(n string, k string) (string, error) {
	c := pool.Get()
	defer c.Close()
	return redis.String(c.Do("HGET", n, k))
}

func HLen(n string) (int, error) {
	fmt.Printf("%s\n", n)
	c := pool.Get()
	defer c.Close()
	return redis.Int(c.Do("LLEN", n))
}
