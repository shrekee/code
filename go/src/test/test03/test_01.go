package main

/**
 * author       : liwenqiang
 *creating_time : 19-7-17 上午9:31
 * file_name    : test_01.py
 * IDE          : GoLand
**/

import (
	"fmt"
	//"github.com/pkg/errors"
	//"database/sql"
	"github.com/pkg/errors"
)

func main() {

	s := make([]int, 0)
	s = append(s, 1, 2, 3, 4, 5)
	fmt.Println(s)
	fmt.Println("hello")
	fmt.Println("world")
	fmt.Println("s[:1]", s[:1])
	fmt.Println("s[:2]", s[2:])

	for i, v := range s {
		if v == 2 {
			fmt.Println(s)
			s = append(s[:i], s[i+1:]...)
			fmt.Println(s)
			s = append(s[:0], s[1:]...)
			fmt.Println(s)
			fmt.Println(s[:len(s)-1])
		}
	}

	s1 := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(s1)
	s1, _ = sliceRemove(0, s1)
	fmt.Println("remove index: 0 ==", s1)
	//s1 = []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(s1)
	s1, _ = sliceRemove(5, s1)
	fmt.Println("remove index: 5 ==", s1)
	//s1 := []int{1, 2, 3, 4, 5, 6, 7}
	//s1 = Remove(s1,)
	s2 := []int{1, 2, 3}
	s2, err := sliceRemove(1, s2)
	if err != nil {
		fmt.Println(err)
		return
	}

}
func sliceRemove(index int, s []int) ([]int, error) {
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

//
//func testRemove(index int, s *[]int) error{
//	if index == len(s) -1{
//		return errors.New("erros")
//	}
//	return errors.New("error")
//}
//
//func Insert(slice []interface{},start,end int)[]interface{}  {
//	reutrn []interface{}
//
//}

//func Remove(slice []interface{},start,end int)[]interface{}  {
//	result:=make([]interface{},len(slice)-(end-start))
//	at:=COPY(result,slice[:start])
//	COPY(result[at:],slice[end:])
//	return result
//
//}
