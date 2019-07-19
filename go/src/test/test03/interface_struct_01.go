package main

/**
 * author       : liwenqiang
 *creating_time : 19-7-16 下午1:58
 * file_name    : interface_struct_01.py
 * IDE          : GoLand
**/

import (
	"fmt"
	//"bytes"
	"bytes"
)

type Eater interface {
	eat()
}

type eatType struct {
	Eater
}

func main() {
	var a eatType
	fmt.Println(a)
	fmt.Printf("%T,%v\n", a, a)

	b1 :=[]byte("this is a first string")
	b2:=[]byte("this is a second string")
	fmt.Printf("%T,%v\n%T,%v\n",b1,b1,b2,b2)

	var buffer bytes.Buffer
	buffer.Write(b1)
	buffer.Write(b2)
	b3:=buffer.Bytes()
	fmt.Printf("%T,%v\n",b3,b3)

}
