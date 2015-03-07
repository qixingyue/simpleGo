#!/bin/sh

go build
rm -rf /data0/shareGo/bin/leafScanDownloader
mv leafScanDownloader /data0/shareGo/bin/leafScanDownloader
/data0/shareGo/bin/leaftool.sh restart_download
