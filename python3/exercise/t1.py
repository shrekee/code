#!/usr/local/bin/python3

x = 3;
z = 'global'

def f1():
	x = 'x from f1'
	y = 'y from f1 '
	print(x,y)
	def f2():
		print(z)
	f2()

f1()
print("1"*10)
