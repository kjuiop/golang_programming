package network

import (
	"chat-kafka/service"
	"chat-kafka/types"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
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

	Join  chan *client // Socket 이 연결되는 경우에 적용
	Leave chan *client // Socket 이 끊어지는 경우에 대해서 적용

	Clients map[*client]bool // 현재 방에 있는 Client 정보를 저장

	service *service.Service
}

func NewRoom(service *service.Service) *Room {
	return &Room{
		Forward: make(chan *message),
		Join:    make(chan *client),
		Leave:   make(chan *client),
		Clients: make(map[*client]bool),
		service: service,
	}
}

func (r *Room) SocketServe(c *gin.Context) {

	socket, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}

	userCookie, err := c.Request.Cookie("auth")
	if err != nil {
		log.Fatalln(err.Error())
	}

	client := &client{
		Socket: socket,
		Send:   make(chan *message, types.MessageBufferSize),
		Room:   r,
		Name:   userCookie.Value,
	}

	r.Join <- client

	defer func() {
		r.Leave <- client
	}()

	go client.Write()

	client.Read()

}

func (r *Room) Run() {
	// Room 에 있는 모든 채널 값들을 받는 역할
	for {
		select {
		case client := <-r.Join:
			r.Clients[client] = true
		case client := <-r.Leave:
			r.Clients[client] = false
			close(client.Send)
			delete(r.Clients, client)
		case msg := <-r.Forward:

			go r.service.InsertChatting(msg.Name, msg.Message, msg.Room)

			for client := range r.Clients {
				client.Send <- msg
			}
		}
	}
}

type message struct {
	Name    string
	Message string
	Time    int64
	Room    string
}

type client struct {
	Send   chan *message
	Room   *Room
	Name   string
	Socket *websocket.Conn
}

func (c *client) Read() {
	defer c.Socket.Close()
	// client 가 메시지를 읽는 함수

	for {
		var msg *message
		if err := c.Socket.ReadJSON(&msg); err != nil {
			if !websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				break
			} else {
				log.Fatalln(err.Error())
			}
		}
		msg.Time = time.Now().Unix()
		msg.Name = c.Name

		c.Room.Forward <- msg
	}
}

func (c *client) Write() {
	defer c.Socket.Close()
	// client 가 메시지를 전송하는 함수

	for msg := range c.Send {
		if err := c.Socket.WriteJSON(msg); err != nil {
			log.Fatalln(err.Error())
		}
	}
}
