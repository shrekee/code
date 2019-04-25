#!/bin/bash

#lsing
#
read -p"Input a user name" usr
cat /etc/passwd | grep -w $usr | grep -v grep &> /dev/null
[ $? -eq 0 ] && echo "该用户存在" || echo "无此用户"
