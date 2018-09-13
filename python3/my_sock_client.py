#!/usr/local/bin/python3 
#_*_ coding:utf-8 _*_


#导入 socket sys模块
import socket
import sys

#创建 socket 对象
s = socket.socket(socket.AF_INET,socket.SOCK_STREAM)

#获取本机主机名
host = socket.gethostname()

#设置端口号
port = 12345

#连接服务,指定主机和端口
s.connect((host,port))

#接受小于1024字节的数据
msg = s.recv(1024)
s.close()
print(msg.decode('utf-8'))
