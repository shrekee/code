package main

/**
 * author       : liwenqiang
 *creating_time : 19-7-10 下午1:29
 * file_name    : reflect.py
 * IDE          : GoLand
**/

import (
	"fmt"
	"reflect"
)

type Bird struct {
	Name string
	LifeExpectance int
}

func (b *Bird) Fly()  {
	fmt.Println("I am flying...")
}

func main()  {
	sparrow := &Bird{"sparrow",3}
	s := reflect.ValueOf(sparrow).Elem()
	typeOfT := s.Type()

	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i, typeOfT.Field(i).Name, f.Type(),
			f.Interface())
	}
}