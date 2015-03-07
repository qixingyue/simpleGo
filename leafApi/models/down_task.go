package models

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"syscall"
	"time"
)

type DownloadPre struct {
	Url    string
	AimMd5 string
	//以下的数据是需要添加进去的，传进来的原始数据不包含此数据
	DownloadPath string
	RsyncPath    string
}

type DownloadReply struct {
	Uniqueid  string
	RsyncPath string
}

// Jsondata中需要包含Url,AimMd5
func (this *ItemFilter) Download(callObj *TaskModel, infoWritoRedis *TaskWriteToRedis) (bool, string) {
	rsyncPath, dirPath, diskHash := confirmDownloadPath()
	fullPath := fmt.Sprintf("%s/%s", dirPath, infoWritoRedis.Uniqueid)
	rsyncPath = fmt.Sprintf("%s::download_rsync_%d/%s/%s", ip, diskHash, rsyncPath, infoWritoRedis.Uniqueid)
	dp := new(DownloadPre)
	e := json.Unmarshal([]byte(infoWritoRedis.Jsondata), &dp)
	if nil != e {
		return false, "Json Unshal error ...."
	}
	dp.DownloadPath = fullPath
	dp.RsyncPath = rsyncPath
	infoWritoRedis.ReplaceJsonData(dp)
	callObj.SetAdditionJson(&DownloadReply{infoWritoRedis.Uniqueid, rsyncPath})
	return true, ""
}

func confirmDownloadPath() (string, string, int) {
	now_time := time.Now()
	year, month, day := now_time.Date()
	var diskHash int
	for {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		diskHash = r.Intn(12)
		if isPathFree(diskHash) {
			break
		}
	}
	rsyncSubPath := fmt.Sprintf("%04d%02d%02d", year, month, day)
	dir_name := fmt.Sprintf("/data%d/download/%s", diskHash, rsyncSubPath)
	if !isDirExists(dir_name) {
		os.Mkdir(dir_name, 0777)
	}
	return rsyncSubPath, dir_name, diskHash
}

func isDirExists(path string) bool {
	fi, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	} else {
		return fi.IsDir()
	}
	return false
}

type DiskStatus struct {
	All  uint64 `json:"all"`
	Used uint64 `json:"used"`
	Free uint64 `json:"free"`
}

func DiskUsage(path string) (disk DiskStatus) {
	fs := syscall.Statfs_t{}
	err := syscall.Statfs(path, &fs)
	if err != nil {
		return
	}
	disk.All = fs.Blocks * uint64(fs.Bsize)
	disk.Free = fs.Bfree * uint64(fs.Bsize)
	disk.Used = disk.All - disk.Free
	return
}

func isPathFree(diskHash int) bool {
	path := fmt.Sprintf("/data%d", diskHash)
	ds := DiskUsage(path)
	return ds.Free >= 20971520
}
