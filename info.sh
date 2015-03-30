#!/bin/sh

f_name=$1
full_path=$( echo $line | awk '{print $2}')
curl -s -d "fname=$f_name" "http://10.44.3.23:8080/vdisk_info.php" > tmp.info
title=$(cat tmp.info | jq ".name")
mime_type=$(cat tmp.info | jq ".data.mime_type")
description=$(cat tmp.info | jq ".data.description")
size=$(cat tmp.info | jq ".data.size")
echo "<p>$title $mime_type $size</p>" 
echo "<p>$description</p>"
echo "<p>$full_path</p>" 
echo "</div>" 
echo $f_name
rm tmp.info
