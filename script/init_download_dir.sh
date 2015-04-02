#!/bin/sh

for((i=0;i<12;i++))
do
	if [ -d "/data$i/download" ] 
	then
		echo "exists dir /data$i/download"
	else 
		mkdir "/data$i/download"
	fi
done



for((i=0;i<12;i++))
do
	echo "[download_rsync_$i]"
	echo "path = /data$i/download"	
	echo "read only = no"
	echo "uid = root"
	echo "gid = root"
	echo "hosts allow = 10.44.3.23,10.44.3.24,10.13.0.41,127.0.0.1,10.29.8.25"
	echo "list = no"
	echo "auth users=rsync_auth"
	echo "secrets file=/data0/script/.download_rsync.sect"
	echo ""
done
