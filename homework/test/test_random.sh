#!/bin/bash


for i in user{1..5};do
#	echo $i
	useradd $i
	echo "$[${RANDOM}%10000+1000]" | tee -a shadow.txt
done
for i in {1..5};do
	sort -R shadow.txt
	echo `head -1 shadow.txt` | passwd --stdin user$i
done

