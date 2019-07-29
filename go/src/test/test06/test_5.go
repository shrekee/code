package main

/**
 * author       : liwenqiang
 *creating_time : 19-7-29 下午2:50
 * file_name    : test_5.py
 * IDE          : GoLand
**/
import (
	"fmt"
	//"runtime"
)

func main()  {
	var buf [5]byte
	fmt.Println("len of(buf[:]) is: ",len(buf[:]),buf[:])
	fmt.Println(buf[0:0:len(buf)])
	//n:=runtime.Stack(buf[:],false)

	//newBuf:=buf[:]
	//fmt.Printf("buf[5] %T, buf[:] %T",buf,newBuf)

}
