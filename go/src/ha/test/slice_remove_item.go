package main

/**
 * author       : liwenqiang
 *creating_time : 19-7-24 上午10:43
 * file_name    : slice_remove_item.py
 * IDE          : GoLand
**/

import (
	"fmt"
	"strings"
)

func main()  {
	s,v:="ping:1.1.1.1:ok","ping"
	if strings.HasPrefix(s,v){
		fmt.Printf("%v hasPrefix %v",s,v)
	}else {
		fmt.Printf("%v does not hasPrefix %v",s,v)
	}


}
