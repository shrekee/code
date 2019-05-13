package main

import(
	"fmt"
	"time"
	"math"
)
var i,j uint64 = 1,1
var nub uint64 = 400000
var k int = 1
//var prime []uint32
var prime []uint64 =make([ ]uint64,1)
func main() {
	prime[0] =2
	start_time := time.Now()

	for i =3; i<=nub; i+=2 {
		mid := uint64(math.Sqrt(float64(i)))
		for j =2; j<= (mid + 1); j++ {
			if i%j == 0 {
				break
			}
			if j == (mid + 1) {
				//prime[k] = i
				prime = append(prime, i)
				//k++
			}
		}
	}

	end_time := time.Now()
	diff := end_time.Sub(start_time)

	for i = 0; i< uint64(len(prime)); i++ {
		fmt.Printf("%d ",prime[i])
	}

	fmt.Printf("\n总共耗时：%f 秒\n在%d 个自然数中有%d 个素数\n",diff.Seconds(),nub,len(prime))
}
