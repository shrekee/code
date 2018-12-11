#!/usr/local/bin/python3
#_*_ coding:utf-8 _*_

#导入正则模块
import re
if __name__ == "__main__":
#打开文件'/proc/meminfo' 
#	with open('/proc/meminfo') as f:
##因为打开后的文件是可迭代的,于是,我们可以使用for 来遍历文件内容
#		for i in f:
#			if re.match('MemTotal:[ ]+(\d+) kB',i):
#				tot = re.match('MemTotal:[ ]+(\d+) kB',i)
#			if re.match('MemFree:[ ]+(\d+) kB',i):
#				lef = re.match('MemFree:[ ]+(\d+) kB',i)
#				break
#		total = tot.group(1)
#		left = lef.group(1)
#		f.close()
		
#print("内存还剩余%d%%"%(int(left)*100/int(total))
	f = open('/proc/meminfo')
	for i in f:
		if re.match('MemTotal:[ ]+(\d+) kB',i):
			tot = re.match('MemTotal:[ ]+(\d+) kB',i)
		if re.match('MemFree:[ ]+(\d+) kB',i):
			lef = re.match('MemFree:[ ]+(\d+) kB',i)
			break
	f.close()
	total =int(tot.group(1))
	left = int(lef.group(1))
	#print("内存还剩余%d"%(left*100/total)
	print("内存剩余%.2f%%"%(left*100/total))
