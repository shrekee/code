package main

import (
	"fmt"
	"time"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)
	fmt.Println(time.StampMicro)
}
