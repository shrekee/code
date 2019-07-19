package main

import "fmt"

/**
 * author       : liwenqiang
 *creating_time : 19-7-16 上午9:56
 * file_name    : interface_valuing_02.py
 * IDE          : GoLand
**/

type Integer int

func (a Integer) Less(b Integer) bool {
	return a < b
}

func (a *Integer) Add(b Integer) Integer {
	return *a + b
}

type LessAdder interface {
	Less(b Integer) bool
	Add(b Integer) Integer
}

func main() {
	var a Integer = 1
	var b LessAdder = &a
	fmt.Println(b)

	var b1 Lesser = a
	fmt.Println(b1)
}

type Lesser interface {
	Less(b Integer) bool
}

