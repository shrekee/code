package main

import "fmt"

func main() {
    a := make([]int,1)
    b := make([]int,1,2)

    a[0] = 321
    fmt.Println("len(a)=",len(a),"cap(a)=",cap(a))
    fmt.Println("len(b)=",len(b),"cap(b)=",cap(b))

    a = append(a,123)
    b = append(b,123)

    fmt.Println("len(a)=",len(a),"cap(a)=",cap(a))
    fmt.Println("len(b)=",len(b),"cap(b)=",cap(b))

    c := a[:1]

    fmt.Println("len(a)=",len(a),"cap(a)=",cap(a))
    fmt.Println("len(c)=",len(c),"cap(c)=",cap(c))

    c = append(c,456)
    fmt.Println("len(a)=",len(a),"cap(a)=",cap(a))
    fmt.Println("len(c)=",len(c),"cap(c)=",cap(c))

    p := &c[0]
    fmt.Println(p)

}
