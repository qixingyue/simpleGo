package libs

import (
	"fmt"
	"log"
	"os"
	"time"
)

var manageEmail = "xingyue@staff.sina.com.cn"
var errorLogger = logInit("/var/log/golang.error.log")
var runLogDir = "/data0/shareGo/logs/leafScanDownloader/"
var runLogger *log.Logger
var rdayNum int
var rlogFile *os.File

func init() {
	rdayNum = currentDayNum()
	logFileName := fmt.Sprintf("%s/log_%d_%s.log", runLogDir, rdayNum, os.Args[1])
	runLogger = logInit(logFileName)
}

func currentDayNum() (dayNum int) {
	y, m, d := time.Now().Date()
	dayNum = int(y)*10000 + int(m)*100 + int(d)
	return dayNum
}

func reinitRLog() {
	rlogFile.Close()
	rdayNum = currentDayNum()
	logFileName := fmt.Sprintf("%s/log_%d_%s.log", runLogDir, rdayNum, os.Args[1])
	runLogger = logInit(logFileName)
}

func logInit(fileName string) *log.Logger {
	rlogFile, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0)
	if err != nil {
		os.Exit(1)
	}
	_innerLogger := log.New(rlogFile, "\n", log.Ldate|log.Ltime)
	//_innerLogger := log.New(rlogFile, "\n", log.Ldate|log.Ltime|log.Lshortfile)
	return _innerLogger
}

func E_L(err error, msg string, needEmail bool) {
	errorLogger.Print(err)
	errorLogger.Println(msg)
	if needEmail {
		SendAlarmEmail(manageEmail, msg, msg)
	}
}

func R_L(msg string, needEmail bool) {
	//fmt.Println(msg)
	runLogger.Printf(msg)

	if currentDayNum() != rdayNum {
		rdayNum = currentDayNum()
		reinitRLog()
	}
	if needEmail {
		//SendAlarmEmail(manageEmail, msg, msg)
	}
}
