#!/usr/bin/python36
# coding=utf-8

#石头、剪子、布的游戏
import random
import os
#sum = ((0,0),(0,1),(0,2),(1,0),(1,1),(1,2),(2,0),(2,1),(2,2))
win = (('0','2'),('1','0'),('2','1'))   #定义了赢的组合（元组）
welcome = '''欢迎体验石头、剪子、布游戏
再游戏中，您会和电脑较量..have a fun^_^
1、剪子用数字“0”代替
2、石头用数字“1”代替
3、布 用数字“2”代替
4、输入‘quit’或者‘q’退出本游戏
'''
#建立数字和手法的对应字典
dic = {'0':'剪子','1':'石头','2':'布'}
def getnum():
    client = input("请输入一个数字[0-2]...")
    if client in ['quit','q']:
        os._exit(0)
#判断输入是否为字符[0-2]，否则重新输入
    if client not in ['1','2','0']:
        return getnum()
    return client
while True:
    print(welcome)
    rand_ = str(random.randint(0,2))
#   print(rand_)
    client = getnum()
    print(dic[rand_],dic[client])
#    print(client)
    if rand_ == client:
        print("\033[44;37m您输入的是：%s--电脑输入的是：%s...平局\033[0m"%(dic[client],dic[rand_]))
        continue
    if (rand_,client) in win:
        print("\033[42;30m您输入的是：%s---电脑输入的是：%s...您输掉了比赛\033[0m"%(dic[client],dic[rand_]))
    else:
        print("\033[41;37m您输入的是：%s---电脑输入的是：%s...恭喜您赢得比赛\033[0m"%(dic[client],dic[rand_]))
