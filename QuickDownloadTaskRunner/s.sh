#!/bin/sh
rm QuickDownloadTaskRunner
rm /data0/shareGo/bin/QuickDownloadTaskRunner
go build
go install
ps aux | grep QuickDownloadTaskRunner | grep -v "grep" | awk '{print $2}' | xargs kill -9 
/data0/shareGo/bin/QuickDownloadTaskRunner > /data0/shareGo/logs/QuickDownloadTaskRunner/run.log 2>&1 &
