#!/bin/bash
# File Name: chk_nginx_ok.sh
# Author: Liwqiang
# mail: shrekee@qq.com
# Created Time: Thu 23 Aug 2018 04:42:27 PM CST

#我的目的是
#定义一个资源检查与管理脚本，
curl -I 127.0.0.1 &> /dev/null
if [ $? -ne 0 ] ;then
	service nginx restart
	sleep 2
	curl -I 127.0.0.1 &> /dev/null
	if [ $? -ne 0 ];then
		service keepalived stop
	fi
fi
