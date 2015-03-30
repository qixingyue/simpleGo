package libs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

var callback_url = "http://10.44.3.23:8080/download_callback.php"

func torf(r bool) string {
	if r {
		return "T"
	} else {
		return "F"
	}
}

type RealRunner struct {
	UniqueId   string
	Jsonstring string
	Url        string
	InsertTime string
}

type DownloadItem struct {
	Url          string
	AimMd5       string
	DownloadPath string
	RsyncPath    string
}

func (this *RealRunner) QueueName() string {
	return "download_RealQueue"
}

func (this *RealRunner) QueueSetName() string {
	return "download_checkSet"
}

func (this *RealRunner) Init(jsonstring string, uniqueId string, insertTime string) {
	this.UniqueId = uniqueId
	this.Jsonstring = jsonstring
	this.InsertTime = insertTime
}

func (this *RealRunner) RealDoHandler() (bool, string) {
	LogPrintf("\n=============== NEW ITEM =================\n%s", this.Jsonstring)
	//下边的代码是真正的处理流程
	di := new(DownloadItem)
	e := json.Unmarshal([]byte(this.Jsonstring), &di)
	if e != nil {
		return false, "data parse Error"
	}
	this.Url = di.Url
	res, message := this.downloadFile(di.Url, di.DownloadPath, di.AimMd5, di.RsyncPath, this.UniqueId)

	reportFileName := fmt.Sprintf("/data0/shareGo/logs/DownloadTask/run_%d.log", currentDayNum())
	reportMessage := fmt.Sprintf("%s\t%s\t%s\t%s\t%s\t%s\n", this.UniqueId, di.DownloadPath, this.InsertTime, torf(res), ToUserString(time.Now()), timeStringDiff(this.InsertTime, ToUserString(time.Now())))
	ReportFile(reportFileName, reportMessage)

	return res, message
}

func (this *RealRunner) downloadFile(download_url string, downloadFile string, aimMd5 string, rsyncPath string, uniqueId string) (bool, string) {
	dirName := filepath.Dir(downloadFile)
	fileName := downloadFile[len(dirName)+1:]
	LogPrintf("download dir : %s \n", dirName)
	LogPrintf("Download Url: %s \n", download_url)
	LogPrintf("Download FileName : %s \n", fileName)
	LogPrintf("Download AimMd5: %s \n", aimMd5)
	downloadCmd := exec.Command("/data0/myget012/bin/mytget", "-n", strconv.Itoa(NumCPU), "-d", dirName, "-f", fileName, download_url)
	downloadCmd.Stdin = strings.NewReader("")
	var out bytes.Buffer
	for try_times := 0; try_times < 3; try_times++ {
		downloadCmd.Stdout = &out
		err := downloadCmd.Run()
		if err != nil {
			R_L("Redis download error", false)
			continue
		}
		if this.checkMd5(downloadFile, aimMd5) {
			R_L("Download OK ....", false)
			this.downloadOK(uniqueId, rsyncPath)
			return true, ""
		}
	}
	cmdOutString := string(out.Bytes())
	this.downloadFailed(uniqueId, cmdOutString)
	return false, "download file error ..."
}

func (this *RealRunner) downloadOK(uniqueId string, rsyncPath string) {
	LogPrintf("\nOK %s \n", this.Url)
	data_string := fmt.Sprintf("appkey=vdisk&unique_id=%s&status=ok&fpath=%s", uniqueId, rsyncPath)
	_ = RequestPost(callback_url, data_string)
}

func (this *RealRunner) downloadFailed(uniqueId string, cmdOutString string) {
	LogPrintf("\nFailed : %s \n errorMessage: %s \n", this.Url, cmdOutString)
	m := url.Values{}
	m.Add("appkey", "vdisk")
	m.Add("unique_id", uniqueId)
	m.Add("status", "failed")
	m.Add("fpath", "")
	m.Add("errorMessage", cmdOutString)
	_ = RequestPost(callback_url, m.Encode())
	RemoveUniqueId(uniqueId, this)
}

func (this *RealRunner) checkMd5(fileName string, aimMd5 string) bool {
	cmd := exec.Command("md5sum", fileName)
	cmd.Stdin = strings.NewReader("")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		R_L("Check md5 error ....", false)
	}
	resultString := string(out.Bytes())
	resultArr := strings.Split(resultString, " ")
	return strings.EqualFold(resultArr[0], aimMd5)
}
