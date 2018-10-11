#!/usr/bin/env python
# coding=utf-8

with open("/tmp/passwd",'r') as f:
    sf = f.read()
    b = open("/tmp/test1",'a')
    b.write(sf)
    b.close()



