#!/bin/sh

bee pack
rm /data0/shareGo/bin/leafApi.tar.gz 
mv leafApi.tar.gz /data0/shareGo/bin/leafApi.tar.gz
cd /data0/shareGo/bin/
rm -rf controllers routers conf leafApi
tar zxvf leafApi.tar.gz
/data0/shareGo/bin/leaftool.sh restart_api
rm d.sh
