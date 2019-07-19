package test01

import (
	"fmt"
	"strings"
)

func main() {
	str := "哈哈 ，你好，你个小煞笔。。"
	fmt.Println(strings.Contains(str, "小煞笔"))
	fmt.Println(strings.Contains(str, "xiaoshabi"))
	fmt.Println(strings.Contains("", ""))
	fmt.Println("hello")
}
