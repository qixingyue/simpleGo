package models

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type QuickDownloadPre struct {
	Url string
	//以下的数据是需要添加进去的，传进来的原始数据不包含此数据
	DownloadPath string
	HttpPath     string
}

type QuickDownloadReply struct {
	Uniqueid string
	HttpPath string
}

// Jsondata中需要包含Url
func (this *ItemFilter) QuickDownload(callObj *TaskModel, infoWritoRedis *TaskWriteToRedis) (bool, string) {
	httpSubPath, dirPath := quickConfirmDownloadPath()
	fullPath := fmt.Sprintf("%s/%s", dirPath, infoWritoRedis.Uniqueid)
	httpPath := fmt.Sprintf("http://10.29.8.25/%s/%s", httpSubPath, infoWritoRedis.Uniqueid)
	dp := new(QuickDownloadPre)
	e := json.Unmarshal([]byte(infoWritoRedis.Jsondata), &dp)
	if nil != e {
		return false, "Json Unshal error ...."
	}
	dp.DownloadPath = fullPath
	dp.HttpPath = httpPath
	infoWritoRedis.ReplaceJsonData(dp)
	callObj.SetAdditionJson(&QuickDownloadReply{infoWritoRedis.Uniqueid, httpPath})
	return true, ""
}

func quickConfirmDownloadPath() (string, string) {
	now_time := time.Now()
	year, month, day := now_time.Date()
	httpSubPath := fmt.Sprintf("%04d%02d%02d", year, month, day)
	dir_name := fmt.Sprintf("/data0/webdownload/%s", httpSubPath)
	if !isDirExists(dir_name) {
		os.Mkdir(dir_name, 0777)
	}
	return httpSubPath, dir_name
}
