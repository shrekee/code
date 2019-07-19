package main

/**
 * author       : liwenqiang
 *creating_time : 19-7-18 下午4:18
 * file_name    : test_02.py
 * IDE          : GoLand
**/

import "fmt"

func main()  {
	fun()
	fmt.Println("Done.")

	r := fun_2()
	if value,ok := r.([1]int);ok{
		fmt.Println("Ok, type assert")
		fmt.Println("value is : ",value)
	} else {
		fmt.Println("Faild, type assert")
		fmt.Println("value is : ",value)
		return
	}


}

func fun()int  {
	fmt.Println(11)
	return 11
}

func fun_2()interface{}  {
	return [1]int{1}
}