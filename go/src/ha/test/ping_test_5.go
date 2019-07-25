package main

/**
 * author       : liwenqiang
 *creating_time : 19-7-25 上午10:13
 * file_name    : ping_test_5.py
 * IDE          : GoLand
**/
import (
	"os/exec"
	"fmt"

	"strings"
)

func main() {
	Ping("192.168.123.101")

}
func Ping(dst string)  {
	//out, _ := exec.Command("ping", dst, "-c 2", "-i 2", "-w 5").Output()
	out, _ := exec.Command("ping", dst, "-c 5", "-w 10").Output()
	//out, _ := exec.Command("ping", "xiaona.me", "-c 3", "-w 10").Output()
	//out, _ := exec.Command("ping", "www.baidu.com", "-c 2", "-i 2", "-w 5").Output()
	fmt.Println("out: ", out)
	fmt.Println("out: ", string(out))
	if len(out) == 0 || strings.Contains(string(out), "Destination Host Unreachable") {
		fmt.Println("TANGO DOWN")
	} else {
		fmt.Println("IT'S ALIVEEE")
	}

}