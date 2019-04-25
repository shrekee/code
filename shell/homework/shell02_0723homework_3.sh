#!/bin/bash

read -p"please input a ip address. for example,10.1.1.2.." ip
ping -c1 $ip  &> /dev/null
#[ $? -eq 0 ] && echo "Server $ip is up!" || echo "Server $ip is down!"

if [ $? -eq 0 ];then
	echo "Server $ip is up" | mail -s "ServerHost is up/down" root@localhost
	echo "Server $ip is up" | mail -s "ServerHost is up/down" mail01@localhost
else
	echo "Server $ip is down" | mail -s "ServerHost is up/down" root@localhost
	echo "Server $ip is down" | mail -s "ServerHost is up/down" mail01@localhost
fi	

