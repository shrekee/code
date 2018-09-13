#!/usr/local/bin/python3.7
# coding=utf-8

import getpass
username = input("name:")
password = getpass.getpass("please input your password")
if username == 'tom' and password == '123':
    print("login sucessfully")
else:
    print("username or password is wrong")

