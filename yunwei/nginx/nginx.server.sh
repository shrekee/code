#!/bin/sh
#
# nginx - this script starts and stops the nginx daemon
#
# chkconfig:   - 85 15
# description:  NGINX is an HTTP(S) server, HTTP(S) reverse \
#               proxy and IMAP/POP3 proxy server
# processname: nginx
# config:      /usr/loca/nginx/conf/nginx.conf
# config:      /etc/sysconfig/nginx
# pidfile:     /usr/local/nginx/logs/nginx.pid

##定义两个变量，分别保存nginx二进制的文件和配置文件路径
nginx="/usr/local/nginx/sbin/nginx"
NGINX_CONF_FILE="/usr/local/nginx/conf/nginx.conf"

start() {
	[ -x $nginx ] || exit 1
	[ -f $NGINX_CONF_FILE ] || exit 2
	echo "starting nginx..."
	deamon $nginx -c $NGINX_CONF_FILE
	retval=$?
	return $retval
}
stop() {
	[ -x $nginx ] || exit 3
	[ -f $NGINX_CONF_FILE ] || exit 4
	echo "stoping nginx..."
	$nginx -s quit
	return $?
}
reload() {
	configtest || return $?
	echo "reloadding nginx..."
	$nginx -s reload
}
status() {
	[ -x $nginx ] || exit 5
	[ -f $NGINX_CONF_FILE ] || exit 6
	$nginx -s status
}



configtest() {
	$nginx -t -c $NGINX_CONF_FILE 
}

case "$1" in 
	status)
		status
		;;
	start )
		start
		;;
	stop )
		stop
		;;
	reload )
		reload
esac




