#!/bin/bash
# File Name: check_ip_regular.sh
# Author: Liwqiang
# mail: shrekee@qq.com
# Created Time: Sun 29 Jul 2018 07:55:47 PM CST

##检查ip是否符合规范？

check_ip_regular() {
	read -p"请输入一个规范的ip地址。" ip

	echo $ip | grep -rP '^([1-9]\d*\.){3}[1-9]\d*$'
	[ $? -ne 0 ] && echo "您输入的ip不符合规范" && exit 1

	echo $ip | grep -rP '^([1-9]\d*\.){3}[1-9]\d*$'|tr '.' ' '|while read a b c d; do
		if ((a>=0 && a<=255 && b>=0 && b<=255 && c>=0 && c<=255 && d>=0 && d<=255));then
			echo '您输入的ip地址符合规范'
		else
			echo "您输入的ip地址不符合规范"
			exit 2
		fi
		done
}

#调用函数

check_ip_regular
