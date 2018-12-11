#/bin/bash
#x写一个脚本，对日志进行轮转、打包，然后把30天前的日志删除。
tar -Jcf conn.`date +$F`.tar.xz  conf.log
find / -mtime +30 -type f -exec rm -f {} \;
