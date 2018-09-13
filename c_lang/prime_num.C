#include <math.h>
#include <iostream>
#include <string>
#include <ctime>
#include <vector>
#include <stdlib.h>
using namespace std;
bool isPrime_1( int num );
bool isPrime_2( int num );
bool isPrime_3( int num );

bool isPrime_1( int num )
{
    int tmp =num- 1;
    for(int i= 2;i <=tmp; i++)
      if(num %i== 0)
         return 0 ;
    return 1 ;
}

bool isPrime_2( int num )
{
     int tmp =sqrt( num);
     for(int i= 2;i <=tmp; i++)
        if(num %i== 0)
          return 0 ;
     return 1 ;
}

bool isPrime_3( int num )
{
                 //两个较小数另外处理
                 if(num ==2|| num==3 )
                                 return 1 ;
                 //不在6的倍数两侧的一定不是质数
                 if(num %6!= 1&&num %6!= 5)
                                 return 0 ;
                 int tmp =sqrt( num);
                 //在6的倍数两侧的也可能不是质数
                 for(int i= 5;i <=tmp; i+=6 )
                                 if(num %i== 0||num %(i+ 2)==0 )
                                                 return 0 ;
}
int main()
{
                 int test_num =400000;
                 int tstart ,tstop; //分别记录起始和结束时间
                 //测试第一个判断质数函数
                 tstart=clock ();
                 for(int i= 1;i <=test_num; i++)
                                 isPrime_1(i);
                 tstop=clock();
                 cout<<"方法(1)时间(ms):" <<tstop- tstart<<endl ;//ms为单位
                 //测试第二个判断质数函数
                 tstart=clock();
                 for(int i= 1;i <=test_num; i++)
                                 isPrime_2(i );
                 tstop=clock();
                 cout<<"方法(2)时间(ms):" <<tstop- tstart<<endl ;
                 //测试第三个判断质数函数
                 tstart=clock();
                 for(int i= 1;i <=test_num; i++)
                                 isPrime_3(i );
                 tstop=clock();
                 cout<<"方法(3)时间(ms):" <<tstop- tstart<<endl ;
                 cout<<endl ;
                 system("pause" );
                 return 0 ;
}
