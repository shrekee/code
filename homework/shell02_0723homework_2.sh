#!/bin/bash

#name:lising
#date:2018-07-23
#decribe: homework


#作业二：输入一个路径，判断路径是否存在，而且输出是文件还是目录，
#如果是字符连接，还得输出是有效的连接还是无效的连接 

read -p"please input a path,for example: /root/code/run/test1.sh..." str
if [ -L $str ];then    #判断是否为符号链接，如果是，则把真实路径保存在变量dir中
	dir=`ls -l $str | cut -d">" -f2 | cut -d" " -f2`
	ls $dir &> /dev/null
	[ $? -eq 0 ] && echo "有效的符号链接" || echo "无效的符号链接"
	exit 0
fi
if [ ! -e $str ];then
	echo "您输入的路径不存在" && exit 1
else
	if [ -f $str ];then
		echo "此为文件"
	elif [ -d $str ];then
		echo "此为目录"
	fi
	exit 0
fi

