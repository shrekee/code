package main

/**
 * author       : liwenqiang
 *creating_time : 19-7-25 上午9:12
 * file_name    : time_test_1.py
 * IDE          : GoLand
**/
import (
	"time"
	//"fmt"
	"log"
)

func main()  {
	log.Println(time.Now())
	time.Sleep(1e9)
	log.Println(time.Now())

}
