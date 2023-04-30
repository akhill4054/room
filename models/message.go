package models

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	Uid     string `json:"id" binding:"required"`
	Message string `json:"message" binding:"required"`
}
