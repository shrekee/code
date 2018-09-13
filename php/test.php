<?php
	phpinfo();
#用来测试mysql数据库是否连接正常？？
	$conFlag=@mysql_connect("localhost","root","123");
	##如果数据库连接正常，则返回1，否则返回0
		if ($conFlag){
				    echo "连接成功！";

		}else{

				    echo "连接错误！";

		}
	?>
