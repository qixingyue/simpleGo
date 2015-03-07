#!/bin/sh
while [ 2 -gt 1 ]
do

echo -n "QUEUE SIZE : "
echo -e "LLEN download_queue" | /data0/redis2817/bin/redis-cli 
echo -n "SET SIZE : "
echo -e "SCARD ttl_set" | /data0/redis2817/bin/redis-cli 
#echo -e "lrange download_queue 0 -1"  | /data0/redis2817/bin/redis-cli 

sleep 1
done
