package main

import "time"

/**
 * author       : liwenqiang
 *creating_time : 19-7-12 下午5:16
 * file_name    : select_test_01.py
 * IDE          : GoLand
**/
type Vector []float64

timeout := make(chan bool, 1)

go func () {
	time.Sleep(1e9)
	timeout <- true
}()

//
select {
case <-ch:
//...
fmt.Println("zz")
case <- timeout:
//..
}

func (v Vector) DoSome(i, n int, u Vector, c Chan int) {
	for ; i < n; i++ {
		v[i] += u.Op(v[i])
	}
}

const NCPU = 16

func (v Vector) DoAll(u Vector) {

	c := make(chan int, NCPU)

	for i := 0; i < NCPU; i++ {
		go v.DoSome(i*len(v)/NCPU, (i+1)*len(v)/NCPU, u, c)
	}
}
