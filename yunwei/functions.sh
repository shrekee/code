#!/bin/bash
# File Name: functions.sh
# Author: Liwqiang
# mail: shrekee@qq.com
# Created Time: Thu 23 Aug 2018 07:03:06 PM CST

function input_nginx_dir() {
	read -p"请输入您已经解压缩后的nginx源包路径" nginx_dir
		cd $nginx_dir 
			[ $? -ne 0 ] && echo "路径输入错误，请重新输入" && input_nginx_dir
		}
	
input_nginx_dir
