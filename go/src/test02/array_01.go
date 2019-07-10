package main

import "fmt"

func main() {
    var array = [10] int{1,2,3}
    //array_ := make([]int{1,2,3})
    array_ := make([]int,10)
    array_[0] = 1
    for i,_ := range array_ {
        array_[i] = i * 2
    }
    for i,v := range array_ {
        fmt.Println(i,v)
    }
    s := array[1:3]
    fmt.Println(s,array_,array,len(s))

    //append(100,array)
    fmt.Println(&array_[0])
    array_ = append(array_,100)
    fmt.Println(&array_[0])
    fmt.Println(s,array_,array,len(s))

}
