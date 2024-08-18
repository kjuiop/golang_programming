package types

type BodyRoomReq struct {
	Name string `json:"name" binding:"required"`
}

type FormRoomReq struct {
	Name string `json:"name" binding:"required"`
}
