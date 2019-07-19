package main

/**
 * author       : liwenqiang
 *creating_time : 19-7-15 上午9:54
 * file_name    : test_processing.py
 * IDE          : GoLand
**/
import (
	"time"
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func cal() {
	defer wg.Done()
	var sum uint64 = 0
	var i uint64
	for i = 0; i < 1e8; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func main() {
	startTime := time.Now()
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go cal()
	}
	wg.Wait()
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	fmt.Println("duration is ", duration)
}
