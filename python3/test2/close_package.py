#!/usr/bin/python
# coding:utf-8

def testFun():
    temp = [lambda x : i * x for i in range(4)]
    return temp

for everylambda in testFun():
    print( everylambda(2) )

#print(testFun())
