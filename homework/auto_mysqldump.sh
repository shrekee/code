#!/bin/bash
# File Name: auto_mysqldump.sh
# Author: Liwqiang
# mail: shrekee@qq.com
# Created Time: Sat 04 Aug 2018 09:54:47 PM CST

#在每天凌晨4点（在服务器最冷清时热备份,crontab 来执行），通过mysqldump自动热备份数据库（全备）
#
#热备份的要求：
#    1、因为可能同时有myisam和innodb两种引擎，所以要backup前要锁表；--lock-tables
#    2、全备 所以加参数 --all-databases
#    3、为了将来好还原，可以截断二进制日志，--flush-logs
#    4、为了能认清楚     加参数  --master-data=2（）
#	 5、myisam   温备份（备份之前务必锁表）
#	 6、	--lock-all-tables(全备份)
#	 7、	--lock-tables（只备份几张表）
#	 8、innoDB   热备份
#			--single-transaction  
#			--flush-logs
#			--routines   备份过程和函数
#			--triggers   备份触发器
##			--events  --备份事件
#			--master-data=（0|1|2）  2：一般是2；除非用的主从服务器，用1选项；
##
#MySQL热备份的缺点
#		1、浮点数据丢失精度
#		2、备份出的数据占用更大的存储空间，压缩后可节省空间
#		3、不适合对大的数据库完全备份
#		4、也不是对要求准确的使用。。。。
#		5、而且对于innoDB引擎的更大的缺点：
#			加锁过程可能很长（可能几个小时）；

##注意：在使用mysqldump进行逻辑上的恢复时，会自动应用于bin_log日志，，
##所以，在使用mysqldump恢复时，必须应临时关闭'set sql_bin_log=0';恢复后，然后再打开；
mysqldump -uroot -p123 --flush-logs --master-data=2 --all-datebases --lock-alltables > /tmp/mysqldump_full_backup

###
###     SELECT备份的优缺点：：：：
###			优点：比mysqldump节省空间，而且恢复时，可以选择只恢复部分数据   
###			缺点：不能备份数据结构，而且只能备份单张表格；
###使用	SELECT  * INTO OUTFILE ‘/PATH/TO/NAME ’ FROM  tables_name[WHERE CLAUSE];   创建表格备份；
###     LOAD DATA INFILE 'FILE_NAME' INTO TABLE  TABLE_NAME[WHERE CLAUSE];   恢复数据
###


#######################################################################
######################################################################

#备份方法三    LVM热备份 --SNAPSHOT  
##逻辑卷备份：把数据锁定在瞬间，然后快照，然后释放锁，然后把数据物理copy到别的磁盘
######要求：
#		1、数据文件要求在逻辑卷上；
#		2、此逻辑卷所在卷组必须有足够空间使用快照；
#		3、二进制日志和事物日志，跟数据文件必须在同一个卷；这样才能保持数据一致；
#		4、
#		
#
#

