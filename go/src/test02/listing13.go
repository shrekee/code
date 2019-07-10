package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	counter int64
	wg sync.WaitGroup
)

func main()  {
	//计数加2,表示要等待两个goroutine
	wg.Add(2)

	go incCounter(1)
	go incCounter(2)

	wg.Wait()

	fmt.Println("Final Counter:", counter)

}

func incCounter(id int)  {
	//在函数退出时调用Done来通知main函数工作已经完成
	defer wg.Done()


}

