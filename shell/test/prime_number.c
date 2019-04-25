/*************************************************************************
    > File Name: prime_number.c
    > Author: Liwqiang
    > Mail: shrekee@qq.com 
    > Created Time: Sun 12 Aug 2018 09:54:07 PM CST
 ************************************************************************/

#include<stdio.h>
#include<math.h>
#include<time.h>
//#include<stdlib.h>
int i,j,k=1;
//初步定义n=40万，目的是计算40万内的素数的时间，结果很吃惊，在我的单核虚拟机中，只用了0.12秒
unsigned int n=400000;
int prime[40000];
//定义主函数
int main() {
	prime[0]=2;
	long start=clock();
	for(i=3;i<=n;i+=2){
		int	mid=sqrt(i);		
			for(j=2;j<= (mid+1);j++){
				if(i%j==0)
					break;
				if(j==(mid+1)){
					prime[k]=i;
					k++;
				}
			}

	}
	long end=clock();
	double duration = (double)(end - start) / CLOCKS_PER_SEC;
	for(i=0;i<k;i++){
		printf("%d ",prime[i]);}
	printf("\n");
	printf("总共耗时为： %f 秒\n在%d 个自然数中有%d 个素数\n",duration,n,k);
}
		
