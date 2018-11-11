declare -A ipArray 
#定义一个关联数组
cat /root/code/yunwei/test/access.log | while read a b flag ip d 
do    #统计所有error的IP，保存至数组
if [ $flag = '[error]' ]; then
	let ipArray[$ip]++
fi
done 
 
# 把ip数组依次输入到文档中
for key in ${!ipArray[*]}
do
	echo "${ipArray[$key]} $key" >> /tmp/sort.txt
done

sort /tmp/sort.txt |tail -10 
#排序，并去除前十个IP

