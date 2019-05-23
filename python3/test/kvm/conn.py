#!/usr/bin/python3
# coding: utf-8


import libvirt

conn = libvirt.open("qemu:///system")

names = conn.listDefinedDomains()
print(names)


