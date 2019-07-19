package test01

import "fmt"
import "strings"

func main() {
	//#var p *int = new(int)

	str := "hello, 世界"

	for n,i := range str {
		fmt.Println(n,i)
		fmt.Printf("%d %c\n",n,i)
	}

	fmt.Println(strings.Repeat("*", 30))
	var arr [10] int
	arr[0] = 1
	for n,i := range arr {
		fmt.Println(n,i)
	}
	var arr2 = [...]float32{12.22,232,4445,666.0,1231}
	fmt.Println(strings.Repeat("*", 30))


	for n,i := range arr2{
		fmt.Println(n,i)
	}

	var slice = []string{"1111","22222","33333","444444"}

	fmt.Printf("slice len is: %d\n",len(slice))

	 slice = append(slice, "haha")
	fmt.Printf("slice len is: %d\n",len(slice))

	for n,i := range slice {
		fmt.Println(n,i)
	}





}
