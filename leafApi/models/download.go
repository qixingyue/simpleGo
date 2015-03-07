package models

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"log"
)

type DownloadInfo struct {
	Url      string
	Uniqueid string
	AimMd5   string
}

type DownloadInfoWriteRedis struct {
	Url          string
	Uniqueid     string
	DownloadPath string
	AimMd5       string
	RsyncPath    string
}

func (this *DownloadInfo) checkExist() bool {
	c := pool.Get()
	defer c.Close()
	set_name := "ttl_set"
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

func (this *DownloadInfo) WriteRedis() (string, string) {
	queue_name := "download_queue"
	rsyncPath, dirPath, diskHash := confirmDownloadPath()
	fullPath := fmt.Sprintf("%s/%s", dirPath, this.Uniqueid)
	rsyncPath = fmt.Sprintf("%s::download_rsync_%d/%s/%s", ip, diskHash, rsyncPath, this.Uniqueid)

	diwr := new(DownloadInfoWriteRedis)
	diwr.Url = this.Url
	diwr.Uniqueid = this.Uniqueid
	diwr.DownloadPath = fullPath
	diwr.AimMd5 = this.AimMd5
	diwr.RsyncPath = rsyncPath
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
		return rsyncPath, ""
	}
}
