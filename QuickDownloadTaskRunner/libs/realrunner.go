package libs

import (
	"bytes"
	"encoding/json"
	"fmt"
	_ "net/url"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

type RealRunner struct {
	UniqueId   string
	Jsonstring string
	Url        string
}

type DownloadItem struct {
	Url          string
	DownloadPath string
	HttpPath     string
}

func (this *RealRunner) QueueName() string {
	return "QuickDownload_RealQueue"
}

func (this *RealRunner) QueueSetName() string {
	return "QuickDownload_checkSet"
}

func (this *RealRunner) Init(jsonstring string, uniqueId string) {
	this.UniqueId = uniqueId
	this.Jsonstring = jsonstring
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
	return this.downloadFile(di.Url, di.DownloadPath, di.HttpPath, this.UniqueId)
}

func (this *RealRunner) downloadFile(download_url string, downloadFile string, httpUrl string, uniqueId string) (bool, string) {
	dirName := filepath.Dir(downloadFile)
	fileName := downloadFile[len(dirName)+1:]
	LogPrintf("download dir : %s \n", dirName)
	LogPrintf("Download Url: %s \n", download_url)
	LogPrintf("Download FileName : %s \n", fileName)
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
		this.downloadOK(uniqueId, httpUrl)
		return true, ""
	}
	cmdOutString := string(out.Bytes())
	this.downloadFailed(uniqueId, cmdOutString)
	return false, "download file error ..."
}

func (this *RealRunner) downloadOK(uniqueId string, httpUrl string) {
	title := fmt.Sprintf("Download_%s_OK", uniqueId)
	body := fmt.Sprintf("Get_It_From: %s", httpUrl)
	SendAlarmEmail(manageEmail, body, title)
}

func (this *RealRunner) downloadFailed(uniqueId string, cmdOutString string) {
}
