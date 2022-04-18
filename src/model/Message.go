package model

import (
	"fmt"
	"msgboard/src/paramDto"
	"time"

	"gorm.io/gorm"
)

type Message struct {
	MessageId       int64 `gorm:"column:MessageId;type:auto_increment;primary_key;" json:"MessageId,omitempty"`
	UserId          int64 `gorm:"column:UserId;" json:"UserId,omitempty"`
	User            *User
	MessageContent  *string       `gorm:"column:MessageContent;" json:"MessageContent,omitempty"`
	MessageDateTime time.Time     `gorm:"column:MessageDateTime; autoCreateTime:milli" json:"MessageDateTime,omitempty"`
	IsReplyType     bool          `gorm:"column:IsReplyType;" json:"IsReplyType,omitempty"`
	IsHide          bool          `gorm:"column:IsHide;" json:"IsHide,omitempty"`
	IsLockReply     bool          `gorm:"column:IsLockReply;" json:"IsLockReply,omitempty"`
	UpdateDateTime  time.Time     `gorm:"column:UpdateDateTime; autoUpdateTime:milli" json:"UpdateDateTime,omitempty"`
	MessageReply    *MessageReply `gorm:"foreignKey:MessageId" json:"MessageReply,omitempty"`
}

func (Message) TableName() string {
	return "Message"
}

func NewMessageModel() *Message {
	return &Message{}
}

//create a Message
func (*Message) CreateMessage(db *gorm.DB, message *Message) (err error) {
	err = db.Create(message).Error
	if err != nil {
		return err
	}
	return nil
}

//get Messages
func (*Message) GetMessages(db *gorm.DB, message *[]Message) (err error) {
	err = db.Preload("User").Preload("MessageReply").Find(message).Error
	if err != nil {
		return err
	}
	return nil
}

//get message by id
func (*Message) GetMessage(db *gorm.DB, message *Message, id int64) (err error) {
	err = db.First(message, id).Error
	if err != nil {
		return err
	}
	return nil
}

//update message
func (*Message) UpdateMessage(db *gorm.DB, message *Message) (err error) {
	db.Save(message)
	return nil
}

//get message by id
func (*Message) GetMessagesFlexibleSearch(db *gorm.DB, messages *[]Message, queryParamMessage *paramDto.ParamQueryMessageDto) (err error) {
	fmt.Println(queryParamMessage)
	if queryParamMessage.UserId != nil {
		db = db.Preload("User").Where("\"UserId\" = ?", queryParamMessage.UserId)
	}
	if queryParamMessage.MessageContent != nil {
		db = db.Where("\"MessageContent\" = ?", queryParamMessage.MessageContent)
	}
	if queryParamMessage.MessageDateTimeStart != nil && queryParamMessage.MessageDateTimeTo != nil {
		db = db.Where("\"MessageDateTime\" between ? and ?", queryParamMessage.MessageDateTimeStart, queryParamMessage.MessageDateTimeTo)
	}
	if queryParamMessage.IsReplyType != nil {
		db = db.Where("\"IsReplyType\" = ?", queryParamMessage.IsReplyType)
	}
	if queryParamMessage.IsHide != nil {
		db = db.Where("\"IsHide\" = ?", queryParamMessage.IsHide)
	}
	if queryParamMessage.IsLockReply != nil {
		db = db.Where("\"IsLockReply\" = ?", queryParamMessage.IsLockReply)
	}
	err = db.Preload("User").Preload("MessageReply").Find(messages).Error
	if err != nil {
		return err
	}
	return nil
}

// db.Where(&User{Name: "jinzhu", Age: 20}).First(&user)
// SELECT * FROM users WHERE name = "jinzhu" AND age = 20 LIMIT 1;
