package main

/**
 * author       : liwenqiang
 *creating_time : 19-7-12 下午3:46
 * file_name    : add.py
 * IDE          : GoLand
**/
import "fmt"
//import "time"
import "sync"
import "runtime"

var counter int = 0

func Count(lock *sync.Mutex)  {
	lock.Lock()
	counter++
	fmt.Println(counter)
	lock.Unlock()
}

//func Add(x,y int)  {
//	z := x + y
//	fmt.Println(z)
//
//}

func main()  {
	lock := &sync.Mutex{}
	for i := 0;i < 10; i++ {
		go Count(lock)
	}

	for  {
		lock.Lock()
		c := counter
		lock.Unlock()
		runtime.Gosched()
		if c >= 10 {
			break
		}

	}


	fmt.Println("Done")

}