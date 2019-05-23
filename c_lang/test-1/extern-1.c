#include<stdio.h>

int func();

int main()
{
        func(); //1
        extern int num;
        printf("%d",num); //2
        return 0;

}

int num = 3;

int func()
{
        printf("%d\n",num);

}

