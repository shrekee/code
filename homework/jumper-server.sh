#!/bin/bash
# File Name: jumper-server.sh
# Author: lsing
# mail: shrekee@qq.com
# Created Time: Fri 27 Jul 2018 12:53:49 PM CST

#trap ' ' 1 2 3  #用于捕捉信号，然后对信号做相应的处理/动作

#定义一个菜单函数，用于打印用户可以选择的菜单
menu( ) {
 cat<<EOF
欢迎使用Jumper-server，请选择你要操作的主机：
1. DB1-Master
2. DB2-Slave
3. Web1
4. Web2
5. exit
6. 菜单
EOF
}

#定义一个选择函数，帮助用户选择某个选项，，如果用户不选择或者选择错误，则会提示用户一直选择，直至用户退出
function choise( ) {
			read -p"$1" cho
			while [ -z $cho ]; do
				read -p"$1" cho
			done
			echo $cho
			}


#定义一个检查本地跳板主机，有无ssh公钥和密钥的函数，如果没有，则自动创建一个。
function check_rsa_key_exist( ) {
			if [[ -f /home/yunwei/.ssh/id_rsa && -f /home/yunwei/.ssh/id_rsa.pub && -n /home/yunwei/.ssh/id_rsa  && -n /home/yunwei/.ssh/id_rsa.pubi ]] ; then
				echo
			else	
				mkdir -p /home/yunwei/.ssh
				$(/usr/bin/ssh-keygen -P '' -f /home/yunwei/.ssh/id_rsa)
				echo "密钥创建成功..."

			fi
		}
##把运维主机的公钥推送所有up的主机上（通过ssh-copy-id）的方法，当然只有主机默认ssh端口为22才可以。
function push_public_key( ) {
		rpm -q expect
		if [ $? -ne 0 ] ;then
			 sudo yum -y install expect
			[ $? -ne 0 ] && echo "安装expect软件，自动退出..." && sleep 2 && exit
		fi
		for i in `cat $1`
		do
			/usr/bin/expect <<-EOF
			spawn ssh-copy-id root@${i}
			expect {
				"yes/no" { send "yes\r";exp_continue }
				"password:" { send "123\r" }
			}
			expect eof
			EOF
		done
		}
##检查所有私网的可用的服务主机的ip，并保存一个文件/tmp/ip_up.txt中
function check_server_up( ) {
	> /tmp/ip_up.txt   ##清空当前文件
		ping -c1 10.1.1.1 &> /dev/null
		[ $? -eq 0 ] && echo "10.1.1.1" >> /tmp/ip_up.txt
	for i in {3..20}; do
		
		{
		ping -c1 10.1.1.$i &> /dev/null
		[ $? -eq 0 ] && echo "10.1.1.$i" >> /tmp/ip_up.txt
	    } &
	done
		wait
		#echo `grep -v '  10.1.1.2'   /tmp/ip_up.txt` > /tmp/ipip
		#echo `grep -v '10\.1\.1\.2'  /tmp/ip_up.txt`| cat  > /tmp/ip_up.txt
		echo "/tmp/ip_up.txt"
		}
#定义主函数，用于完成主要功能。
function main( ) {
			check_rsa_key_exist			
			a=`check_server_up`		
			push_public_key $a
			menu
			de=`choise 请输入您的选择`
			case $de in
				1 )
				echo "choise 1"
				;;
				2 )  
				echo "choise 2"
				;;
				3 )  
				echo "choise 3"
				;;
				4 )  
				echo "choise 4"
				;;
				5 )  
				echo "choise 5"
				;;
				6 )  
				menu
				;;
			esac

		}
main

