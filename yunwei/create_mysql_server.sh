#!/bin/bash
# File Name: create_mysql_server.sh
# Author: Liwqiang
# mail: shrekee@qq.com
# Created Time: Wed 01 Aug 2018 09:30:01 PM CST

##通过glibc 安装mysql服务器

trap  '' -k 'q'  ##捕捉键盘信号‘q’时 ，停止自动安装。



echo  '------------------------'
echo  '----glibc-绿色安装包----'
echo  '默认路径/usr/local/mysql'
echo  '--'
echo  '--'
echo  ''

##倒计时10秒的程序
for(i=1;i<=10;i++);do
	trap ' ' -k 'q'   ##捕捉信号




##获取客户需求函数
get_info() {
	while [ -z $input ]; do
		read -p"$1" input
	done
}


while 
