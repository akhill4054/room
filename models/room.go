package models

import (
	"gorm.io/gorm"
)

type Room struct {
	gorm.Model
	Id           string `gorm:"primaryKey;autoIncrement:true"`
	PicUrl       string
	Name         string `gorm:"not null"`
	Admin        User   `gorm:"foreignKey:CompanyRefer;constraint:OnDelete:CASCADE;not null"`
	MembersCount string `gorm:"not null"`
	IsPrivate    string `gorm:"not null"`
	Description  string `gorm:"not null"`
}
