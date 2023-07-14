package controllers

import (
	"net/http"
	"strconv"

	"github.com/akhill4054/room-backend/models"
	"github.com/akhill4054/room-backend/pkg/util"
	"github.com/akhill4054/room-backend/schemas"
	"github.com/gin-gonic/gin"
)

func CreateRoom(c *gin.Context) {
	uid := c.Keys["user"].(*models.User).ID

	var reqBody schemas.CreateRoomRequestSchema

	util.Log.Info(strconv.FormatBool(reqBody.IsPrivate))
	if err := c.BindJSON(&reqBody); err != nil {
		c.IndentedJSON(http.StatusBadRequest, schemas.ErrorResponse{Message: err.Error()})
		return
	}

	room, _ := models.CreateRoom(
		reqBody.Name,
		reqBody.PicUrl,
		uid,
		reqBody.Description,
		reqBody.IsPrivate,
	)

	c.IndentedJSON(http.StatusCreated, asRoomSchema(room))
}

func GetRoom(c *gin.Context) {
	roomId, err := strconv.Atoi(c.Param("roomId"))
	if err != nil {
		response := schemas.ErrorResponse{Message: "Invalid room id"}
		c.IndentedJSON(http.StatusBadRequest, response)
		return
	}

	room, err := models.GetRoom(roomId)
	if err != nil {
		if err == models.ErrRoomNotFound {
			c.IndentedJSON(http.StatusNotFound, err.Error())
			return
		} else {
			panic(err)
		}
	}

	c.IndentedJSON(http.StatusOK, asRoomSchema(room))
}

func GetRooms(c *gin.Context) {
	user := c.Keys["user"].(*models.User)

	ownerUid := c.Query("ouid")

	size, _ := strconv.Atoi(c.DefaultQuery("size", "10"))
	after, _ := strconv.Atoi(c.DefaultQuery("after", "0"))

	rooms, err := models.GetRooms(user, ownerUid, size, after)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, schemas.ErrorResponse{Message: err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, asRoomsSchema(rooms))
}

func UpdateRoom(c *gin.Context) {
	var reqBody schemas.UpdateRoomResponse

	if err := c.BindJSON(&reqBody); err != nil {
		c.IndentedJSON(http.StatusBadRequest, schemas.ErrorResponse{Message: err.Error()})
		return
	}

	roomId, err := strconv.Atoi(c.Param("roomId"))
	if err != nil {
		response := schemas.ErrorResponse{Message: "Invalid room id"}
		c.IndentedJSON(http.StatusBadRequest, response)
		return
	}

	room, err := models.GetRoom(roomId)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, schemas.ErrorResponse{Message: err.Error()})
		return
	}

	room.Name = reqBody.Name
	room.PicUrl = reqBody.PicUrl
	room.Description = reqBody.Description
	room.IsPrivate = reqBody.IsPrivate

	room, err = room.Update()

	c.IndentedJSON(http.StatusOK, asRoomSchema(room))
}

func DeleteRoom(c *gin.Context) {
	user := c.Keys["user"].(*models.User)

	roomId, err := strconv.Atoi(c.Param("roomId"))
	if err != nil {
		response := schemas.ErrorResponse{Message: "Invalid room id"}
		c.IndentedJSON(http.StatusBadRequest, response)
		return
	}

	if err := models.DeleteRoom(roomId, user); err != nil {
		var code int

		if err == models.ErrRoomNotFound {
			code = http.StatusNotFound
		}
		if err == models.ErrRoomDeleteNotAllowed {
			code = http.StatusBadRequest
		}

		c.IndentedJSON(code, schemas.ErrorResponse{Message: err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, schemas.DeleteRoomResponse{Message: "Room deleted"})
}

func asRoomSchema(room *models.Room) *schemas.RoomSchema {
	return &schemas.RoomSchema{
		ID:          room.ID,
		Name:        room.Name,
		Description: room.Description,
		IsPrivate:   room.IsPrivate,
		OwnerUid:    room.OwnerUID,
		CreatedAt:   room.CreatedAt,
	}
}

func asRoomsSchema(rooms *[]models.Room) *[]schemas.RoomSchema {
	roomScheams := []schemas.RoomSchema{}
	for _, room := range *rooms {
		roomScheams = append(roomScheams, *asRoomSchema(&room))
	}

	return &roomScheams
}
