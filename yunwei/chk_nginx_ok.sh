#!/bin/bash
# File Name: chk_nginx_ok.sh
# Author: Liwqiang
# mail: shrekee@qq.com
# Created Time: Thu 23 Aug 2018 04:42:27 PM CST
curl -I 127.0.0.1 &> /dev/null
[ $? -ne 0 ] && service keepalived stop
