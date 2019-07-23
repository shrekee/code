package main

/**
 * author       : liwenqiang
 *creating_time : 19-7-23 下午2:16
 * file_name    : test_05.py
 * IDE          : GoLand
**/

import (
	"fmt"
	"time"
	"sync"
)

var ch = make(chan int,10)
var wg sync.WaitGroup
func main()  {
	//go fun1()
	//go fun2()
	//go fun3()
	//time.Sleep(time.Second*60)
	//return
	//go fun()
	//count :=40
	//for  {
	//	fmt.Println("--main--")
	//	time.Sleep(1e9)
	//	if count > 0{
	//		count--
	//
	//	}else {
	//		return
	//	}
	//}
	wg.Add(2)
	for i:=0;i<10;i++{
		//fmt.Println("len(ch) is: ",len(ch))
		ch <- i
	}

	go fun1(ch)
	go fun2(ch)
	//time.Sleep(time.Duration(10)*time.Second)
	wg.Wait()
	close(ch)


}

func fun1(ch chan int)  {
	defer wg.Done()
	for  {
		item,ok :=<-ch
		if !ok{
			return
		}
		fmt.Println("fun--1--: ",item)
		time.Sleep(time.Second*2)
	}
}

func fun2(ch chan int)  {
	defer wg.Done()
	for  {
		item,ok :=<-ch
		if !ok{
			return
		}
		fmt.Println("fun--2--: ",item)
		time.Sleep(time.Second*2)
	}
}

func fun3()  {
	for  {
		fmt.Println("fun--3--")
		time.Sleep(time.Second*2)
	}
}


//func fun()  {
//	go fun1()
//	go fun2()
//	go fun3()
//	time.Sleep(time.Second*20)
//	fmt.Println("time passed is 20 seconds...")
//	return
//}

func fun4()  {
	for  {
		fmt.Println("fun--4--")
		time.Sleep(time.Second*2)
	}
}
