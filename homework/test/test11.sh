#!/bin/bash

######方法一：  while负责外层   until负责内层

i=1 #外层循环
while ((i<=5));do
	j=1 #内层循环
	until ((j>i));do
		echo -n $j   ## 内层循环中，控制内层的打印次数<=外层数
		let j++ ##内层自加1
	done
	echo      #打印换行符
	let i++   ##外层自加1
done

##方法二  until负责外层，while负责内层

i=1
until ((i>5));do
	j=1
	while ((j<=i));do
		echo -n $j
		let j++
	done
	echo
	let i++
done
	

##方法三      使用for列表方式  ++
for i in {1..5};do
	j=1
	while ((j<=i));do
		echo -n $j
		let j++
	done
	echo 
done

echo   "以上，是用三种不同方法实现的"
#################################

# DEMO2
echo DEMO2

for ((i=5;i>=1;i--));do
	for ((j=5;j>=i;j--));do
		echo -n $j
	done
	echo
done

	
###########################
#以下为
#打印第三种图形的方法

for ((i=1;i<=5;i++));do
	for ((j=5;j>=i;j--));do
		echo -n $j
	done
	echo
done

###############
echo  "打印九九乘法表"
echo

for ((i=1;i<=9;i++));do
	echo
	for ((j=1;j<=i;j++));do
		let k=i*j
		echo -n "$j * $i = $k "
		#echo -n "\t"
	done
	echo
done
i=1
while ((i<=9));do
	j=1
	echo
	until ((j>i));do
		let k=i*j
		echo -ne "$j*$i=${k}\t "
		let j++
	done
	let i++
	echo
done






