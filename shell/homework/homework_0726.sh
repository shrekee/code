#!/bin/bash
# File Name: homework_0726.sh
# Author: lsing
# mail: shrekee@qq.com
# Created Time: Fri 27 Jul 2018 07:43:06 PM CST

function choose( ) {
			read -p"$1" ans
			while [ -z $ans ]; do
				read -p"$1" ans
			done
			echo $ans
			}
function main( ) {
			ip=`choose 请输入您的服务器ip`
			domain_name=`choose 请输入您的域名`
			#echo "$ip,$domain_name,$data_dir"
			echo "开始yum安装appache服务...按ctl+c终止本操作。"
			rpm -q httpd &> /dev/null
			[ $? -ne 0 ] &&  yum install httpd -y &> /dev/null
			[ $? -ne 0 ] && echo "yum源安装失败，请检查yum配置文件.." && exit 1
			echo "this is test page" > "${data_dir}/index.html"
			echo "ServerName ${domain_name}:80" >> /etc/httpd/conf/httpd.conf

}
main
