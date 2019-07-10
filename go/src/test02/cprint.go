package main

/*
# include <stdio.h>
 */
/**
 * author       : liwenqiang
 *creating_time : 19-7-10 下午1:42
 * file_name    : cprint.py
 * IDE          : GoLand
**/
import "C"
import "unsafe"

func main()  {
	cstr := C.CString("Hello, world")
	C.puts(cstr)
	C.free(unsafe.Pointer(cstr))

}
