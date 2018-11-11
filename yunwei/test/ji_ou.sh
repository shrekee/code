
for i in {1..100};do
	let v=($i % 2)
	if [ $v -gt 0 ];then
		echo "$i">> 1.txt
	else
		echo "$i">> 2.txt
	fi
done
