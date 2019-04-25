#!/bin/bash


for i in {0..20};do
	ping -c1 "10.1.1.${i}"
	[ $? -eq 0 ] && echo "10.1.1.${i}" >> /root/code/ip_exist.txt || echo "10.1.1.${i}" >> /root/code/ip_noexist.txt
done


	
