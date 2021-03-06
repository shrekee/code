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
	"os/exec"
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
	pingChan   = make(chan string)
	clientChan = make(chan string)
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
	clusterMem = append(clusterMem, "192.168.1.1", "192.168.1.2",
		"127.0.0.1", "192.168.123.1", "192.168.123.101", "192.168.123.102")
	// HA集群主机成员
	HAHosts = append(HAHosts, "127.0.0.1", "192.168.1.1")
	//clusterMem = append(clusterMem, "192.168.1.1", "192.168.1.2")

	server, err := net.Listen("tcp", ":30000")
	if err != nil {
		Log("Fatal createServer errors: ", err.Error())
		return
	}
	defer server.Close()
	// tcp Server主要逻辑线程: 用于接受clients的socket连接请求，并处理之间的信息沟通
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
			mutex1.Unlock()

			// 把主机添加进 “在线主机列表”； 同时从“超时”主机立列表中删除此主机，如果存在的话
			if !isExisthost {
				mutex1.Lock()
				onlineHosts = append(onlineHosts, remoteAddrIP)
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
				mutex1.Unlock()
				fmt.Printf("Connected to : %s\n", conn.RemoteAddr().String())

				// queue := make(chan string, 100)
				// 与每个client sock处理逻辑的线程，负责主要的沟通逻辑
				go connHandler(conn)
			}
		}
	}()

	// 线程： 用于交互命令行输入
	go readLine()

	//go ping(pingChan)
	go handleSockTimeOut()
	// 阻塞主线程
	select {}

}

func connHandler(conn net.Conn) {
	//remoteAddr := conn.RemoteAddr()

	// 一个socket的是否存活的标志变量
	sockIsAlive := true

	queueHandle := make(chan string)
	buf := make([]byte, 512)

	// 处理写入的线程
	go func() {
		for {
			// 如果socket 连接关闭，则关闭此线程
			if !sockIsAlive {
				return
			}

			command := <-clientChan
			log.Println(command)
			conn.Write([]byte(command))

			// 睡眠一秒，目标：让clientQueue的item发送到每个接受的线程中，刚好每个线程一个。
			time.Sleep(1e9)
		}
	}()

	// 处理读取的线程
	go func(queue chan string) {
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
	}(queueHandle)

	// 用以接受client端的心跳信息
	go HeartBeating(conn, queueHandle, &sockIsAlive)

}

func HeartBeating(conn net.Conn, readerChannel chan string, sockAlive *bool) {

	var deltaTime int = 0
	remoteAddr := conn.RemoteAddr()
	remoteAddIP := strings.Split(remoteAddr.String(), ":")[0]
	for {
		select {
		case r := <-readerChannel:
			deltaTime = 0
			rec := string(r)
			if rec == "1" {
				Log("Received heart beating:", remoteAddr.String())
				conn.SetDeadline(time.Now().Add(time.Duration(sockTimeout) * time.Second))
			} else {
				recSlice := strings.Split(rec, ":")
				if len(recSlice) > 1 {
					switch recSlice[0] {
					case "ping":
						fmt.Println("=====Ping Results=====")
						if strings.HasSuffix(r, "failed") {
							Log("Error: ", remoteAddIP, "->", rec)
						} else {
							Log("Success: ", remoteAddIP, "->", rec)
						}
						fmt.Println(rec)
					case "destory":
						fmt.Println("======destory")
						fmt.Println(rec)
					default:
						fmt.Println("======default")
						fmt.Println(rec)
					}
				}
			}
		case <-time.After(time.Second * sockCheckInterval):
			deltaTime += sockCheckInterval
			Log("Warning: Time elapsed ", deltaTime, "seconds, and no heartbeat for", remoteAddr.String())
			if deltaTime >= sockTimeout {
				// 断开超时的socket连接
				conn.Close()
				*sockAlive = false
				Log("Error: connection timeout: socket", remoteAddr.String())
				Log("Disconnected: after elapsed time"+
					"", sockTimeout, "seconds with no heartbeat")
				// 清理环境："丢失主机切片"和"在线主机切片"
				isExistHost := false
				mutex1.Lock()
				for _, v := range timeoutHosts {
					if v == remoteAddIP {
						isExistHost = true
						break
					}
				}
				mutex1.Unlock()
				if !isExistHost {
					mutex1.Lock()
					timeoutHosts = append(timeoutHosts, strings.Split(remoteAddr.String(), ":")[0])
					mutex1.Unlock()
				}

				mutex1.Lock()
				for k, v := range onlineHosts {
					if remoteAddIP == v {
						var err error
						onlineHosts, err = tools.StrSRemove(onlineHosts, k)
						if err != nil {
							Log("Error: ===================")
							Log("Fatal error: ", err.Error())
						}
					}
				}
				mutex1.Unlock()

				// 把连接超时的远端ip地址发给server的主线程
				pingChan <- remoteAddIP
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
			fmt.Println("online hosts:", onlineHosts)
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

func handleSockTimeOut() {
	for {
		params := <-pingChan
		exitCode := Ping(params)
		if exitCode == 1 {
			// failed.
			Log("Error: serverHost ping host failed:", params)
			Log("=====onlineHost is:", onlineHosts)
			for i := 0; i < len(onlineHosts); i++ {
				for j := 0; j < 4; j++ {
					clientChan <- ("ping:" + params)
					Log("======================================")
					Log("Send Num_: ", j+1)
					Log("======================================")
					time.Sleep(time.Millisecond * 100)
				}
			}
		} else {
			// ping success.
			Log("==============================================")
			Log("Warning: Host does not have heart beat:", params)
			Log("==============================================")

			// 下一步： 通过主机的ssh，连接到该client，重启该客户端一次
			// 。。。。。

		}
	}
}

func Ping(dst string) int {
	// return 0, if ping success
	// return 1, if ping failed
	Log("Try to ping: ", dst)
	out, _ := exec.Command("ping", dst, "-c 5", "-w 10").Output()
	fmt.Println("out: ", string(out))
	if len(out) == 0 || strings.Contains(string(out), "0 received") {
		// failed
		return 1
	} else {
		// success
		return 0
	}
}
