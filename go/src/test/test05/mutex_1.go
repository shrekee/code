package main

/**
 * author       : liwenqiang
 *creating_time : 19-7-26 下午5:29
 * file_name    : mutex_1.py
 * IDE          : GoLand
**/
import (
	"sync"
)

//var mutex1 sync.Mutex
var (
	counter int64
	wg sync.WaitGroup

)

