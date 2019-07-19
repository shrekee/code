package ipc

import "encoding/json"

type IpcClient struct {
	conn chan string
}

func NewIpcClient(server *IpcServer) *IpcServer {
	c := server.Connect()

	return &IpcServer{c}

}

func (client *IpcServer)Call(method,params string)(resp *Response,err error)  {
	req := &Request{method,params}

	var b []byte
	b,err =json.Marshal(req)
	if err != nil {
		return
	}
	client.conn <- string(b)
	str := <-client.conn

}

