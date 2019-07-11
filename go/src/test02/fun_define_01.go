package main

/**
 * author       : liwenqiang
 *creating_time : 19-7-11 下午2:49
 * file_name    : fun_define_01.py
 * IDE          : GoLand
**/
import "errors"

func Add(a,b int) (ret int, err error) {
	if a < 0 || b < 0 {
		err = errors.New("Should be non-negarice numbers!")
		return
	}

	return a + b, nil
}