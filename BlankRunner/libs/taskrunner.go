package libs

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"runtime"
	"time"
)

type TaskRunningItem struct {
	UniqueId string
	Jsondata string
}

type TaskRunner interface {
	Init(jsonString string, uniqueId string)
	RealDoHandler() (bool, string)
	QueueName() string
	QueueSetName() string
}

var taskRunner TaskRunner

// 重写生成连接池方法
func newPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:   30,
		MaxActive: 2 * NumCPU, // max number of connections
		Dial: func() (redis.Conn, error) {
			c, err := redis.DialTimeout("tcp", "127.0.0.1:6379", 0, 1*time.Second, 1*time.Second)
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

func getItem() (*TaskRunningItem, error) {
	c := pool.Get()
	defer c.Close()
	downloadModel := new(TaskRunningItem)
	if nil != checkRedisCon(c) {
		c = pool.Get()
	}
	rstring, err := redis.String(c.Do("LPOP", taskRunner.QueueName()))
	if nil != err {
		return downloadModel, err
	}
	if 0 == len(rstring) {
		R_L(fmt.Sprintf(" empty queue  %s sleep 3 second  \n", rstring), false)
	}
	e := json.Unmarshal([]byte(rstring), &downloadModel)
	return downloadModel, e
}

func begin(dataChan chan *TaskRunningItem) {
	for {
		select {
		case item := <-dataChan:
			taskRunner.Init(item.Jsondata, item.UniqueId)
			taskRunner.RealDoHandler()
		default:
			time.Sleep(1000 * time.Millisecond)
		}
	}
}

func RemoveUniqueId(id string) {
	c := pool.Get()
	defer c.Close()
	c.Do("SREM", taskRunner.QueueSetName(), id)
}

func Run(tr TaskRunner, max int) {
	taskRunner = tr
	runtime.GOMAXPROCS(NumCPU)
	runtime.Gosched()

	dataChan := make(chan *TaskRunningItem, NumCPU)
	for i := 0; i < NumCPU; i++ {
		go begin(dataChan)
	}

	item, err := getItem()

	if max == -1 {
		for {
			if nil == err {
				select {
				case dataChan <- item:
					item, err = getItem()
				default:
					time.Sleep(1000 * time.Millisecond)
				}
			} else {
				item, err = getItem()
			}
		}
	} else {
		for m := 0; m < max; m++ {
			if nil == err {
				select {
				case dataChan <- item:
					item, err = getItem()
				default:
					time.Sleep(1000 * time.Millisecond)
				}
			} else {
				item, err = getItem()
			}
		}
	}
}
