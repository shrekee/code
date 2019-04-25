package test

import "fmt"

func main() {

	fmt.Println("hello, world")
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)
}
