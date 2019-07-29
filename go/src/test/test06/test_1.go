package main

/**
 * author       : liwenqiang
 *creating_time : 19-7-29 上午8:55
 * file_name    : test_1.py
 * IDE          : GoLand
**/

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello, world!")
	//a()
	//f()
	callback(1,Add)
}

func a() {
	i := 0
	defer fmt.Println(i)
	i++
	return
}

func f() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}

}
func Add(a,b int) int {
	return a+b
}

func callback(z int,f func(x,y int)int) {
	fmt.Println("sum is: ", f(z, 2))
}

func callball_2()  {
	strings.IndexFunc()

}
