package schemas

import (
	"time"

	"github.com/akhill4054/room-backend/models"
)

type CreateRoomRequestSchema struct {
	Name        string `json:"name" binding:"required"`
	PicUrl      string `json:"pic_url,omitempty"`
	Description string `json:"description" binding:"required"`
	IsPrivate   bool   `json:"is_private" binding:"boolean"`
}

type RoomSchema struct {
	Id          int       `json:"id" binding:"required"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description"`
	IsPrivate   bool      `json:"is_private" binding:"required"`
	PicUrl      string    `json:"pic_url,omitempty"`
	OwnerUid    int       `json:"owner_uid" binding:"required"`
	CreatedAt   time.Time `json:"created_at" binding:"required"`
}

type RoomOwner struct {
	Uid      int    `json:"uid" binding:"required"`
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

type CreateRoomMemberRequest struct {
	MemberUid int                   `json:"member_uid" binding:"required"`
	Role      models.RoomMemberRole `json:"role" binding:"required"`
}

type CreateRoomMemberResponse struct {
	Message string `json:"message" binding:"required"`
}

type JoinRoomResponse struct {
	Message string `json:"message" binding:"required"`
}

type DeleteRoomResponse struct {
	Message string `json:"message" binding:"required"`
}

type DeleteRoomMemberResponse struct {
	Message string `json:"message" binding:"required"`
}
