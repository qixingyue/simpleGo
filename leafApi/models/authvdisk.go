package models

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"log"
)

type AuthvdiskM struct {
	Uniqueid string
	Jsondata string
}

type AuthvdiskMWriteRedis struct {
	Uniqueid string
	Jsondata string
}

func (this *AuthvdiskM) checkExist() bool {
	c := pool.Get()
	defer c.Close()
	set_name := "authvdisk_ttl_set"
	v, _ := c.Do("TTL", set_name)
	ttl, _ := v.(int64)
	if ttl < 0 {
		c.Do("SADD", set_name, "ttlkey")
		redis.Bool(c.Do("EXPIRE", set_name, 60*60*24))
	}
	if ok, _ := redis.Bool(c.Do("SADD", set_name, this.Uniqueid)); ok {
		return ok
	}
	return false
}

func (this *AuthvdiskM) WriteRedis() (string, string) {
	queue_name := "authvdiskqueue"
	diwr := new(AuthvdiskMWriteRedis)
	diwr.Uniqueid = this.Uniqueid
	diwr.Jsondata = this.Jsondata

	b, e := json.Marshal(diwr)
	fmt.Printf(string(b))

	if nil != e {
		return "", "json to string error"
	} else {
		c := pool.Get()
		defer c.Close()
		if this.checkExist() {
			if ok, err := redis.Bool(c.Do("LPUSH", queue_name, string(b))); ok {
			} else {
				log.Print(err)
			}
		} else {
			return "", "repeat record"
		}
		return "ok", ""
	}
}
