#!/bin/bash

#name:lising
#date:2018-07-23
#decribe: homework

#作业一：判断/tmp/run目录是否存在，如果不存在就建立，
#如果存在就删除目录里所有文件
[ ! -d /tmp/run ] && $(mkdir /tmp/run)|| rm -f /tmp/run/* > /dev/null

#作业二：输入一个路径，判断路径是否存在，而且输出是文件还是目录，
#如果是字符连接，还得输出是有效的连接还是无效的连接 

