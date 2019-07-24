package main

/**
 * author       : liwenqiang
 *creating_time : 19-7-24 上午9:26
 * file_name    : scanln_test_06.py
 * IDE          : GoLand
**/

import (
	"fmt"
)
var(
	firstName,lastName,s string
	i int
	f float32
	input ="56.12 / 5212 Go"
	format ="%f / %d / %s"
)

func main()  {
	fmt.Println("Please enter your full name: ")
	fmt.Scanln(&firstName,&lastName)
	fmt.Printf("%s  %s",firstName,lastName)

}
