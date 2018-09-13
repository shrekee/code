#!/bin/bash
# File Name: sum_code_row.sh
# Author: Liwqiang
# mail: shrekee@qq.com
# Created Time: Sun 05 Aug 2018 06:43:19 PM CST
# Desc: 统计某个目录下的代码总量（以递归的方式遍历目录查询），注释总量，空行总量等，并输出有意思的内容。。

#统计一个文件夹中的所有代码的之和
#即：统计下自己在学习中总共付出多少汗水？
#最后计算下，有效代码的数量 代码总量的75%
function auto_sum() {
#		path=$1
#		path=`echo $path |sed -r 's/(.+)\/$/\1/g'`  
#		echo '`````````````````````````````````````````'
#		filename=`ls $1`  #统计当前目录下，有多少脚本？并保存在变量filename中
#		
#		tip:在递归查询中，不能使用变量保存别的变量，比如，不能令filename=`ls $1`,因为shell语言，默认自定义的变量都是全局变量，
#  导致递归失败，除非使用关键字“local” 来声明变量，才可以。切记！！切记！
#

			for i in `ls $1`;do
				#如果是目录，则自动保存在数组中,然后递归
				if [ -d "$1/$i" ]; then
					dir_array[j]="$1/$i"   #把子目录中的目录保存在数组中
					let j++
					echo "$1/$i"
					auto_sum "$1/$i"
				#如果是文件或者是有效的符号链接则统计代码行数
				elif [ -f "$1/$i" -o -L "$1/$i" -a -e "$1/$i" ]; then
					let valid_num+=`sed -r '/^$|^#/d' $1/$i |wc -l`   #统计有用的代码行
					let null_num+=`sed  '/^$/!d' $1/$i |wc -l`		#统计空行
					let comment_num+=`sed -r '/^#/!d' $1/$i |wc -l`	#统计注释行（每行的开头是#，则为注释行）
				#既不是文件，也不是目录的，则放在else_array数组中
				else
					else_array[k]="$1/$i"
					let k++
				fi
			done
}
function main(){
		echo '````````````````````````````````````````````````````````````````````'
		echo '````````````````````````````````````````````````````````````````````'
		echo '此脚本用来统计某个目录下的脚本的有效代码行数(递归的方式)，步骤如下：'
		echo -e "第一步： 分别统计所有以#开头的注释行\t空行\t代码行"
		echo '第二步： 计算有效代码行。其计算公式为: 有效代码行=代码总行*75%'
		echo -e "最后分别输出:  注释行\t空行\t代码行"
		echo '````````````````````````````````````````````````````````````````````'
		echo '````````````````````````````````````````````````````````````````````'

		read -p'请输入您想统计目录的路径，默认为统计当前目录'  path
		#处理下用户输入的路径，如果末尾带‘/’，则自动删除最后的‘/’
		path=`echo $path |sed -r 's/(.+)\/$/\1/g'`  
		echo '````````````````````````````````````````````````````````````````````'
		auto_sum $path
		echo -e "目录:\t${dir_array[@]}"
		echo -e "其他:\t${else_array[@]}"
		echo '````````````````````````````````````````````````````````````````````'
		echo '统计结果如下'
		echo '````````````````````````````````````````````````````````````````````'
		let valid_code_row=$valid_num*75/100
		let mix_code_comment_num=$comment_num+$valid_num
		echo -e "空行.......总量:\t${null_num}"
		echo -e "注释行.....总量:\t$comment_num"
		echo -e "注释行+代码总量:\t$mix_code_comment_num"
		echo -e "代码.......总量:\t$valid_num"
		echo -e "有效代码...总量:\t$valid_code_row"
}

main
