package schemas

type CreateRoomRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	IsPrivate   bool   `json:"is_private" binding:"required"`
}

type RoomDetailsResponse struct {
	Id            string `json:"id" binding:"required"`
	Name          string `json:"name" binding:"required"`
	Description   string `json:"description" binding:"required"`
	IsPrivate     bool   `json:"is_private" binding:"required"`
	PicUrl        string `json:"pic_url" binding:"required"`
	AdminUsername string `json:"admin_username" binding:"required"`
}

type GetRoomsResponse struct {
	Rooms []RoomDetailsResponse
}

type CreateRoomResponse struct {
	RoomDetailsResponse
}

type UpdateRoomResponse struct {
	RoomDetailsResponse
}

type DeleteRoomResponse struct {
	Message string `json:"message" binding:"required"`
}
