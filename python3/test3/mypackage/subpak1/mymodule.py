#!/usr/bin/python
# coding:utf-8

import logging

logger = logging.getLogger("MyModule")

def doIt():
    logger.debug("doin' stuff")

    raise TypeError("bogus type error for testing")


