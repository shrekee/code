#!/bin/bash


# 通过find命令找出空的文件，然后通过tee命令把文件名保存在empty.txt文件，然后再删除
find ./ -type f  -empty | tee -a empty.txt | xargs rm -f
cat empty.txt | wc -l  ##把已经删除的空文件打印出来，并统计个数
rm -f empty.txt    ##把创建出来中间文件再次删除
sync  #把文件缓冲内容同步到文件系统
