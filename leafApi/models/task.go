package models

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"reflect"
	"strings"
	"time"
)

type TaskModel struct {
	Uniqueid     string
	Jsondata     string
	QueueName    string
	CheckSetName string
	SetTTL       int
	ReturnJson   string
	RealQueue    string
}

type TaskWriteToRedis struct {
	Uniqueid   string
	Jsondata   string
	InsertTime string
}

func (this *TaskModel) Init(Uniqueid string, QueueName string, Jsondata string) (bool, string) {
	this.Uniqueid = Uniqueid
	this.QueueName = QueueName
	this.Jsondata = Jsondata
	return this.initQueueInfo()
}

func (this *TaskModel) initQueueInfo() (bool, string) {
	qm := new(QueueModel)
	if ok, message := qm.GetInfo(this.QueueName); ok {
		this.CheckSetName = qm.CheckSetName
		this.SetTTL = qm.SetTTL
		this.RealQueue = qm.RealQueue
		return true, ""
	} else {
		return false, message
	}
}

func (this *TaskModel) checkExist() bool {
	c := pool.Get()
	defer c.Close()
	v, _ := c.Do("TTL", this.CheckSetName)
	ttl, _ := v.(int64)
	if ttl < 0 {
		c.Do("SADD", this.CheckSetName, "default_check_set_value")
		redis.Bool(c.Do("EXPIRE", this.CheckSetName, this.SetTTL))
	}
	if ok, _ := redis.Bool(c.Do("SADD", this.CheckSetName, this.Uniqueid)); ok {
		return ok
	}
	return false
}

func (this *TaskModel) WriteRedis() (bool, string) {
	diwr := new(TaskWriteToRedis)
	diwr.Uniqueid = this.Uniqueid
	diwr.Jsondata = this.Jsondata
	now_time := time.Now()
	diwr.InsertTime = now_time.Format("2006-01-02 15:04:05")
	ifilter := new(ItemFilter)
	ifilterType := reflect.TypeOf(ifilter)

	methodName := strings.Join([]string{strings.ToUpper(this.QueueName[0:1]), this.QueueName[1:]}, "")
	if _, ok := ifilterType.MethodByName(methodName); ok {
		vfilterValue := reflect.ValueOf(ifilter)
		m := vfilterValue.MethodByName(methodName)
		vfilterParams := []reflect.Value{reflect.ValueOf(this), reflect.ValueOf(diwr)}
		v := m.Call(vfilterParams)
		if 2 != len(v) {
			return false, "reflect Call error...."
		}
		if !v[0].Bool() {
			return false, v[1].String()
		}
	}

	b, e := json.Marshal(diwr)
	fmt.Printf(string(b))
	fmt.Printf("%s\n", this.RealQueue)

	if nil != e {
		return false, "json to string error"
	} else {
		c := pool.Get()
		defer c.Close()
		if this.checkExist() {
			if ok, _ := redis.Bool(c.Do("LPUSH", this.RealQueue, string(b))); ok {
			} else {
				return false, "push error"
			}
		} else {
			return false, "repeat record"
		}
		return true, ""
	}
}

func (this *TaskModel) SetAdditionJson(d interface{}) {
	b, _ := json.Marshal(d)
	this.ReturnJson = string(b)
}

func (this *TaskWriteToRedis) ReplaceJsonData(d interface{}) {
	b, _ := json.Marshal(d)
	this.Jsondata = string(b)
}
