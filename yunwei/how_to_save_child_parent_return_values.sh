#!/bin/bash

#通过管道输出的信息，会进入一个子进程中。。所以，在子进程中的信息，，无法在父进程中使用。。
#如何把通过管道输出的变量保存到数组中。。。答案：通过中间变量的方法。。
#exec  declare -a array
 outfile=`df |tr -s " "|cut -d' ' -f6|grep '^/'` 

 while read i;do
  array[$j]=$i;
  let j++;
  echo $j;
done <<EOF
$outfile
EOF


echo ${array[@]}
echo ${#array}
