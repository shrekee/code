/*************************************************************************
    > File Name: c_1.c
    > Author: Liwqiang
    > Mail: shrekee@qq.com 
    > Created Time: Mon 13 Aug 2018 06:39:32 PM CST
	>有一个字符串’1a2b3d4z‘,写一个函数实现以下功能：
	>功能1：把偶数位字符挑选出来，组成一个字符串1
	>功能2：把奇数位字符挑选出来，组成一个字符串2
	>功能3：把字符串1和字符串2，通过函数参数，传递给main函数，并打印
	>功能4：主函数能测试通过。
	>int getStr1Str2(char *source, char *buf1, char *buf2)
 ************************************************************************/
/*指针也是一种变量，也有空间，32位程序
 
 */

#include<stdio.h>
//定义一个函数，用来获取子字符串1和2，分别只取主字符串的奇数位和偶数位；
int getStr1Str2(char *source, char *buf1, char *buf2) {
		int i=0,j=0,k=0;
	 while(source[i])
	 {
		if((i)%2  == 1)
		{	*(buf1+j)= source[i];
			j++;
		}
		else
		{
			*(buf2+k)= source[i];
			k++;
		}
		i++;
	 }

	

}
//定义一个主函数，，并调用子函数。
int main() {	
	char str[] = "1a2b3d4z";
	char str1[20];
	char str2[20];
	getStr1Str2(str,str1,str2);
	printf("%s,%s,%s",str,str1,str2);
	
	


}

