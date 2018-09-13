#!/bin/bash
# File Name: monitor_disk_usage.sh
# Author: Liwqiang
# mail: shrekee@qq.com
# Created Time: Sun 29 Jul 2018 07:00:39 PM CST

#检测某个已经挂载磁盘分区的事情情况，如果使用量超过90%则mail给管理员

#定义一个函数，用于获取某个磁盘分区的使用情况

disk_usage() {
	if [ $# -eq 1 ]; then
		df  | tr -s " " |while read a b c d e f;do if [[ $f = "$1" ]] ;then let k=c*100/b;echo $k;fi;done
	else
		echo "usage : disk_usage '/'" && exit 1
	fi
}
#disk_usage /boot

##定义一个主函数，用于当磁盘剩余空间不足90%时，自动给管理员发邮件
main() {
	per=`disk_usage /`
	echo $per

##如何使用mail命令发邮件？需要好好学习以下
###	((per>90)) && echo "$per分区使用量已超过90%“  | mail -s "disk_90%_usage_emergent"  root@loacalhost
}
main

