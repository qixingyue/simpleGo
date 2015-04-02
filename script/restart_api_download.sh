#!/bin/sh

source /etc/profile
source /etc/bashrc
source /root/.bashrc
source /root/.bash_profile

export GOROOT=/data0/go
export GOPATH="/data0/shareGo"
export PATH="$PATH:$GOROOT/bin:$GOPATH/bin"

/data0/shareGo/bin/leaftool.sh stop_api
/data0/shareGo/bin/leaftool.sh clean_api_logs
/data0/shareGo/bin/leaftool.sh start_api

### do not need stop download 
#/data0/shareGo/bin/leaftool.sh stop_download
/data0/shareGo/bin/leaftool.sh clean_download_logs
#/data0/shareGo/bin/leaftool.sh start_download
