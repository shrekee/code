#!/usr/bin/env python
# coding=utf-8

def login():
    global count
    if count > 2:
        print("您已三次输入密码错误...请5分钟后重试")
        return 1
    name = input("yourname...")
    password = input("password...")
    if name not in pa.keys(): 
        print("用户名或密码有错，请重新输入..")
        count += 1
        return login()
    if pa[name] != password:
        print("用户名或密码有错，请重新输入..")
        count += 1
        return login()
    else:
        print("login successfully")
        return 0
pa = dict(tom='123')
count = 0
login()
