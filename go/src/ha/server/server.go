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
		  host1:
		  - name: hp01
		  - ip: 1.1.1.1
		  - port:
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
	hostCluster []string = make([]string, 0, 10)
	timeoutConn []string = make([]string, 0, 100)
	lostConn    []string = make([]string, 0, 100)
)

const (
	//集群的配置参数
	// in second
	sockTimeout       = 20
	sockCheckInterval = 2
)

var mutex1  sync.Mutex
//var wg sync.WaitGroup
func main() {

	// 通过读取配置文件，获取集群的主机列表
	// 目标是读取yaml格式的文件，获取所有主机的信息
	// 通过第三方模块 spf13/viper
	//模拟初始集群的主机为：
	hostCluster = append(hostCluster, "192.168.1.1.", "192.168.1.2")

	server, err := net.Listen("tcp", ":30000")
	if err != nil {
		Log("Fatal createServer errors: ", err.Error())
		return
	}
	go func() {
		for {
			conn, err := server.Accept()
			if err != nil {
				Log("Fatal: connection errors: ", err.Error())
				continue
			}
			fmt.Printf("Connected to : %s\n", conn.RemoteAddr().String())
			//queue := make(chan string, 100)
			go connHandler(conn, hostCluster, timeoutConn, lostConn)

		}

	}()
	// 阻塞主线程
	for  {
		read,err:=fmt.Scan()
		if err!=nil{
			fmt.Println("Fatal: ",err.Error())
		}

	}


}

func connHandler(conn net.Conn, hostCluster, timeoutConn, lostConn []string) {
	//remoteAddr := conn.RemoteAddr()

	queueHandle := make(chan string)
	buf := make([]byte, 512)

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
			//Log("got item: ", cntString)

			queue <- cntString
		}

	}(queueHandle)
	go HeartBeating(conn, queueHandle)

}

func HeartBeating(conn net.Conn, readerChannel chan string) {

	var deltaTime int = 0
	remoteAddr := conn.RemoteAddr()
	for {
		select {
		case fk := <-readerChannel:
			deltaTime = 0
			Log(remoteAddr.String(), "receive data string:", string(fk))
			conn.SetDeadline(time.Now().Add(time.Duration(sockTimeout) * time.Second))
		case <-time.After(time.Second * sockCheckInterval):
			deltaTime += sockCheckInterval
			Log("Warning: Time elapsed ", deltaTime, "seconds, and no heartbeat for", remoteAddr.String())
			if deltaTime >= sockTimeout {
				Log("Error: Connection timeout: socket", remoteAddr.String(), "is disconnected after elapsed time"+
					"", sockTimeout, "seconds with no heartbeat")
				//把连接超时的远端ip地址发给server的主线程
				mutex1.Lock()
				{
					timeoutConn = append(timeoutConn, remoteAddr.String())
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
