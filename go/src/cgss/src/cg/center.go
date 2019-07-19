package cg

/**
 * author       : liwenqiang
 *creating_time : 19-7-16 下午5:14
 * file_name    : center.py
 * IDE          : GoLand
**/
import (
	"cgss/src/ipc"
	"sync"
	"encoding/json"
	"github.com/pkg/errors"
)

var _ ipc.Server = &CenterServer{}

type Message struct {
	From    string "from"
	To      string "to"
	Content string "content"
}

type CenterServer struct {
	servers map[string]ipc.Server
	players []*Player
	rooms   []*Room
	mutex   sync.RWMutex
}

func NewCenterServer() *CenterServer {
	servers := make(map[string]ipc.Server)
	players := make([]*Player, 0)

	return &CenterServer{servers: servers, players: players}
}

func (server *CenterServer) addPlayer(params string) error {
	player := NewPlayer()
	err := json.Unmarshal([]byte(params), &player)
	if err != nil {
		return err
	}

	server.mutex.Lock()
	defer server.mutex.Unlock()
	server.players = append(server.players, player)
	return nil
}

func (server *CenterServer) removePlayer(params string) error {
	server.mutex.Lock()
	defer server.mutex.Unlock()

	for i, v := range server.players {
		if v.Name == params {
			if i == len(server.players)-1 {
				server.players = server.players[:i]
			} else {
				server.players = append(server.players[:i], server.players[i+1:]...)
			}

			return nil
		}
	}
	return errors.New("Player not found.")
}

func (server *CenterServer)listPlayer(params string)(players string,err error)  {
	server.mutex.Lock()
	defer server.mutex.Unlock()

	if len(server.players) >0 {
		b,_ := json.Marshal(server.players)
		players = string(b)
	}else{
		err = errors.New("No player online.")
	}
	return
}

func (server *CenterServer)broadcast(params string) error {

	var message Message
	err:= json.Unmarshal([]byte(params),&message)
	if err !=nil{
		return err
	}

	server.mutex.Lock()
	defer server.mutex.Unlock()

	if len(server.players) >0 {
		for _,player := range server.players{
			player.mq <- &message
		}
	}else{
		err =errors.New("No player online.")
	}
	return err
}

func (server *CenterServer)Handle(method,params string) *ipc.Response  {
	switch method {
	case "addplayer":
		err := server.addPlayer(params)
		if err != nil{
			return &ipc.Response{Code:err.Error()}
		}
	case "removeplayer":
		err := server.removePlayer(params)
		if err != nil{
			return &ipc.Response{Code:err.Error()}
		}
	case "listplayer":
		players,err :=server.listPlayer(params)
		if err!=nil{
			return &ipc.Response{Code:err.Error()}
		}
		return &ipc.Response{"200",players}
	case "broadcast":
		err:=server.broadcast(params)
		if err!=nil{
			return &ipc.Response{Code:err.Error()}
		}
		//return &ipc.Response{Code:"200"}
	default:
		return &ipc.Response{Code:"404",Body:method}
	}
	return &ipc.Response{Code:"200"}
}

func (server *CenterServer)Name()string  {
	return "CenterServer"
}














