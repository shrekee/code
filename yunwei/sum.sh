#!/bin/bash
# File Name: sum.sh
# Author: Liwqiang
# mail: shrekee@qq.com
# Created Time: Mon 03 Dec 2018 10:07:25 AM CST

for i in $*;do
	let sum=sum+i
done
echo $sum
