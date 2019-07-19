package main

/**
 * author       : liwenqiang
 *creating_time : 19-7-12 下午4:15
 * file_name    : channel_test_01.py
 * IDE          : GoLand
**/
import "fmt"
import "time"

func Count(ch chan int, value int) {
	ch <- value
	fmt.Println("Counting ")
	time.Sleep(1e9)
}

func main() {
	chs := make([]chan int, 10)
	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		go Count(chs[i], i)
	}

	for _, ch := range chs {
		value := <-ch
		fmt.Println("value ", value)
		//time.Sleep(1e7)
	}
	time.Sleep(1e9)
}
