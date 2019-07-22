package main

import (
	"net"
	"time"
	"log"
)

/**
 * author       : liwenqiang
 *creating_time : 19-7-22 下午5:22
 * file_name    : socker_test_01.py
 * IDE          : GoLand
**/
//长连接入口
func handleConnection(conn net.Conn,timeout int) {

	buffer := make([]byte, 2048)
	for {
		n, err := conn.Read(buffer)

		if err != nil {
			LogErr(conn.RemoteAddr().String(), " connection error: ", err)
			return
		}
		Data :=(buffer[:n])
		messnager := make(chan byte)
		//心跳计时
		go HeartBeating(conn,messnager,timeout)
		//检测每次Client是否有数据传来
		go GravelChannel(Data,messnager)
		Log( "receive data length:",n)
		Log(conn.RemoteAddr().String(), "receive data string:", string(Data))

	}
}

//心跳计时，根据GravelChannel判断Client是否在设定时间内发来信息
func HeartBeating(conn net.Conn, readerChannel chan byte,timeout int) {
	select {
	case fk := <-readerChannel:
		Log(conn.RemoteAddr().String(), "receive data string:", string(fk))
		conn.SetDeadline(time.Now().Add(time.Duration(timeout) * time.Second))
		//conn.SetReadDeadline(time.Now().Add(time.Duration(5) * time.Second))
		break
	case <-time.After(time.Second*5):
		Log("It's really weird to get Nothing!!!")
		conn.Close()
	}

}

func GravelChannel(n []byte,mess chan byte){
	for _ , v := range n{
		mess <- v
	}
	close(mess)
}


func Log(v ...interface{}) {
	log.Println(v...)
}
