package main

/**
 * author       : liwenqiang
 *creating_time : 19-7-24 下午2:39
 * file_name    : test_06.py
 * IDE          : GoLand
**/
import (
	"fmt"
	//"os"
	"strings"
)

func main()  {
	//fmt.Println("Finish.")
	//os.Exit(1)
	s:="hello"
	if strings.Contains(s,"hello"){
		fmt.Println("hello Contains hello")
	}else{
		fmt.Println("hello does not Contain hello")
	}


}