#!/bin/bash
# File Name: recursion_dir.sh
# Author: Liwqiang
# mail: shrekee@qq.com
# Created Time: Sun 05 Aug 2018 09:46:05 PM CST

#递归遍历目录
function list_alldir() {
for file in `ls $1`
do
	if [ -d $1"/"$file ]
	then
		echo $1"/"$file
		list_alldir $1"/"$file
#	else
#		dos2unix $1"/"$file
#	echo $1"/"$file
	fi
done
}
list_alldir $1
