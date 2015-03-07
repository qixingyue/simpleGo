#!/bin/sh

#curl -i -d "name=world&ttl=86400" "http://10.29.8.25:8088/queue/create"
#curl -i  "http://10.29.8.25:8088/queue/status?name=world"

# running_queue queuelist
# queue_name 对应hash
# check_
# real_


redisCli(){
	echo "$1" | /data0/redis2817/bin/redis-cli
}

queueList(){
	redisCli "SMEMBERS running_queue"	
}

delqueue(){
	redisCli ""	
}

queueInfo(){
	redisCli "HKEYS down_config"	
	redisCli "HVALS down_config"	
}

#queueList
queueInfo
