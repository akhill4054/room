package models

import "gorm.io/gorm"

type Chat struct {
	gorm.Model
	Id            int  `gorm:"primaryKey;autoIncrement:true"`
	SenderUid     int  `gorm:"not null;default: null"`
	SenderUser    User `gorm:"foreignKey:SenderUid;references:Id;constraint:OnDelete:SET NULL"`
	RecipientUid  int  `gorm:"not null;default: null"`
	RecipientUser User `gorm:"foreignKey:RecipientUid;references:Id;constraint:OnDelete:SET NULL"`
}
