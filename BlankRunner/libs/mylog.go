package libs

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"time"
)

var NumCPU int
var manageEmail = "xingyue@staff.sina.com.cn"
var errorLogger = logInit("/var/log/golang.error.log")
var runLogDir = "/tmp"
var runLogger *log.Logger
var rdayNum int
var rlogFile *os.File
var runnerId int

func init() {
	NumCPU = runtime.NumCPU()
	if 2 <= len(os.Args) {
		runnerId, _ = strconv.Atoi(os.Args[1])
	} else {
		runnerId = 9999
	}
	rdayNum = 0
}

func ConfigManageEmail(email string) {
	manageEmail = email
}

func ConfigLogDir(logdir string) {
	runLogDir = logdir
	reinitRLog()
}

func ErrorLogFile(errorFile string) {
	errorLogger = logInit(errorFile)
}

func currentDayNum() (dayNum int) {
	y, m, d := time.Now().Date()
	dayNum = int(y)*10000 + int(m)*100 + int(d)
	return dayNum
}

func reinitRLog() {
	rlogFile.Close()
	rdayNum = currentDayNum()
	logFileName := fmt.Sprintf("%s/log_%d_%d.log", runLogDir, rdayNum, runnerId)
	runLogger = logInit(logFileName)
}

func logInit(fileName string) *log.Logger {
	rlogFile, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0)
	if err != nil {
		os.Exit(1)
	}
	_innerLogger := log.New(rlogFile, "\n", log.Ldate|log.Ltime)
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
	if currentDayNum() != rdayNum {
		rdayNum = currentDayNum()
		reinitRLog()
	}
	runLogger.Printf(msg)
	if needEmail {
		//SendAlarmEmail(manageEmail, msg, msg)
	}
}

func LogPrintf(format string, args ...interface{}) {
	message := fmt.Sprintf(format, args)
	R_L(message, false)
}
