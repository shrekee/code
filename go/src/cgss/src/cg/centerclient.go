package cg

/**
 * author       : liwenqiang
 *creating_time : 19-7-17 上午11:11
 * file_name    : centerclient.py
 * IDE          : GoLand
**/
import (
	"errors"
	"encoding/json"
	"cgss/src/ipc"
)

type CenterClient struct {
	*ipc.IpcClient
}

func (client *CenterClient)Addplayer(player *Player) error  {
	b,err :=json.Marshal(*player)
	if err!=nil{
		return err
	}

	resp,err := client.Call("addplayer",string(b))
	if err==nil&&resp.Code=="200"{
		return nil
	}
	return err
}

func (client *CenterClient)RemovePlayer(name string) error  {
	ret,_ :=client.Call("removeplayer",name)
	if ret.Code =="200"{
		return nil
	}
	return errors.New(ret.Code)
}

func (client *CenterClient)ListPlayer(params string)(ps []*player,err error){
	resp,_ :=client.Call("listplayer",params)
	if resp.Code != "200"{
		err =errors.New(resp.Code)
		return
	}
	err = json.Unmarshal([]byte(resp.Body),&ps)
	return
}

func (client *CenterClient)Broadcast(message string)error  {
	m:=&Message{Content:message}
	b,err :=json.Marshal(m)
	if err!=nil{
		return err
	}

	resp,_ := client.Call("broadcast",string(b))
	if resp.Code == "200"{
		return ni
	}
	return errors.New(resp.Code)
}








