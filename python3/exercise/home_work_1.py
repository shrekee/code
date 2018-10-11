#!/usr/bin/env python
# coding=utf-8

__author__ = 'liwenqiang'

with open("/root/code/python3/login.py") as f1:
    f2 = open("/tmp/11110000.py",'w')
    f2.write(f1.read())
    f1.close()
    f2.close()
