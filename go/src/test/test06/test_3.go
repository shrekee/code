package main

/**
 * author       : liwenqiang
 *creating_time : 19-7-29 上午10:40
 * file_name    : test_3.py
 * IDE          : GoLand
**/
import (
	"fmt"
)

func main()  {
	var t interface{}
	//t=[]string{"hello"}
	t=1

	switch t:=t.(type) {
	default:
		fmt.Printf("unexpected type %T",t)
	case int:
		fmt.Printf("int %t\n",t)
	case string:
		fmt.Printf("sting %s\n",t)
	}

}
