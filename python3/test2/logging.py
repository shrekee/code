#!/usr/bin/python
# coding:utf-8

import logging
def testFunc1():
    log = logging.getLogger("father")
    log.debug("1: debug:zzzz")
    print("testFunc1")

def testFunc2():
    log = logging.getLogger("child")
    log.info("2: info: aaaaa")
    print("testFunc2...")


if __name__ == "__main__":
    testFunc1()
    testFunc2()
