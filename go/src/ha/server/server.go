package main

/**
 * author       : liwenqiang
 *creating_time : 19-7-22 上午10:54
 * file_name    : server.py
 * IDE          : GoLand
**/
import (
	"fmt"

	"net"
	"io"
	"bytes"
)

func main() {
	r := []byte("hello,world")
	fmt.Println(r)
	fmt.Println([]byte("haha"))
	server, err := net.Listen("tcp", ":30000")
	for {
		conn, err := server.Accept()
		conn.SetDeadline()
		fmt.Printf("Connected to : %s\n", conn.RemoteAddr().Network())
		go connHandler(conn)

	}
}

func connHandler(conn net.Conn) {
	for {
		fmt.P
		var buf [512]byte
		cnt, err := conn.Read(buf[0:])
		cntString := string(cnt)

	}

}
func HandleAgent() {
	fmt.Println("pass...")
}

}

func ReadAll(conn net.Conn) ([]byte, error) {
	//defer conn.Close()

	result := bytes.NewBuffer(nil)
	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		result.Write(buf[0:n])
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
	}
	return result.Bytes(), nil
}

func register(addr string) {

}
