package test01
import (
	"time"
	"fmt"
)

// 生成2, 3, 4, ... 到 channel 'ch'中.

func Generate(ch chan<- int) {

	for i := 2; ; i++ {

		ch <- i	// Send 'i' to channel 'ch'.

	}

}

// 从管道复制值 'in' 到 channel 'out',

// 移除可整除的数 'prime'.

func Filter(in <-chan int, out chan<- int, prime int) {

	for {

		i := <-in	// 接收值 'in'.

		if i%prime != 0 {

			out <- i	// 传入 'i' 到 'out'.

		}

	}

}

func main() {
	start_time := time.Now()

	ch := make(chan int)	// Create a newchannel.

	go Generate(ch)	// Launch Generate goroutine.

	for i := 0; i < 33860; i++ {

		prime := <-ch

		print(prime, "\n")

		ch1 := make(chan int)

		go Filter(ch, ch1, prime)

		ch = ch1

	}
	end_time := time.Now()
	//fmt.Printf("\n总共耗时：%f 秒\n在%d 个自然数中有%d 个素数\n")
	fmt.Printf("\n总共耗时：%f 秒",end_time.Sub(start_time).Seconds())

}
