#!/bin/bash
# File Name: sum_code_row.sh
# Author: Liwqiang
# mail: shrekee@qq.com
# Created Time: Sun 05 Aug 2018 06:43:19 PM CST
# Desc: 统计某些目录下的代码总量，注释总量，空行总量等，并输出有意思的内容。。

#统计一个文件夹中的所有代码的之和
#即：统计下自己在学习中总共付出多少汗水？
#最后计算下，有效代码的数量 代码总量的75%

echo '`````````````````````````````````````````'
echo '`````````````````````````````````````````'
echo '此脚本用来统计某个目录下的脚本的有效行数，步骤如下：'
echo '基本策略是： 分别统计所有以#开头的行；空行；以及代码行，然后打印在屏幕'
echo '然后计算有效代码行,计算公式为: 有效代码总行=代码总行*75%'
echo '`````````````````````````````````````````'
echo '`````````````````````````````````````````'

read -p'请输入您想统计目录的路径，默认为统计当前目录'  path
#处理下用户输入的路径，如果末尾带‘/’，则自动删除最后的‘/’
path=`echo $path |sed -r 's/(.+)\/$/\1/g'`  
echo '`````````````````````````````````````````'
filename=`ls $path`  #统计当前目录下，有多少脚本？并保存在变量filename中
j=0
for i in $filename;do
	#如何是文件或者是有效的符号链接则统计代码行数
	if [ -f "$path/$i" -o -L "$path/$i" -a -e "$path/$i" ]; then
		let valid_num+=`sed -r '/^$|^#/d' $path/$i |wc -l`
		let null_num+=`sed  '/^$/!d' $path/$i |wc -l`
		let comment_num+=`sed -r '/^#/!d' $path/$i |wc -l`
	#如果是目录，则自动保存在数组中
	elif [ -d "$path/$i" ]; then
		dir_array[j]="$path/$i"
		let j++
#		echo "$path/$i"
	#既不是文件，也不是目录的，则放在else_array数组中
	else
		else_array[k]="$path/$i"
#		echo "$path/$i"
		let k++
	fi
done
echo -e "目录:\t${dir_array[@]}"
echo -e "其他:\t${else_array[@]}"
echo '`````````````````````````````````````````'
echo '统计结果如下'
echo '`````````````````````````````````````````'
let valid_code_row=$valid_num*75/100
let mix_code_comment_num=$comment_num+$valid_num
echo -e "完全空行总量是:\t${null_num}"
echo -e "完全注释行总量是:\t$comment_num"
echo -e "注释行+代码总量是:\t$mix_code_comment_num"
echo -e "代码总量是:\t$valid_num"
echo -e "有效代码总量是:\t$valid_code_row"


