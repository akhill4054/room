package models

import "gorm.io/gorm"

type Chat struct {
	gorm.Model
	Id            int `gorm:"primaryKey;autoIncrement:true"`
	SenderUid     int
	SenderUser    User `gorm:"foreignKey:SenderUid;references:Id;constraint:OnDelete:SET NULL"`
	RecipientUid  int
	RecipientUser User `gorm:"foreignKey:RecipientUid;references:Id;constraint:OnDelete:SET NULL"`
}
