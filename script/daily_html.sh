#!/bin/sh


# 每天5点，定时生成每天的下载文件报告

file_info(){
	f_name=$1
	full_path=$2
	curl -s -d "fname=$f_name" "http://10.44.3.23:8080/vdisk_info.php" > tmp.info
	title=$(cat tmp.info | jq ".name")
	mime_type=$(cat tmp.info | jq ".data.mime_type")
	description=$(cat tmp.info | jq ".data.description")
	size=$(cat tmp.info | jq ".data.size")
	echo "<div>"
	echo "<p>$title $mime_type $size</p>" 
	echo "<p>$description</p>"
	echo "<p>$full_path</p>" 
	echo "</div>" 
	echo $f_name
	echo "</div>"
	rm tmp.info
}



dname=$(date +"%Y%m%d" -d "-1 day")
htmlname="$dname.html"

for((m=0;m<=12;m++))
do
	for f in $( ls /data${m}/download/$dname) 
	do
		file_info $f /data${m}/download/$dname/$f >> $htmlname
	done
	break
done



