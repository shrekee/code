#!/usr/local/bin/python3

#简单的实现网络聊天室功能：

import socket
import threading

sock = socket.socket(socket.AF_INET,socket.SOCK_STREAM)
sock.bind(('localhost',5550))

sock.listen(5)

print('Server',socket.gethostbyname('localhost'),'listening...')

mydict = dict()
mylist = list()

#把whatToSay传递给除了exceptNum的所有人

def tellOthers(exceptNum, whatToSay):
	for i in mylist:
		if i.fileno
