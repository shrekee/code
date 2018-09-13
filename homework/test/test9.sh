#!/bin/bash

#交互式判断输入的一个年份是否为闰年？
read -p"please Input a year number,,e.g: 2018" year

if [ $[year%4] -eq 0 -a $[year%100] -ne 0 -o $[year%400] -eq 0 ];then
	echo "$year是闰年"
else
	echo "$year不是闰年"
fi
