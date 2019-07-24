package main

import (
	"net"
	"fmt"
	//"io"
	"time"
	//"go/constant"
	//"bytes"
	"log"
	"sync"
	"strings"
	"bufio"
	"os"
	"ha/tools"
)

/**
 * author       : liwenqiang
 *creating_time : 19-7-22 上午10:54
 * file_name    : server.py
 * IDE          : GoLand
**/

/*
集群必要条件
	1、集群server和Host——client的时间同步
	2、DNS域名解析
	3、ssh免密登陆
	4、
 */

/*
yaml配置文件
cluster
	cluster1:
	- name: mycluster1
	- socketTimeOutinSecond: 20
	- heartBeatCheckIntervalinSecond: 2
		hosts:
		  - ip: 1.1.1.1
		  - port:
		  host1:
		  - name: hp01
		  - ha: yes
		  - vms:
			  vm01:
			  - name: vmtest11
			  - ip: 1.1.1.2
			  - ha: yes

	cluster2:
	- name: mycluster2
	- socketTimeOutinSecond: 20
	- heartBeatCheckIntervalinSecond: 2
	- hosts
		host

 */

// 集群全局设置
var (
	clusterMem   = make([]string, 0, 10)
	timeoutHosts = make([]string, 0, 100)
	onlineHosts  = make([]string, 0, 100)
	//lostHosts    = make([]string, 0, 100)
	HAHosts       = make([]string, 0, 100)
	activeHAHosts = make([]string, 0, 100)
	//  channel
	pingChan = make(chan string)
	clientChan= make(chan string)
)

const (
	//集群的配置参数
	// in second
	sockTimeout       = 20
	sockCheckInterval = 2
)

// 锁：1、在操作超时连接主机列表，防止竞争
var mutex1 sync.Mutex
//var err error
//var wg sync.WaitGroup
func main() {

	// 通过读取配置文件，获取集群的主机列表
	// 目标是读取yaml格式的文件，获取所有主机的信息
	// 通过第三方模块 spf13/viper
	//初始集群主机成员
	clusterMem = append(clusterMem, "192.168.1.1", "192.168.1.2", "127.0.0.1","192.168.123.1")
	// HA集群主机成员
	HAHosts = append(HAHosts, "127.0.0.1", "192.168.1.1")
	//clusterMem = append(clusterMem, "192.168.1.1", "192.168.1.2")

	server, err := net.Listen("tcp", ":30000")
	if err != nil {
		Log("Fatal createServer errors: ", err.Error())
		return
	}
	// 线程: 用于接受socket连接
	go func() {
		for {
			conn, err := server.Accept()
			if err != nil {
				Log("Fatal: connection errors: ", err.Error())
				continue
			}
			// 判断此socket连接是否在集群主机列表中?
			hostAddr := conn.RemoteAddr().String()
			remoteAddrIP := strings.Split(hostAddr, ":")[0]
			isValidHost := false
			for _, v := range clusterMem {
				if remoteAddrIP == v {
					isValidHost = true
					break
				}
			}
			if !isValidHost {
				fmt.Println("Warning: the remoteHostAddr", hostAddr, "is not valid host! disconnected")
				conn.Close()
				continue
			}
			// 把给主机ip添加进在线主机列表
			isExisthost := false
			mutex1.Lock()
			{
				// 添加新的主机host的IP,加入已经连接，则提示已经连接，并断开此新的连接
				for _, v := range onlineHosts {
					if v == remoteAddrIP {
						isExisthost = true
						Log("Warning: ", remoteAddrIP, " already in this cluster")
						Log("Disconnecting....")
						conn.Close()
						break
					}
				}
				// 删除超时连接的host的IP
				for k, v := range timeoutHosts {
					if v == remoteAddrIP {
						timeoutHosts, err = tools.StrSRemove(timeoutHosts, k)
						if err != nil {
							Log("Error: ", err.Error())
						}
						break
					}
				}
			}
			mutex1.Unlock()

			if !isExisthost {
				mutex1.Lock()
				onlineHosts = append(onlineHosts, remoteAddrIP)
				mutex1.Unlock()
				fmt.Printf("Connected to : %s\n", conn.RemoteAddr().String())
				//queue := make(chan string, 100)
				go connHandler(conn)
				isExisthost = false
			}
		}
	}()

	// 线程: 用于打印超时的连接.
	//go func() {
	//	for  {
	//		fmt.Println("TimeoutConnetions: ",timeoutHosts)
	//		time.Sleep(time.Duration(2)*time.Second)
	//
	//	}
	//}()
	// 线程： 用于交互命令行输入
	go readLine()
	//go ping(pingChan)
	go handleSockTimeOut(pingChan,clientChan)
	// 阻塞主线程
	select {}

}

func connHandler(conn net.Conn) {
	//remoteAddr := conn.RemoteAddr()

	queueHandle := make(chan string)
	buf := make([]byte, 512)

	go func(clientQueue,queue chan string) {
		// 处理写入的线程
		go func() {
			command:=<-clientQueue
			conn.Write([]byte(command))
		}()
		for {
			//var buf [512]byte
			n, err := conn.Read(buf)
			if err != nil {
				Log("Fatal errors: ", err.Error())
				conn.Close()
				return
			}
			cntString := string(buf[0:n])
			queue <- cntString
		}
	}(clientChan,queueHandle)

	go HeartBeating(conn, queueHandle)

}

func HeartBeating(conn net.Conn, readerChannel chan string) {

	var deltaTime int = 0
	remoteAddr := conn.RemoteAddr()
	remoteAddIP := strings.Split(remoteAddr.String(), ":")[0]
	for {
		select {
		case r := <-readerChannel:
			deltaTime = 0
			rec:=string(r)
			if rec=="1"{
				Log("Received heart beating:",remoteAddr.String())
				conn.SetDeadline(time.Now().Add(time.Duration(sockTimeout) * time.Second))
			}else {
				recSlice:=strings.Split(rec,":")
				if len(recSlice)>1{
					switch recSlice[0] {
					case "ping":
						fmt.Println("ping")
						fmt.Println(rec)
					case "destory":
						fmt.Println("destory")
						fmt.Println(rec)
					default:
						fmt.Println("default")
						fmt.Println(rec)
					}
				}
			}
		case <-time.After(time.Second * sockCheckInterval):
			deltaTime += sockCheckInterval
			Log("Warning: Time elapsed ", deltaTime, "seconds, and no heartbeat for", remoteAddr.String())
			if deltaTime >= sockTimeout {
				Log("Error: Connection timeout: socket", remoteAddr.String(), "is disconnected after elapsed time"+
					"", sockTimeout, "seconds with no heartbeat")
				//把连接超时的远端ip地址发给server的主线程
				pingChan<-remoteAddIP
				isExistHost := false
				mutex1.Lock()
				{
					for _, v := range timeoutHosts {
						if v == remoteAddIP {
							isExistHost = true
							break
						}
					}
					if !isExistHost {
						timeoutHosts = append(timeoutHosts, strings.Split(remoteAddr.String(), ":")[0])
					}

					for k, v := range onlineHosts {
						if remoteAddIP == v {
							var err error
							onlineHosts, err = tools.StrSRemove(onlineHosts, k)
							if err != nil {
								fmt.Println("Fatal error: ", err.Error())
								return
							}
						}
					}
				}
				mutex1.Unlock()
				/*
				do something to tell server: this remote connection is closed~~
				 */
				conn.Close()
				return
			}
		}
	}
}

func Log(v ...interface{}) {
	log.Println(v...)
}

func readLine() {
	//commands :=[]string{"status","list"}
	var (
		inputReader *bufio.Reader
		input       string
		err         error
	)

	for {
		inputReader = bufio.NewReader(os.Stdin)
		fmt.Print("input >: ")
		input, err = inputReader.ReadString('\n')
		if err != nil {
			fmt.Println("input error: ", err.Error())
			continue
		}
		if len(input) <= 1 {
			continue
		}
		inputS2 := strings.Fields(strings.TrimSpace(input))
		if len(inputS2) == 0 {
			continue
		}
		switch inputS2[0] {
		case "status":
			fmt.Println("status: ", timeoutHosts)
		case "list":
			fmt.Println("list: ", timeoutHosts)
		case "cluster":
			fmt.Println("cluste members:", clusterMem)
			fmt.Println("online hosts:", onlineHosts)
			fmt.Println("timeout hosts", timeoutHosts)
		case "ha":
			fmt.Println("ha members:", HAHosts)
			fmt.Println("ha online hosts", activeHAHosts)

		default:
			fmt.Println(`Usage: <command> <patametes>
							<status>
							<list>
							<ha>
							<cluster>`)
		}
	}
}

func ping(ch chan string) error {
	params := <-ch
	err := tools.Ping(params)
	if err != nil {
		fmt.Println("Error: ", err.Error())
		return err
	}
	return nil
}

func handleSockTimeOut(ch,clientChan chan string){
	for  {
		params := <-ch
		err := tools.Ping(params)
		if err!=nil{
			fmt.Println("Error: ", err.Error())
			for i:=0;i<len(onlineHosts) ;i++  {
				clientChan<-params
			}
		}
	}
}


