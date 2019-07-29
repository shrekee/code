package main

import (
	"fmt"
	"errors"
)

/**
 * author       : liwenqiang
 *creating_time : 19-7-29 上午10:14
 * file_name    : test_2.py
 * IDE          : GoLand
**/
func main() {
	//s := []string{"hello"}
	//////s, err := remove(s, "hello")
	//
	//fmt.Println(s, err)
	assert_("hello")

}

func remove(slice []interface{}, elem interface{}) ([]interface{}, error) {
	if len(slice) == 0 {
		return slice, nil
	}

	flag := false
	for i, v := range slice {
		if v == elem {
			slice = append(slice[:i], slice[i+1:]...)
			flag = true
			break
		}
	}
	if flag == true {
		return slice, nil
	} else {
		return slice, errors.New("not find item")
	}
}

func assert_(i interface{})error  {
	if value,ok := i.(string);ok{
		fmt.Printf("type is string, value is: %v, type is: %T",value,value)
		return nil
	}
	return errors.New("type is not string")

}