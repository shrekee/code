#!/bin/bash


#自动安装nfs，并且启动nfs服务 ；
#把服务器/data 作为服务目录，以只读方式共享给私网（内网）所有人；
/bin/rpm -q rpcbind &> /dev/null
if [ $? -ne 0 ];then
	yum install -y rpcbind 
	if [ $? -ne 0 ] ;then
   	echo "yum源安装rpcbind软件失败，请检查yum环境" 
	 exit 1
	 fi
fi
/bin/rpm -q nfs-ntils &> /dev/null
if [ $? -ne 0 ] ;then
	yum install -y nfs-utils    # &> /dev/null
	if [ $? -ne 0 ] ;then
		echo "yum源安装nfs软件失败，请检查yum环境"
	   	exit 2
	fi
fi
#假设共享的目录是 /data 而且默认以只读方式共享给所有用户

echo "/data  *(ro)" > /etc/exports  2> /dev/null
service rpcbind start &> /dev/null
[ $? -ne 0 ] && echo "rpcbind服务进程启动失败。。。" && exit 3 || echo "rpc服务启动  ok"
service nfs start &> /dev/null
[ $? -ne 0 ] && echo "nfs服务进程启动失败"  && exit 4   || echo "nfs 服务启动失败 "
