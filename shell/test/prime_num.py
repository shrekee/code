#!/usr/local/bin/python
#coding:utf8

#导入统计时间模块
import time
#记录开始时间，用变量start保存
start = time.time()
#定义列表，用于保存素数，从2开始。
lst=[2]
for i in xrange(3,400000,2):
	for j in xrange(2,int(i**0.5)+1):
		if i%j == 0:
			break
	else:
		lst.append(i)
#记录结束时间
end = time.time()
#结束时间减去开始时间，等于deru
deru=end-start
#打印所有的素数
print(lst)
#打印用时
print("总共用时为：%f 秒"%deru)
#打印总的素数个数
print("素数的个数为：%d 个"%len(lst))











print("---------------分隔线--------------")
##以下是第二种方法求素数，，通过马哥教育的wayne老师所学，，老师吹嘘性能好，其实是垃圾的一批，相比第一种方法慢好几倍。
#prime_num=[2]
#for i in range(3,100000,2):
#	for j in prime_num:
#		if i%j == 0:
#			break
#	else:
#		prime_num.append(i)
#
#end2=time.process_time()
#deru=end2-end
#print(deru)
#print(len(prime_num))

