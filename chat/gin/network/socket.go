package network

import (
	"chat-gin/types"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = &websocket.Upgrader{
	ReadBufferSize:  types.SocketBufferSize,
	WriteBufferSize: types.MessageBufferSize,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Room struct {
	Forward chan *message // 수신되는 메시지를 보관하는 값
	// 들어오는 메시지를 다른 클라이언트에게 전송을 합니다.

	Join  chan *Client // Socket 이 연결되는 경우에 적용
	Leave chan *Client // Socket 이 끊어지는 경우에 대해서 적용

	Clients map[*Client]bool // 현재 방에 있는 Client 정보를 저장
}

func NewRoom() *Room {
	return &Room{
		Forward: make(chan *message),
		Join:    make(chan *Client),
		Leave:   make(chan *Client),
		Clients: make(map[*Client]bool),
	}
}

func (r *Room) SocketServe(c *gin.Context) {

	_, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		panic(err)
	}

}

type message struct {
	Name    string
	Message string
	Time    int64
}

type Client struct {
	Send   chan message
	Room   *Room
	Name   string
	Socket *websocket.Conn
}
