#!/bin/bash
# File Name: chk_ssh+mysql+nginx.sh
# Author: Liwqiang
# mail: shrekee@qq.com
# Created Time: Sat 18 Aug 2018 10:51:26 PM CST

#每隔一小时检查mysql，nginx和ssh的是否正常运行，写入日志：

##   sql_sock="/usr/local/mysql/mysql.sock"

#方法二：可以接受参数，并且把参数放在数组中，，，
#array=（$*）
#
#


#定义三个需要查询的对象
obj="msyql nginx sshd"
##遍历以上的对象，并把日志追加到日志（包含日期）
for i in $obj;do
	netstat -ntlp | grep $i > /dev/null
	[ $? -eq 0 ] && echo "$i is up"`date '+%F %T'` >> /tmp/chk_ssh+mysql+nginx.log || echo "$i is down "`date '+%F %T'` &>> /tmp/chk_ssh+mysql+nginx.log
done

##取出系统目前正在监听的端口号，并保存在变量stat中
stat=`netstat -tnlp | awk -F'[ ]+' '{print $4}' | sed -rn  "s#([0-9]+\.){3}[0-9]+:([0-9]+)#\2#pg" `
##遍历所有的端口，并输出
for i in $stat;do
	echo $i
done
