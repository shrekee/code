#!/bin/bash
# File Name: homework_clean_environment.sh
# Author: Liwqiang
# mail: shrekee@qq.com
# Created Time: Sat 28 Jul 2018 06:45:09 PM CST
#desc:  清理当前系统环境，包括根据当前IP自动配置主机名为（如：ip是192.168.0.88，则主机名改为server88.itcast.cc）
#自动配置可用yum源
#自动关闭防火墙和selinux

#自动修改主机名

get_host_ip() {
  ip a | sed -nr 's/.*inet (([0-9]+\.){3}[0-9]{1,3}).*/\1/p'  | while read i ;do
	echo $i
	let j++
	done
}	
sum=0;j=0
sum=`get_host_ip`
echo $sum
echo $j
#get_host_ip
