package main

import (
	"fmt"
	"time"
	//"sync"
)

/**
 * author       : liwenqiang
 *creating_time : 19-7-29 下午1:25
 * file_name    : test_4.py
 * IDE          : GoLand
**/

func main()  {
	fun()
	go func() {
		for  {
			fmt.Println("In mainthread...")
			time.Sleep(2e9)
		}
	}()
	select {}

}
func fun_1()  {
	for  {
		fmt.Println("======  fun_1  ======")
		time.Sleep(1e9)
	}
}

func fun_2()  {
	for  {
		fmt.Println("======  fun_2  ======")
		time.Sleep(1e9)
	}
}

func fun()  {
	go fun_1()
	go fun_2()
	return

}
