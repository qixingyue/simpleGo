package main

import (
	"os"
	"leafApi/redisLib"
	"fmt"
	"encoding/json"
	"math/rand"
	"time"
)

type Message struct {
	Url string
	Uniqueid string
}

func confirmDownloadPath(uniqueId string) (string,int) {
	now_time := time.Now()
	year,month,day := now_time.Date()
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	diskHash := r.Intn(12)
	rsyncSubPath  := fmt.Sprintf("%d%d%d",year,month,day)
	dir_name := fmt.Sprintf("/data%d/download/%s",diskHash,rsyncSubPath)
	if !isDirExists(dir_name) {
		os.Mkdir(dir_name,0777)
	}
	return rsyncSubPath,diskHash
}

func isDirExists(path string) bool {
    fi, err := os.Stat(path)
    if err != nil {
        return os.IsExist(err)
    }else{
        return fi.IsDir()
    }
    return false
}

func main(){
	uniqueId := "309988999"
	rsyncPath,diskHash := confirmDownloadPath(uniqueId)
	fmt.Printf("Hash : %d , subdir : %s \n" , diskHash, rsyncPath)
	fmt.Printf("\n================================================\n")
	msg := new(Message)
	msg.Url = "http://weibo.com"
	msg.Uniqueid = "xxxxxxxxxxxxxyyyyyyyyyx"
	b,e := json.Marshal(msg)
	if e != nil {
		fmt.Printf("Error")
	}
	fmt.Printf(string(b))
	redisLib.Lpush("QNAME",string(b))
	rbytes := redisLib.Lpop("QNAME")
	msg2 := new(Message)
	e = json.Unmarshal([]byte(rbytes), &msg2)
	fmt.Printf("%#v",msg2)
}
