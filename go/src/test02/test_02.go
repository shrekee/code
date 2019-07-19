package main

/**
 * author       : liwenqiang
 *creating_time : 19-7-15 上午10:35
 * file_name    : test_02.py
 * IDE          : GoLand
**/
import "fmt"

var v uint64 = 1

func main() {

	for i := 0; i < 100; i++ {
		v += uint64(i)
	}
	fmt.Println(v)
	switch 1 {
	case 0:
		fmt.Println(0)
	case 1:
		fmt.Println(1)

	}
}
