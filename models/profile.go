package models

import "gorm.io/gorm"

type UserProfile struct {
	gorm.Model
	ID     int    `gorm:"primaryKey;"`
	UID    int    `gorm:"not null;unique;index;default: null"`
	Name   string `gorm:"not null;default:null;"`
	Bio    string
	PicUrl string
	User   User `gorm:"foreignKey:UID;constraint:OnDelete:CASCADE;not null;"`
}
