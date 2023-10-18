package models

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	Id               int `gorm:"primaryKey;autoIncrement:true"`
	SenderUid        int
	Sender           User `gorm:"foreignKey:SenderUid;references:Id;constraint:OnDelete:SET NULL"`
	Text             string
	ReplyToMessageId int
	RepliedToMessage *Message `gorm:"foreignKey:RepliedToMessageId;references:Id;constraint:OnDelete:SET NULL"`
	attatchements    []string `gorm:"type:text[]"`
}
