package main

/**
 * author       : liwenqiang
 *creating_time : 19-7-16 上午11:15
 * file_name    : channel_test_01.py
 * IDE          : GoLand
**/

import (
	"fmt"
	"time"
)

func Count(ch chan int)  {
	ch <- 1
	fmt.Println("Counting")
}

func main()  {
	chs := make([]chan int, 10)
	for i:=0;i<10 ;i++  {
		chs[i] = make(chan int)
		go Count(chs[i])
	}

	for _,ch := range chs{
		<-ch

	}
	time.Sleep(1e8)

}
