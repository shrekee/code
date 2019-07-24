package main

/**
 * author       : liwenqiang
 *creating_time : 19-7-24 下午2:25
 * file_name    : ping_test_3.py
 * IDE          : GoLand
**/
import (
	"fmt"
	"github.com/tatsushid/go-fastping"
	"net"
	"os"
	"time"
)

func main() {
	p := fastping.NewPinger()
	ra, err := net.ResolveIPAddr("ip4:icmp", os.Args[1])
	fmt.Println("ip: ", ra)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	p.AddIPAddr(ra)
	p.OnRecv = func(addr *net.IPAddr, rtt time.Duration) {
		fmt.Printf("IP Addr: %s receive, RTT: %v\n", addr.String(), rtt)
	}
	p.OnIdle=func(){
		fmt.Println("Finish")
	}
	err=p.Run()
	if err!=nil{
		fmt.Println(err.Error())
		os.Exit(2)
	}

}
