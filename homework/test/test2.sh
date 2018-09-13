#!/bin/bash
rpm -aq vsftpd
[ $? -eq 0 ] && echo "已安装"  || yum install -y vsftpd &> /dev/null
