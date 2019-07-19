package main

/**
 * author       : liwenqiang
 *creating_time : 19-7-16 上午9:39
 * file_name    : interface_valuing_01.py
 * IDE          : GoLand
**/
import (
	"fmt"
)

func main()  {
	var zero Integer = 0
	var a_ Integer = 2
	var a Integer = a_
	fmt.Println(a.Less(zero),a.Add(zero))
	fmt.Println(a.Less(0),a.Add(5))
	//var b LessAdder = &a
	//var c LessAdder = &a
	//fmt.Println(b,c)
}

type Integer int

func (i *Integer) Less(b Integer) bool {
	return *i < b
}

func (i *Integer) Add(b Integer) Integer {
	return *i + b
}

type LessAdder interface {
	Less(a *Integer) bool
	Add(zz *Integer) Integer
}
