#!/bin/bash

pgrep -l vsftpd
if [ $? -eq 0  ];then 
	echo vsftpd的服务器已启动...
	echo vsftpd监听的端口：
	echo "vsftpd的进程PID是$(pgrep -l vsftpd | cut -d' ' -f1)"
fi
