#!/bin/bash
# File Name: clean_environment.sh
# Author: lsing
# mail: shrekee@qq.com
# Created Time: Sat 28 Jul 2018 08:43:08 AM CST

##定义一个函数，用于获取当前系统下所有的ip地址，然后echo输出；
##通过定义变量var=$(get_host_ip), 可以获取本函数的输出值。
get_host_ip() {                                                                                                                  
	ip a | sed -nr 's/.*inet (([0-9]+\.){3}[0-9]{1,3}).*/\1/p'| while read i;do
	{
		echo $i
	}
done
}




##配置yum环境
#配置一个本地的yum源
create_local_yum() {
	check_iso_on
	
	if [ ! -f /etc/yum.repos.d/local.repo ]; then
	echo `menu` > /etc/yum.repos.d/local.repo
	[ $? -ne 0 ] && "创建本地yum源失败，请检查本地环境"
	fi
}

##自配置一个yum源，，并打印出来。；然后通过调用本函数重定向，来产生一个本地的yum源文件
menu() {
	cat<<-EOF
[local]
name=local-server
enabled=1
gpgcheck=0
baseurl=file:///mnt
EOF
}
##检测系统光盘/dev/sr0是否已经挂载？

check_iso_on() {
	df -h | grep '/dev/sr[0-9]'
	if [ $? -ne 0 ] ;then
		mount /dev/sr0 /mnt 
		[ $? -ne 0 ] && echo "挂载失败，请检查光盘情况。5秒后退出。" && sleep 5 && exit  
	fi
}


##自动关闭防火墙和selinux

turn_off_firewall() {
	service iptables stop
	setenforce 0
}

##定义主函数
main() {

turn_off_firewall  ##关闭系统防火墙和selinux
count=0;j=0
ip_group=`get_host_ip`  ##获取系统所有的IP地址，并保存在本变量中
declare -a array      ##定义数组，用于保存所有ip地址
for i in $ip_group;do   ##遍历ip地址，并保存到数组中   变量j可用保存ip的个数
	array[$j]=$i
	let j++
done
read -p "您一共有${j}个ip地址，都在本行以下显示。"
echo ${array[@]}
echo "请选择：2-${j}之间的一个数字作为服务器地址呢？" cho
let cho--
name_num=`echo ${array[$cho]} |cut -d. -f4`
name_num=`echo "server${name_num}.itcast.cc"`
#echo $name_num
#定义主机名
hostname $name_num
}


##调用主函数
main

