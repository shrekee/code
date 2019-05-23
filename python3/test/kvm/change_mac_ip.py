#!/usr/bin/python3
# coding:utf-8

import sys
from xml.dom.minidom import parseString

def main():
    print(sys.argv)
    target = sys.argv[1]
    print(sys.argv[1])
    number = int(sys.argv[2])
    print(sys.argv[2])

    xml = open(target, 'r').read()
    doc = parseString(xml)

    for host in doc.getElementsByTagName('name'):
        print(dir(host))
        print('hello')
        ip = host.getAttribute('ip')
        parts = ip.split('.')
        parts[-1] = str(int(parts[-1]) + number)


if __name__ == '__main__':
    main()
