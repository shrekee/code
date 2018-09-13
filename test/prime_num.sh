#!/bin/bash
# File Name: prime_num.sh
# Author: Liwqiang
# mail: shrekee@qq.com
# Created Time: Sun 12 Aug 2018 09:32:05 PM CST

#求40万内的素数，
#同样的算法，python的用时：3-4秒，GCC4.4版本下的c用时：0.12秒左右，，看下shell的速度如何
#先看下计算10000内素数的时间，初步估算为时间10000的时间*40倍；

begin=`date +%s`
array[0]=2
k=1
for i in `seq 3 2 10000`;do
	sqrt=`echo "sqrt($i)"|bc`
	let sqrt++
	for j in `seq 2 $sqrt`;do
		let mid=i%j
		if [[ $mid -eq 0 ]]; then
			break
		fi
		if [[ $j -eq $sqrt ]];then
			array[k]=$i
			let k++
		fi
	done
done
end=`date +%s`
let deru=end-begin;
echo ${array[@]}
echo "用时为：$deru 秒"
echo "1万内素数的个数为： ${#array[@]} 个"
