#!/bin/bash
#nginx 源码安装nginx,并根据用户是需求，自定义配置文件、二进制文件路径


 #定义一个输入路径的函数，递归方法实现
function input_nginx_dir() {
	read -p"请输入您已经解压缩后的nginx源包路径" nginx_dir
	cd $nginx_dir 
	#通过cd 命令，来判断输入路径是否正确，如果出错则重新输入。
	[ $? -ne 0 ] && echo "路径输入错误，请重新输入" && input_nginx_dir
	#扫描当前目录下的文件名称，并保存。
	list=`/bin/ls`
	#定义一个变量，用来保存解压nginx源码后，一般会存在的文件
	nginx_pack_names="html configure conf man"
	for i in $nginx_pack_names;do
		echo $list | grep "$i" &>/dev/null
		#逐一取出常用的名字，并用来和list中做比对，如果不存在，则提示用户：
		if [ $? -ne 0 ]; then
			read -p"您输入的目录中并不包含$i文件，您确定吗？yes/no" ans1
			if [ "$nas1" = 'yes' -o "$ans1" = 'y' ];then
				continue
			else
				input_nginx_dir
			fi
		fi
	done				
}


##第一步，首先解决源码包安装nginx所需要的依赖“gcc pcre-devel openssl-devel zlib-devel”
pack="gcc pcre-devel openssl-devel zlib-devel"
yum -y install $pack

#遍历软件是否已全部安装成功?
for i in $pack 
do	
	rpm -q $i &> /dev/null
	[ $? -ne 0 ] && echo "$i 未能成功安装,请检查您的环境" && exit 1
done


#判断在“/etc/passwd”否有nginx用户，如果没有则自动创建
#创建nginx系统用户，无登陆权限，无家目录

#取出“/etc/passwd”中的用户
user_name=`cat /etc/passwd | cut -d':' -f1`
echo $user_name | grep 'nginx' > /dev/null
if [ $? -ne 0 ]; then
#此处不太严谨,并没有判断在/etc/group中有无nginx组用户
	groupadd -r nginx &>/dev/null
	useradd -r -g nginx -s /sbin/nologin  -M nginx
fi



#让用户指出解压缩后的nginx目录,如果目录输入错误则要求用户一再输入;而且具有基本的判断功能，
#即帮助用户来判断：用户输入的目录是否就是正确的目录
input_nginx_dir

#打印菜单,包含nginx的常用安装路径
cat <<EOF
*******默认安装路径*******
默认用户为nginx，默认组为nginx
默认安装路径：/usr/local/nginx
默认的二进制路径：/usr/local/nginx/sbin
默认的配置路径： /usr/local/nginx/conf
默认启用openssl模块

EOF
read -p“是否同意以上默认配置，yes/no” ans

if [ $ans = 'yes' -o $ans = 'y' ] ;then
	./configure --user=nginx --group=nginx --with-http_ssl_module
	[ $? -eq 0 ] && make && make install || echo "安装失败，请查看错误信息"
#创建二进制文件的软连接
	ln -sv /usr/local/nginx/sbin/nginx /usr/bin/nginx
#判断是否安装成功
	/usr/bin/nginx -V
	[ $? -eq 0 ] && echo "恭喜您，安装成功"
##




else
	read -p"请输入nginx的安装路径" nginx_install
	read -p"请输入二进制安装路径" bin_dir
	read -p"请输入配置文件路径"  conf_dir
 
 ./configure  --prefix=$nginx_install --user=nginx --group=nginx --with-http_ssl_module --sbin-path=$bin_dir ----conf-path=$conf_dir
[ $? -eq 0 ] && make && make install || echo "安装失败，请检查错误信息"
fi
 
 
 
 
 
 
 
 
 
 
 
 
 
