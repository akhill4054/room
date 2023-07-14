package models

import (
	"errors"
	"fmt"

	"github.com/akhill4054/room-backend/pkg/util"
	"gorm.io/gorm"
)

type Room struct {
	gorm.Model
	ID           int    `gorm:"primaryKey;autoIncrement:true"`
	Name         string `gorm:"not null"`
	PicUrl       string
	OwnerUID     int    `gorm:"index;not null;default: null"`
	Description  string `gorm:"not null"`
	IsPrivate    bool   `gorm:"index;not null"`
	MembersCount string `gorm:"not null"`
	Owner        User   `gorm:"foreignKey:OwnerUID;constraint:OnDelete:SET NULL;not null"`
}

func CreateRoom(
	name string,
	picUrl string,
	adminUid int,
	description string,
	isPrivate bool,
) (*Room, error) {
	room := Room{
		Name:        name,
		PicUrl:      picUrl,
		OwnerUID:    adminUid,
		Description: description,
		IsPrivate:   isPrivate,
	}

	if err := db.Create(&room).Error; err != nil {
		panic(err)
	}

	return &room, nil
}

func GetRoom(id int) (*Room, error) {
	room := Room{}
	err := db.Where("id = ?", id).First(&room).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrUserProfileNotFound
		}
		panic(err)
	}

	return &room, nil
}

func GetRooms(user *User, ownerUid string, size int, after int) (*[]Room, error) {
	var rooms []Room

	query := db

	if ownerUid != "" {
		query = query.Where("owner_uid = ?", ownerUid)
	} else if !user.IsAdmin {
		return nil, fmt.Errorf("%w: User does not have permissions to fetch all runs", ErrInvalidRequest)
	}

	query = query.Offset(after).Limit(size).Order("created_at desc")

	err := query.Find(&rooms).Error
	if err != nil {
		panic(err)
	}

	return &rooms, nil
}

func (r *Room) Update() (*Room, error) {
	if err := db.Save(r).Error; err != nil {
		panic(err)
	}

	return r, nil
}

func DeleteRoom(roomId int, user *User) error {
	room := Room{}
	err := db.Where("id = ?", roomId).First(&room).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return ErrRoomNotFound
		}
		util.Log.Sugar().Errorf("Unable to get room (room id: %d)", roomId)
		panic(err)
	}

	if !user.IsAdmin && room.OwnerUID != user.ID {
		return ErrRoomDeleteNotAllowed
	}

	if err = db.Unscoped().Delete(&Room{}, "id = ?", roomId).Error; err != nil {
		util.Log.Sugar().Errorf("Unable to delete room (room id: %d)", roomId)
		panic(err)
	}
	return nil
}

var (
	ErrRoomNotFound         = errors.New("Room not found")
	ErrInvalidRequest       = errors.New("Invalid request")
	ErrRoomDeleteNotAllowed = errors.New("A room can be deleted only by the room owner or an admin")
)
