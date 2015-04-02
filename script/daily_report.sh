#!/bin/sh

mail_echo (){ 
	echo "$*%13"
}

taskProcessStatus=$(ps aux | grep DownloadTaskRunner | grep -v "grep")
apiStatus=$(ps aux | grep leafApi | grep -v "grep")
diskStatus=$(df -h)

send_time=$(date +"%Y-%m-%d %H:%M:%S")

echo "DAILY REPORT $send_time" >> mail.content
echo "BODYABOVE" >>mail.content

echo "
<style>
div{
	margin-top:10px;
}
</style>
" >> mail.content

echo "<div>taskProcessStatus: <span style='color:red'> $taskProcessStatus </span></div>" >> mail.content
echo "<div>apiStatus :  <span style='color:red'>$apiStatus</span></div>" >> mail.content
echo "<div>disk status :  <span style='color:red'>$diskStatus</span></div>" >> mail.content

mail_content=$(cat mail.content)
date=$(date +"%Y-%m-%d")
curl -s -T mail.content  --connect-timeout 3 "http://172.16.30.169/qsend.php?id=007&to=xingyue@staff.sina.com.cn"

rm mail.content
