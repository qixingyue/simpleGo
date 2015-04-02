#!/bin/sh

#按照类型和日期来分析下载文件
#使用方法: sh info.sh <type> <day>

info(){
	f_name=$1
	full_path=$2
	curl -s -d "fname=$f_name" "http://10.44.3.23:8080/vdisk_info.php" > tmp.info
	title=$(cat tmp.info | jq ".name")
	mime_type=$(cat tmp.info | jq ".data.mime_type")
	description=$(cat tmp.info | jq ".data.description")
	size=$(cat tmp.info | jq ".data.size")
	echo $title $full_path $size
	#echo "<div>"
	#echo "<p>$title $mime_type $size</p>" 
	#echo "<p>$description</p>"
	#echo "<p>$full_path</p>" 
	#echo "</div>" 
	#echo $f_name
	rm tmp.info
}

typelist(){
	type=$1
	d=$2
	cat run_$d.log | grep -w "T" | grep "$type" | while read line 
	do
		l_array=($line)
		f_name=${l_array[0]}
		full_path=${l_array[1]}
		info $f_name $full_path
	done
}

if [ $# -eq 2 ]
then
	typelist $1 $2
else
	echo "Help : sh info.sh <type> <daynum>"
fi

