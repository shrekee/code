#!/bin/bash


#计算距离2019年新年元旦凌晨零时零秒还有多少天多少时多少分  多少秒？
day=0；hour=0;minute=0;second=0;
future=`date -d "2019-1-1 0:0:0 " "+%s"`
while true ;do
	now=`date "+%s"`
	let time=future-now
	let day=time/86400
	let shi=time%86400
	let hour=shi/3600
	let fen=shi%3600
	let minute=fen/60
	let second=fen%60

	echo "距离2019年新年元旦还有${day}天${hour}时${minute}分${second}秒"
	sleep 1
done
