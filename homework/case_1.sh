#!/bin/bash

read -p"请输入一个等级（A-E），查看每个等级的成绩" grade

case $grade in
	A|a)
		echo "90~100分"
	;;
	B|b)
		echo "80~89分"
	;;
	C|c)
		echo "70~79分"	
	;;
	D|d)
		echo "60~69分"
	;;
	E|e)
		echo "50~59分"
	;;
	F|f)
		echo "40~49分"
	;;
	*)
		echo "请输入A-E之间的字母"
	;;
esac
