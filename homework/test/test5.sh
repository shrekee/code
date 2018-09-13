#!/bin/bash
declare -i sum=0
for ((i=1;i<=100;i++)) ;do
	if [ $[$i % 2] -ne 0 ] ;then
		let sum+=i
	fi

done
echo $sum
		
