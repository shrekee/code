#!/usr/local/bin/python3
#_*_ coding:utf-8 _*_

import socket
#创建socket对象
s = socket.socket(socket.AF_INET,socket.SOCK_STREAM)
#获取本机主机名
host = socket.gethostname()
port = 12345
#绑定端口
s.bind((host,port))
#等待客户端连接
s.listen(5)
while True:
	c,addr = s.accept()
	print("the address is %s and port is %s"%addr)
	#print(addr)
	msg = '欢迎访问菜鸟教程!\r\n'
	c.send(msg.encode('utf-8'))
	c.close()
