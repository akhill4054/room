package models

import "gorm.io/gorm"

type Chat struct {
	gorm.Model
	UID1 string `json:"uid_1" binding:"required"`
	UID2 string `json:"uid_2" binding:"required"`
}
