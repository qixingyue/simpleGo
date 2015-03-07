package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	lm "leafApi/models"
	"leafScanDownloader/libs"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var NumCPU int

func init() {
	NumCPU = runtime.NumCPU()
}

// 重写生成连接池方法
func newPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:   30,
		MaxActive: 2 * NumCPU, // max number of connections
		Dial: func() (redis.Conn, error) {
			c, err := redis.DialTimeout("tcp", "127.0.0.1:6379", 0, 1*time.Second, 1*time.Second)
			if err != nil {
				libs.E_L(err, "create pool failed ...", true)
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

func downloadFile(url string, downloadFile string, aimMd5 string, rsyncPath string, uniqueId string) {
	dirName := filepath.Dir(downloadFile)
	fileName := downloadFile[len(dirName)+1:]
	libs.R_L(fmt.Sprintf("Download Dir : %s \n", dirName), false)
	libs.R_L(fmt.Sprintf("Download Url: %s \n", url), false)
	libs.R_L(fmt.Sprintf("Download FileName : %s \n", fileName), false)
	fmt.Printf("Download FileName : %s \n", fileName)
	libs.R_L(fmt.Sprintf("Download AimMd5: %s \n", aimMd5), false)
	//downloadCmd := exec.Command("/data0/myget012/bin/mytget", "-n", "8", "-d", dirName, "-f", fileName, url)
	downloadCmd := exec.Command("/data0/myget012/bin/mytget", "-n", strconv.Itoa(NumCPU), "-d", dirName, "-f", fileName, url)
	downloadCmd.Stdin = strings.NewReader("")

	for try_times := 0; try_times < 3; try_times++ {
		var out bytes.Buffer
		downloadCmd.Stdout = &out
		err := downloadCmd.Run()
		if err != nil {
			libs.R_L("Redis download error", false)
			continue
		}
		if checkMd5(downloadFile, aimMd5) {
			libs.R_L("Download OK ....", false)
			downloadOK(uniqueId, rsyncPath)
			return
		}
		//break
	}
	downloadFailed(uniqueId)
}

func downloadOK(uniqueId string, rsyncPath string) {
	libs.R_L(fmt.Sprintf("Post result %s  %s OK ", uniqueId, rsyncPath), false)
	url := "http://10.44.3.23:8080/download_callback.php"
	data_string := fmt.Sprintf("appkey=vdisk&unique_id=%s&status=ok&fpath=%s", uniqueId, rsyncPath)
	libs.R_L(data_string, false)
	postReturnString := libs.RequestPost(url, data_string)
	libs.R_L(postReturnString, false)
}

func downloadFailed(uniqueId string) {
	libs.R_L(fmt.Sprintf("Post result %s failed ", uniqueId), false)
	url := "http://10.44.3.23:8080/download_callback.php"
	data_string := fmt.Sprintf("appkey=vdisk&unique_id=%s&status=failed&fpath=%s", uniqueId, "")
	postReturnString := libs.RequestPost(url, data_string)
	libs.R_L(postReturnString, false)

	//remove key from ttl_set
	c := pool.Get()
	defer c.Close()
	c.Do("SREM", "ttl_set", uniqueId)

}

func checkMd5(fileName string, aimMd5 string) bool {
	cmd := exec.Command("md5sum", fileName)
	cmd.Stdin = strings.NewReader("")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		libs.R_L("Check md5 error ....", false)
	}
	resultString := string(out.Bytes())
	resultArr := strings.Split(resultString, " ")
	libs.R_L(fmt.Sprintf("download file md5:  %s \n compare Md5 :  %s \n", resultArr[0], aimMd5), false)
	return strings.EqualFold(resultArr[0], aimMd5)
}

func getItem() (*lm.DownloadInfoWriteRedis, error) {
	c := pool.Get()
	defer c.Close()
	downloadModel := new(lm.DownloadInfoWriteRedis)
	if nil != checkRedisCon(c) {
		c = pool.Get()
	}
	rstring, err := redis.String(c.Do("LPOP", "download_queue"))
	if nil != err {
		return downloadModel, err
	}
	if 0 == len(rstring) {
		libs.R_L(fmt.Sprintf(" empty queue  %s sleep 3 second  \n", rstring), false)
	}
	libs.R_L(fmt.Sprintf(" download item %s \n", rstring), false)
	e := json.Unmarshal([]byte(rstring), &downloadModel)
	return downloadModel, e
}

func begin(dataChan chan *lm.DownloadInfoWriteRedis) {
	for {
		select {
		case downloadModel := <-dataChan:
			fmt.Printf("\ndownload url: %#v\n", downloadModel)
			downloadFile(downloadModel.Url, downloadModel.DownloadPath, downloadModel.AimMd5, downloadModel.RsyncPath, downloadModel.Uniqueid)
		default:
			time.Sleep(1000 * time.Millisecond)
		}
	}
}

func main() {
	libs.R_L("Download started ....", true)
	runtime.GOMAXPROCS(NumCPU)
	runtime.Gosched()

	dataChan := make(chan *lm.DownloadInfoWriteRedis, NumCPU)
	for i := 0; i < NumCPU; i++ {
		go begin(dataChan)
	}

	//flag := true
	//var item *lm.DownloadInfoWriteRedis
	//var err error
	item, err := getItem()

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
}
