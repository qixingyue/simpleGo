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
	//NumCPU = 1
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
	SendAlarmEmail(manageEmail, "REINIT LOG TELL YOU I AM RUNNING ...", "REINIT LOG TELL YOU I AM RUNNING ...")
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
	message = fmt.Sprintf("%s:%s", time.Now().Local(), message)
	R_L(message, false)
}

func ToUserString(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func timeStringDiff(t string, e string) string {
	tt, _ := time.Parse("2006-01-02 15:04:05", t)
	et, _ := time.Parse("2006-01-02 15:04:05", e)
	d := et.Sub(tt)
	return d.String()
}

func ReportFile(fileName string, text string) {
	f, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0x644)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.WriteString(text)
	f.Sync()
}
