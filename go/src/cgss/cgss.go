package main

/**
 * author       : liwenqiang
 *creating_time : 19-7-17 下午2:03
 * file_name    : cgss.py
 * IDE          : GoLand
**/
import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"cgss/src/cg"
	"cgss/src/ipc"
)

var centerClient *cg.CenterClient

func startCenterService()error  {
	server:=ipc.NewIpcServer(&cg.CenterServer{})
	client :=ipc.NewIpcClient(server)
	centerClient=&cg.CenterClient{client}

	return nil
}

func Help(args []string)  int{
	fmt.Println(`
		Commands:
			login<username><level><exp>
			logout <username>
			send<message>
			listplayer
			quit(q)
			help(h)
		`)
	return 0
}

func Quit(args []string) int {
	return 1

}

func Logout(args []string) int {
	if len(args) !=2{
		fmt.Println("USAGE: logout<username>")
		return 0
	}

	centerClient.RemovePlayer(args[1])
	return 0
}

func Login(args []string) int  {
	if len(args) != 4{
		fmt.Println("USAGE: login<username><level><exp>")
		return 0
	}

	level,err:=strconv.Atoi(args[2])
	if err!=nil{
		fmt.Println("Invalid Parametes: <level> should be an integer.")
		return

	}

}
