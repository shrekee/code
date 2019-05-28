#!/usr/bin/python
## coding:utf-8

#import sys
#sys.path.append("..")
#import main
#
#print("main.sqrt(72) is %f"%main.square_root(72))

import logging, mymodule

logging.basicConfig()
logger = logging.getLogger("MyApp")

logger.setLevel(logging.DEBUG)
logger.info("Start my app")
logger = logging.getLogger("Another App")
logger.info("another.. my app")
logger.info("Start my app")
logger = logging.getLogger("MyApp")
logger.info("another.. my app")
logger.info("Start my app")

logger = logging.getLogger("Another App")

try:
    mymodule.doIt()
except Exception as e:
    logger.exception("There was a problem")

logger.info("Ending my app")
