package tools

import "errors"

/**
 * author       : liwenqiang
 *creating_time : 19-7-24 上午10:21
 * file_name    : readline.py
 * IDE          : GoLand
**/

func IntSRemove(s []int, index int) ([]int, error) {
	//fmt.Println("len(slice) is ", len(s))
	if index > len(s)-1 || index < 0 {
		return nil, errors.New("index out of range")
	}
	if index == len(s)-1 {
		s = s[:index]
	} else {
		s = append(s[:index], s[index+1:]...)
	}
	//fmt.Println("s is : ", s)
	return s, nil
}

func StrSRemove(s []string, index int) ([]string, error) {
	//fmt.Println("len(slice) is ", len(s))
	if index > len(s)-1 || index < 0 {
		return nil, errors.New("index out of range")
	}
	if index == len(s)-1 {
		s = s[:index]
	} else {
		s = append(s[:index], s[index+1:]...)
	}
	//fmt.Println("s is : ", s)
	return s, nil
}
