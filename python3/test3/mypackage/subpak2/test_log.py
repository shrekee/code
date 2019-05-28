#!/usr/bin/python
# coding:utf-8

import logging
import os

logging.basicConfig(level=logging.INFO, format='%(asctime)s - %(name)s - %(levelname)s -%(message)s',filename='/tmp/test.log')
logger = logging.getLogger(__name__)
print os.path.dirname(__file__)

logger.info('This is a log info')
logger.debug('Debugging')
logger.warning('Warning exists')
logger.info('Finish')
