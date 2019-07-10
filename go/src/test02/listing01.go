package main

import (
    "fmt"
    "runtime"
    "sync"
)

func main() {
    // 分配一个逻辑处理器给调度器使用
    runtime.GOMAXPROCS(1)

    // wg用来等待程序完成
    //
    var wg sync.WaitGroup
    wg.Add(2)

    fmt.Println("Start Goroutines")

    //
    go func() {
        defer wg.Done()

        for count := 0; count < 3; count++ {
            for char := 'a'; char < 'a'+26; char++ {
                fmt.Printf("%c ", char)
            }
        }
    }()

    //声明一个匿名函数，并创建一个goroutine
    go func() {
      defer wg.Done()

      //显示字母表三次
      for count := 0;count < 3; count++ {
          for char := 'A'; char < 'A'+26; char++ {
              fmt.Printf("%c ",char)
          }
      }
    }()
    // 等待goroutine结束
    fmt.Println("Waiting To Finish...")

    wg.Wait()

    fmt.Println("\nTerminating Program...")

}
