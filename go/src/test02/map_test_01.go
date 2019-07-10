package main

import "fmt"

func main() {
    var m = map[string]string{"name":"lsing","gender":"male"}
    for k,v := range m {
        fmt.Println(k,v)
    }
}
