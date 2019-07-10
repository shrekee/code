package main

import (
	"fmt"
	"sync"
	"runtime"
)

var wg sync.WaitGroup

func main() {
	runtime.GOMAXPROCS(2)

	wg.Add(2)
	fmt.Println("Create goroutines")
	go printPrinme("A")
	go printPrinme("B")

	fmt.Println("Waiting To Finish.")

	wg.Wait()
	fmt.Println("Teminating Program..")


}

func printPrinme(prefix string) {
	defer wg.Done()

	next:
		for outer :=2;outer < 5000; outer++ {
			for inner :=2; inner < outer; inner++ {
				if outer % inner == 0 {
					continue next
				}
			}
			fmt.Printf("%s:%d\n",prefix, outer)
		}
		fmt.Println("completed", prefix)
}
