package models

import (
	"errors"

	"gorm.io/gorm"
)

type UserProfile struct {
	gorm.Model
	Id     int    `gorm:"primaryKey;autoIncrement:true"`
	Uid    int    `gorm:"unique;index;not null;default: null"`
	Name   string `gorm:"not null;default:null"`
	Bio    string
	PicUrl string
	User   User `gorm:"foreignKey:Uid;references:Id;constraint:OnDelete:SET NULL"`
}

func CreateUserProfile(
	uid int,
	name string,
	bio string,
	picUrl string,
) (*UserProfile, error) {

	var profileCount int64

	db.Where("uid = ?", uid).Model(&UserProfile{}).Count(&profileCount)
	if profileCount != 0 {
		return nil, ErrUserProfileAlreadyExists
	}

	profile := UserProfile{
		Uid:    uid,
		Name:   name,
		Bio:    bio,
		PicUrl: picUrl,
	}

	if err := db.Create(&profile).Error; err != nil {
		panic(err)
	}

	return &profile, nil
}

func UpdateUserProfile(
	user *User,
	profileId int,
	name string,
	bio string,
	picUrl string,
) (*UserProfile, error) {
	profile := UserProfile{}
	err := db.Where("id = ?", profileId).Joins("User").First(&profile).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrUserProfileDoesNotExist
		}
		panic(err)
	}

	if profile.Uid != user.Id && !user.IsAdmin {
		return nil, ErrProfileUpdateNotAllowed
	}

	// Update fields
	profile.Name = name
	profile.Bio = bio
	profile.PicUrl = picUrl

	if err := db.Save(&profile).Error; err != nil {
		panic(err)
	}

	return &profile, nil
}

func GetUserProfile(uid int) (*UserProfile, error) {
	profile := UserProfile{}
	err := db.Where("uid = ?", uid).Joins("User").First(&profile).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrUserProfileNotFound
		}
		panic(err)
	}

	return &profile, nil
}

var (
	ErrUserProfileAlreadyExists = errors.New("User profile already exists")
	ErrUserProfileNotFound      = errors.New("User profile not found")
	ErrUserProfileDoesNotExist  = errors.New("User profile does not exist")
	ErrProfileUpdateNotAllowed  = errors.New("Profile update not allowed")
)
