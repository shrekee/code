#!/bin/bash

#方法一
sum=0
for((i=0;i<=50;i+=2));do
	let sum+=i

done
echo "1到50的偶数之和${sum}"


#方法二

sum1=0
for i in {0..50..2};do
	let sum1+=i
done


echo "1到50的偶数之和${sum1}"


#方法三

sum2=0

for i in {1..50};do
	if [ $[$i%2] -ne 0 ];then
		continue;
		
	else
		let sum2+=i
	fi

done


echo "1到50的偶数之和${sum2}"


#方法四

sum3=0
for i in `seq 0 2 50 `;do
	let sum3+=i

done


echo "1到50的偶数之和${sum3}"






