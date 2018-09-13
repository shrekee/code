#!/bin/bash

if [ $# -ne 1  ];then
	echo "“usage:/root/code/test10.sh  hello or world"
	exit
fi
if [ $1 = 'hello' -o $1 = "world" ];then
	if [ $1 = 'hello' ];then
		echo world
	else
		echo hello
	fi
else
	echo "“usage:/root/code/test10.sh  hello or world"
	exit
fi
