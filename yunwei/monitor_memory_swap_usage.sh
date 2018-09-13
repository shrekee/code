#!/bin/bash
# File Name: monitor_memory_swap_usage.sh
# Author: Liwqiang
# mail: shrekee@qq.com
# Created Time: Sun 29 Jul 2018 07:45:37 PM CST

##监控系统内存和交换分区使用情况

monitor_memory_swap_usage() {
	
	free -m | sed -nr '/^Mem|^Swap/p'| tr -s ' '|cut -d' ' -f1,2,3| while read a b c ;do
			let k=c*100/b
			echo "$a已经使用了$k%"
	done

}

monitor_memory_swap_usage
