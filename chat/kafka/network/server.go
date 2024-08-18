package network

import (
	"chat-kafka/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

type api struct {
	server *Server
}

func registerServer(s *Server) *api {

	a := &api{server: s}

	s.engine.GET("/room-list", a.roomList)
	s.engine.GET("/room", a.room)
	s.engine.GET("/enter-room", a.enterRoom)

	s.engine.POST("/make-room", a.makeRoom)

	//r := NewRoom()
	//go r.RunInit()

	//s.engine.GET("/room", r.SocketServe)

	return a
}

func (a *api) roomList(c *gin.Context) {
	if res, err := a.server.service.RoomList(); err != nil {
		response(c, http.StatusInternalServerError, err.Error())
	} else {
		response(c, http.StatusOK, res)
	}
}

func (a *api) room(c *gin.Context) {
	var req types.FormRoomReq

	if err := c.ShouldBindQuery(&req); err != nil {
		response(c, http.StatusUnprocessableEntity, err.Error())
	} else if res, err := a.server.service.Room(req.Name); err != nil {
		response(c, http.StatusInternalServerError, err.Error())
	} else {
		response(c, http.StatusOK, res)
	}
}

func (a *api) enterRoom(c *gin.Context) {
	var req types.FormRoomReq

	if err := c.ShouldBindQuery(&req); err != nil {
		response(c, http.StatusUnprocessableEntity, err.Error())
	} else if res, err := a.server.service.EnterRoom(req.Name); err != nil {
		response(c, http.StatusInternalServerError, err.Error())
	} else {
		response(c, http.StatusOK, res)
	}

}

func (a *api) makeRoom(c *gin.Context) {
	var req types.BodyRoomReq

	if err := c.ShouldBindJSON(&req); err != nil {
		response(c, http.StatusUnprocessableEntity, err.Error())
	} else if err := a.server.service.MakeRoom(req.Name); err != nil {
		response(c, http.StatusInternalServerError, err.Error())
	} else {
		response(c, http.StatusOK, "Success")
	}

}
