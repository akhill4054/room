package schemas

import "time"

type CreateRoomRequestSchema struct {
	Name        string `json:"name" binding:"required"`
	PicUrl      string `json:"pic_url,omitempty"`
	Description string `json:"description" binding:"required"`
	IsPrivate   bool   `json:"is_private" binding:"boolean"`
}

type RoomSchema struct {
	ID          int       `json:"id" binding:"required"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description"`
	IsPrivate   bool      `json:"is_private" binding:"required"`
	PicUrl      string    `json:"pic_url,omitempty"`
	OwnerUid    int       `json:"owner_uid" binding:"required"`
	CreatedAt   time.Time `json:"created_at" binding:"required"`
}

type RoomOwner struct {
	UID      int    `json:"uid" binding:"required"`
	Username string `json:"username" binding:"required"`
}

type GetRoomsResponse struct {
	Rooms []RoomSchema
}

type CreateRoomResponse struct {
	RoomSchema
}

type UpdateRoomResponse struct {
	CreateRoomRequestSchema
}

type DeleteRoomResponse struct {
	Message string `json:"message" binding:"required"`
}
