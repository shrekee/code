package main

/**
 * author       : liwenqiang
 *creating_time : 19-7-16 上午10:22
 * file_name    : interface_test_02.py
 * IDE          : GoLand
**/

import "fmt"

func main()  {
	var a int

	if _,ok :=  a.(int);ok{
		fmt.Println("a is int type")
	}

}

