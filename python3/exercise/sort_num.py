#!/usr/local/bin/python3
#coding:utf-8


#快速排序法
lst=[3,1,2,5,6,9]


def sort(array) :
	i=1
	l=len(array)
	bigger=list()
	smaller=list()
	while i<= l-1:
		if array[i]>=array[0]:
			bigger.append(array[i])
		else :
			smaller.append(array[i])
		i+=1
	lst_mid=[array[0]]
	lst_new=smaller+lst_mid+bigger
	print(lst_new)

if __name__ =='__main__':
	sort(lst)


		
		
