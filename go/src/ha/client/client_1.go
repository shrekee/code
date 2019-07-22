package client

/**
 * author       : liwenqiang
 *creating_time : 19-7-22 上午11:21
 * file_name    : client_1.py
 * IDE          : GoLand
**/
import (
	"fmt"
	"os"
	"net"
	"bytes"
	"io"
)

func clientAgent() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
		os.Exit(1)
	}

}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
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
