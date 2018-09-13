#!/bin/bash
# File Name: if_not_then_create_ftp_server.sh
# Author: Liwqiang
# mail: shrekee@qq.com
# Created Time: Sat 28 Jul 2018 10:39:51 AM CST

#检查系统是否有已安装ftp服务器，如果没有，则自行yum安装一个，并满足以下要求：
#1.不支持本地用户登陆
#2.匿名用户可以上传 新建 和删除
#3. 匿名用户限速500KBps

#第一步，检查系统是否已经安装ftp服务？
check_vsftpd_installed() {
	rpm -q vsftpd &> /dev/null
	if [ $? -ne 0 ];then
		yum install -y vsftpd  
		[ $? -ne 0 ] && echo "安装失败，请检查yum源。。" && exit 1
	fi
}

##禁用本地用户登陆,以及允许你用户登陆上传下载等

deny_localusers_login() {
	sed -ir 's/(^local_enable=)YES/\1NO/g' /etc/vsftpd/vsftpd.conf
	sed -ir 's/(^anon_mkdir_write_enable=)NO/\1YES/g' /etc/vsftpd/vsftpd.conf
	sed -ir 's/(^anon_other_write_enable=)NO/\1YES/g' /etc/vsftpd/vsftpd.conf
	sed -ir 's/(^anon_upload_enable=)NO/\1YES/g' /etc/vsftpd/vsftpd.conf
grep '^anon_mkdir_write_enable=YES' /etc/vsftpd/vsftpd.conf	
[ $? -ne 0 ] echo anon_mkdir_write_enable=YES >> /etc/vsftpd/vsftpd.conf
grep '^anon_other_write_enable=YES' /etc/vsftpd/vsftpd.conf	
[ $? -ne 0 ] echo anon_other_write_enable=YES >> /etc/vsftpd/vsftpd.conf
grep '^anon_upload_enable=YES' /etc/vsftpd/vsftpd.conf	
[ $? -ne 0 ] echo anon_upload_enable=YES >> /etc/vsftpd/vsftpd.conf

#anon_other_write_enable
#anon_upload_enable 
#NO
}

###禁用独立模式，进入xinetd模式

standby_mode() {
	sed -ir 's/(^listen=)YES/\1NO/g' /etc/vsftpd/vsftpd.conf
}

##把vsftpd托管给xinetd服务

##配置xinetd下的vsftpd服务配置

vsftpd_xinetd_conf() {
	cat<<-EOF
service ftp                                                                                                                 
{
        disable                 = no
        socket_type             = stream
        wait                    = no
        user                    = root
		only_from               = 10.1.1.0/24
		server                  = /usr/sbin/vsftpd
		per_source              = 5
		instances               = 200
		banner_fail             = /etc/vsftpd.busy_banner
		log_on_success          += PID HOST DURATION
		log_on_failure          += HOST
 }
EOF

}

##定义一个主函数
function main() {
	check_vsftpd_installed
	deny_localusers_login
	standby_mode
	echo `vsftpd_xinetd_conf` > /etc/xinetd.d/vsftpd
}

###调用主函数
main
