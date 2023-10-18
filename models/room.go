package models

import (
	"database/sql/driver"
	"errors"
	"fmt"

	"github.com/akhill4054/room-backend/pkg/util"
	"gorm.io/gorm"
)

type Room struct {
	gorm.Model
	Id           int    `gorm:"primaryKey;autoIncrement:true"`
	Name         string `gorm:"not null"`
	PicUrl       string
	OwnerUid     int    `gorm:"index;not null;default: null"`
	Description  string `gorm:"not null"`
	IsPrivate    bool   `gorm:"index;not null"`
	MembersCount string `gorm:"not null"`
	Owner        User   `gorm:"foreignKey:OwnerUid;constraint:OnDelete:SET NULL;not null"`
}

type RoomMemberRole string

const (
	ROOM_ADMIN     RoomMemberRole = "ADMIN"
	ROOM_MODERATOR RoomMemberRole = "MODERATOR"
	ROOM_MEMBER    RoomMemberRole = "MEMBER"
)

func (role *RoomMemberRole) Scan(value interface{}) error {
	*role = RoomMemberRole(value.(string))
	return nil
}

func (role RoomMemberRole) Value() (driver.Value, error) {
	return string(role), nil
}

type RoomMember struct {
	gorm.Model
	RoomId int            `gorm:"primaryKey;autoIncrement:false;not null"`
	Uid    int            `gorm:"primaryKey;autoIncrement:false;not null"`
	Role   RoomMemberRole `sql:"type:ENUM('ADMIN', 'MODERATOR', 'MEMBER')" gorm:"column:role; not null; default MEMBER"`
	Member User           `gorm:"foreignKey:Uid;constraint:OnDelete:CASCADE;not null"`
	Room   Room           `gorm:"foreignKey:RoomId;constraint:OnDelete:CASCADE;not null"`
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
		OwnerUid:    adminUid,
		Description: description,
		IsPrivate:   isPrivate,
	}

	if err := db.Transaction(func(tx *gorm.DB) error {
		if result := tx.Create(&room); result.Error != nil {
			return result.Error
		}

		roomMember := RoomMember{
			RoomId: room.Id,
			Uid:    adminUid,
			Role:   ROOM_ADMIN,
		}
		if result := tx.Create(&roomMember); result.Error != nil {
			return result.Error
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return &room, nil
}

func hasRoomModeratorAccess(roomId int, memberUid int) bool {
	userMemberEntry := RoomMember{}
	if err := db.Where("room_id = ?", roomId).Where("uid = ?", memberUid).First(&userMemberEntry).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false
		}
		panic(err)
	}

	return userMemberEntry.Role == ROOM_ADMIN || userMemberEntry.Role == ROOM_MODERATOR
}

func GetRoom(id int) (*Room, error) {
	room := Room{}
	err := db.Where("id = ?", id).First(&room).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrRoomNotFound
		}
		panic(err)
	}

	return &room, nil
}

func GetRooms(
	user *User,
	ownerUid string,
	isShowOnlyJoined bool,
	size int,
	after int,
) (*[]Room, error) {
	var rooms []Room

	var err error

	if ownerUid != "" {
		err = db.Where("owner_uid = ?", ownerUid).
			Offset(after).Limit(size).Order("created_at desc").
			Find(&rooms).Error
	} else if isShowOnlyJoined {
		t1 := db.Table("room_members").Select("room_id", "uid").Where("uid = ?", user.Id)
		err = db.Model(&Room{}).Joins("RIGHT JOIN (?) t1 ON t1.room_id = rooms.id", t1).
			Offset(after).Limit(size).Order("created_at desc").
			Scan(&rooms).Error
	} else {
		err = db.Offset(after).Limit(size).Order("created_at desc").
			Find(&rooms).Error
	}

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

	if !user.IsAdmin && room.OwnerUid != user.Id {
		return ErrRoomDeleteNotAllowed
	}

	if err = db.Unscoped().Delete(&Room{}, "id = ?", roomId).Error; err != nil {
		util.Log.Sugar().Errorf("Unable to delete room (room id: %d)", roomId)
		panic(err)
	}
	return nil
}

func CreateRoomMember(user *User, memberUid int, roomId int, role RoomMemberRole) error {
	if _, err := GetUser(memberUid); err != nil {
		return err
	}

	room, err := GetRoom(roomId)
	if err != nil {
		return err
	}

	// Check if member already exists
	var memberCount int64
	db.Where("room_id = ?", roomId).Where("uid = ?", memberUid).Model(&RoomMember{}).Count(&memberCount)

	if memberCount != 0 {
		return ErrRoomMemberAlreadyExist
	}

	if !user.IsAdmin &&
		(!room.IsPrivate || user.Id != memberUid) &&
		!hasRoomModeratorAccess(roomId, memberUid) {
		return ErrNotAllowedToCreateMembers
	}

	roomMember := RoomMember{
		RoomId: roomId,
		Uid:    memberUid,
		Role:   role,
	}
	if err := db.Create(&roomMember).Error; err != nil {
		panic(err)
	}

	return nil
}

func GetRoomMember(roomId int, memberId int) (*RoomMember, error) {
	roomMemeber := RoomMember{}
	if err := db.Where("room_id = ?", roomId).Where("uid = ?", memberId).First(&roomMemeber).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrRoomMemeberOrRoomNotFound
		}
		panic(err)
	}

	return &roomMemeber, nil
}

func DeleteRoomMember(user *User, roomId int, memberUid int) error {
	room, err := GetRoom(roomId)
	if err != nil {
		return err
	}

	if user.Id != memberUid &&
		!user.IsAdmin &&
		!hasRoomModeratorAccess(roomId, memberUid) {
		// Check for room moderator access
		return ErrNotAllowedToCreateMembers
	}

	if memberUid == room.OwnerUid {
		return fmt.Errorf("Room admin cannot be removed: %w", ErrInvalidRequest)
	}

	if err = db.Unscoped().Delete(&RoomMember{}, "room_id = ?", roomId, "id = ?", memberUid).Error; err != nil {
		panic(err)
	}

	return nil
}

var (
	ErrMissingPermissions = errors.New("User does not have required permissions to perform this action")
	ErrRoomNotFound       = errors.New("Room not found")
	ErrInvalidRequest     = errors.New("Invalid request")

	ErrRoomMemeberOrRoomNotFound = errors.New("No such room room or room member")

	ErrRoomDeleteNotAllowed      = errors.New("A room can be deleted only by the room owner or an admin")
	ErrNotAllowedToCreateMembers = errors.New("User is not allowed to create room members")
	ErrRoomMemberAlreadyExist    = errors.New("Room member already exists")
	ErrAlreadyPresentInRoom      = errors.New("Already present in the room")
	ErrNotAllowedToJoinRoom      = errors.New("User not allowed to join room")
)
