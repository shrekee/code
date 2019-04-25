#!/bin/bash


#相亲大会
#提示客户输入信息，如果客户拒绝输入，再重复要求输入信息
get_info() {
	input=''
	while [[ -z $input  ]]; do
		read -p"$1" input
	done
	echo $input
}


function main() {

		name=`get_info	请输入您的姓名`
		gender=`get_info 请输入您的性别`
		age=`get_info 请输入您的年龄`

		case $gender in
			'')
				echo "您是泰国来的吧？"
				;;
			male|nan|man|男)
				if (( age < 18 )); then	
					echo "小毛孩来凑什么热闹？"
				else
					read -p"${name}先生，您结婚了吗？" ans1
						case $ans1 in
							yes|y|是 )
								echo "结了来凑什么热闹？"
								;;
							no|n|没有 )
								read -p"您有房有车吗" ans2
									case $ans2 in
										y|yes )
											echo "咱们去民政局吧?"
											;;
										n|no )
											echo "我去下洗手间"
											;;
										'' )
											echo "别浪费时间，请正面作答"
											;;
									esac
									;;
								'' )  
									echo "你到有没有结婚"
									;;
						esac
					fi
					;;

			woman|female|nv|女 )
				if [[ age>=18 ]]; then
					echo "$name女士，您好"
				else
					echo "$name小姐，您好"
				fi
				;;

		esac
				
}
main
