package tools

import (
	"github.com/tatsushid/go-fastping"
	"net"
	"fmt"
	"time"
)

/**
 * author       : liwenqiang
 *creating_time : 19-7-24 下午2:46
 * file_name    : ping.go.py
 * IDE          : GoLand
**/
func Ping(parms string) error {
	p := fastping.NewPinger()
	ra, err := net.ResolveIPAddr("ip4:icmp", parms)
	fmt.Println("ip: ", ra)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	p.AddIPAddr(ra)
	p.OnRecv = func(addr *net.IPAddr, rtt time.Duration) {
		fmt.Printf("IP Addr: %s receive, RTT: %v\n", addr.String(), rtt)
	}
	p.OnIdle = func() {
		fmt.Println("Finish")
	}
	err = p.Run()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil

}
