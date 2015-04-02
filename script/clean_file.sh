#!/bin/sh

for((i=0;i<12;i++))
do
	cd /data$i/download
	find . -ctime +20 -exec rm {} \;
	ls 
done

