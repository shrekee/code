#!/bin/bash

#把etc文件中的用户名和uid截取出来
cat /etc/passwd | cut -d: -f1,3 | tr ':' " " > mulu.txt
declare -A super_usr
declare -A sys_usr
declare -A common_usr
i=0;j=0;k=0
while read user uid ;do
	if [ $uid -gt 500 ];then
		common_usr[$i]=$user
		let i++
	elif [ $uid -gt 0 ];then
		sys_usr[$j]=$user
		let j++
	elif [ $uid -eq 0 ];then
		super_usr[$k]=$user
		let k++
	fi
done < mulu.txt

echo ${super_usr[@]} >> usr.txt
echo ${sys_usr[@]} >> usr.txt
echo ${common_usr[@]} >> usr.txt

