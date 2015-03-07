package main

import (
	. "DownloadTaskRunner/libs"
	"fmt"
)

func main() {

	//处理panic异常
	defer func() {
		fmt.Println("c")
		if err := recover(); err != nil {
			fmt.Println(err) // 这里的err其实就是panic传入的内容，55
		}
		fmt.Println("d")
	}()

	ConfigManageEmail("xingyue@staff.sina.com.cn")
	ConfigLogDir("/data0/shareGo/logs/DownloadTask/")
	Run(new(RealRunner), -1)
}
