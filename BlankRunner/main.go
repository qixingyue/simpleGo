package main

import (
	. "DownloadTaskRunner/libs"
)

func main() {
	ConfigManageEmail("xingyue@staff.sina.com.cn")
	ConfigLogDir("/data0/shareGo/logs/DownloadTask/")
	Run(new(RealRunner), -1)
}
