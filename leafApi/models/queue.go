package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var queueSetName = "running_queue"

type QueueModel struct {
	QueueName    string
	CheckSetName string
	RealQueue    string
	SetTTL       int
}

type QueueStatusModel struct {
	Size int
}

func (this *QueueModel) Init(name string, existTTL int) {
	this.QueueName = name
	this.CheckSetName = strings.Join([]string{name, "checkSet"}, "_")
	this.RealQueue = strings.Join([]string{name, "RealQueue"}, "_")
	this.SetTTL = existTTL
}

func (this *QueueModel) GetQueues() (names []string) {
	return SetMembers(queueSetName)
}

func (this *QueueModel) GetInfo(name string) (bool, string) {
	if !IsInSet(queueSetName, name) {
		return false, "Queue not exist"
	}
	this.QueueName = name
	queueConfigName := strings.Join([]string{this.QueueName, "config"}, "_")
	ttl, err := HGet(queueConfigName, "ttl")
	CheckSetName, err := HGet(queueConfigName, "CheckSetName")
	this.CheckSetName = CheckSetName
	this.SetTTL, err = strconv.Atoi(ttl)
	this.RealQueue, err = HGet(queueConfigName, "RealQueue")
	if nil == err {
		return true, ""
	} else {
		return false, "Get Info error"
	}
}

func (this *QueueModel) CheckStatus() (bool, string) {
	l, e := HLen(this.RealQueue)
	if nil != e {
		return false, "size get error"
	}
	status := new(QueueStatusModel)
	status.Size = l
	b, e := json.Marshal(status)
	if nil != e {
		return false, "size error"
	}
	return true, string(b)
}

func (this *QueueModel) Save() (bool, error) {
	if IsAddInSet(queueSetName, this.QueueName, TTL_FOREVER) {
		fmt.Println("Alreay exist")
		return false, errors.New("Alreay exist")
	}
	//创建配置信息
	queueConfigName := strings.Join([]string{this.QueueName, "config"}, "_")
	ok, err := HSet(queueConfigName, "ttl", strconv.Itoa(this.SetTTL))
	ok, err = HSet(queueConfigName, "CheckSetName", this.CheckSetName)
	ok, err = HSet(queueConfigName, "RealQueue", this.RealQueue)
	if ok {
		return true, nil
	} else {
		fmt.Printf("\n%#v\n", err)
		return false, errors.New("Create Failed ...")
	}
}
