package main

/**
 * author       : liwenqiang
 *creating_time : 19-7-24 上午9:32
 * file_name    : bufio_test_01.py
 * IDE          : GoLand
**/
import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"golang.org/x/text/cases"
)

var(
	inputReader *bufio.Reader
	input1 string
	err error
)

func main()  {
	for {
		inputReader = bufio.NewReader(os.Stdin)
		fmt.Print("input >: ")
		input1, err = inputReader.ReadString('\n')
		if err != nil {
			fmt.Println("input error: ", err.Error())
			continue
		}
		//fmt.Print("len is: ",len(input1),input1)
		if len(input1) <= 1{
			continue
		}
		//inputS :=strings.Split(input1," ")
		inputS2 := strings.Fields(strings.TrimSpace(input1))
		if len(inputS2) == 0{
			continue
		}
		////for _,v := range inputS{
		////	fmt.Printf("The input was : %s\n",v)
		//}
		//fmt.Println(strings.TrimSpace(input1))
		//fmt.Printf("len is %d,\ncontent is:%s",len(inputS2),inputS2)
		for _, v := range inputS2 {
			fmt.Printf("The input was : %s\n", v)
		}
	}
	select {}
	//readLine()

}

func readLine()  {
	//commands :=[]string{"status","list"}
	for {
		inputReader = bufio.NewReader(os.Stdin)
		fmt.Print("input >: ")
		input1, err = inputReader.ReadString('\n')
		if err != nil {
			fmt.Println("input error: ", err.Error())
			continue
		}
		if len(input1) <= 1{
			continue
		}
		inputS2 := strings.Fields(strings.TrimSpace(input1))
		if len(inputS2) == 0{
			continue
		}
		switch inputS2[0]{
		case "status":
			fmt.Println("status: ",timeoutConn)
		case "list":
			fmt.Println("list: ",timeoutConn)
		default:
			fmt.Println("Usage <command><patametes>" +
				"<status>...")
		}
	}

}
