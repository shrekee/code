#!/usr/bin/python
# coding:utf-8
from multiprocessing import Process, Queue,Pool
import multiprocessing
import os, time, random

# 写数据进程执行的代码:
def write(q):
    for value in ['A', 'B', 'C']:
        print 'Put %s to queue...' % value
        q.put(value)
        time.sleep(random.random())

# 读数据进程执行的代码:
def read(q):
    while True:
        print "subprocess -- reading..."
        if not q.empty():
            value = q.get(True)
            print 'Get %s from queue.' % value
            time.sleep(random.random())
        else:
            break

#if __name__=='__main__':
#    # 父进程创建Queue，并传给各个子进程：
#    q = Queue()
#    pw = Process(target=write, args=(q,))
#    pr = Process(target=read, args=(q,))
#    # 启动子进程pw，写入:
#    pw.start()
#    # 等待pw结束:
#    pw.join()
#    # 启动子进程pr，读取:
#    pr.start()
#    pr.join()
#    # pr进程里是死循环，无法等待其结束，只能强行终止:
#    print
#    print '所有数据都写入并且读完'
#    print
#    print '所有数据都写入并且读完'
#if __name__=='__main__':
#    # 父进程创建Queue，并传给各个子进程：
#    q = Queue()
#    p = Pool()
#    pw = p.apply_async(write,args=(q,))
#    pr = p.apply_async(read,args=(q,))
#    p.close()
#    p.join()
if __name__=='__main__':
    manager = multiprocessing.Manager()
    # 父进程创建Queue，并传给各个子进程：
    q = manager.Queue()
    p = Pool()
    pw = p.apply_async(write,args=(q,))
    time.sleep(0.5)
    pr = p.apply_async(read,args=(q,))
    p.close()
    p.join()
    print
    print '所有数据都写入并且读完'
