package main

/**
 * author       : liwenqiang
 *creating_time : 19-7-16 上午9:14
 * file_name    : struct_test_01.py
 * IDE          : GoLand
**/

import (
	"fmt"
)

func main()  {
	//s := Student{name:"lsing",age:34,cource:"go_lang",gender:"male"}
	s1 := Student{People{"shrek",32},"golang","male"}
	////act := s
	act := s1
	fmt.Println(act)
	act.eat()
	act.sing()
	var s Student
	s.name = "lsing"
	s.age = 34
	s.gender = "male"
	s.cource = "go_lang"
	s.name="student_1"
	s.age = 13

	fmt.Println(s)
	s.sing()
	s.eat()
	fmt.Println("*****************")
	p := People{"xiaohong",18}
	p.eat()
	p.sing()
}

type People struct {
	name string
	age uint
}

type Student struct {
	People
	cource string
	gender string
}

func (p People)eat()  {
	fmt.Printf(" %s is eating...\n",p.name)
}

func (p People)sing()  {
	fmt.Printf("%s is singing...\n",p.name)
}

type action interface {
	eat()
	sing()
}
