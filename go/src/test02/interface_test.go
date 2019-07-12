package main

import "math/rand"

/**
 * author       : liwenqiang
 *creating_time : 19-7-12 下午1:15
 * file_name    : interface_test.py
 * IDE          : GoLand
**/
type File struct {
	// ...
}

func (f *File) Read(buf []byte) (n int, err error){
	//...
}
func (f *File) Write(buf []byte) (n int, err error)  {
	//...
}

type Music struct {
	ID string
	Name string
	Artist string
	Source string
	Type string
}
