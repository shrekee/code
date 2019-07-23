package main

import (
	"fmt"
	"time"
)

/**
 * author       : liwenqiang
 *creating_time : 19-7-23 上午9:51
 * file_name    : select_test_01.py
 * IDE          : GoLand
**/
var exitChan chan bool

func Waiting1() {
	defer func() {
		fmt.Println("waiting1 exit")
	}()
	//do someting
	time.Sleep(time.Second * 10)
	exitChan <- true
}

func Waiting2() {
	defer func() {
		fmt.Println("waiting2 exit")
	}()
	for {
		select {
		case <-time.After(time.Second * 2):
			fmt.Println("tick event...")
		case <-exitChan:
			fmt.Println("exit event...")
			//break
			return
		}
	}

}

func main() {
	exitChan = make(chan bool)
	go Waiting1()
	go Waiting2()
	<-time.After(time.Second * 60)
	fmt.Println("main return")

}
