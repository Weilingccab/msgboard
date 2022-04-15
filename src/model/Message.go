package model

import (
	"time"

	"gorm.io/gorm"
)

type Message struct {
	MessageId       int64         `gorm:"column:MessageId;type:auto_increment;primary_key;" json:"MessageId,omitempty"`
	UserId          int64         `gorm:"column:UserId;" json:"UserId,omitempty"`
	User            *User         `gorm:"foreignKey:UserId;"`
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
func (Message) CreateMessage(db *gorm.DB, message *Message) (err error) {
	err = db.Create(message).Error
	if err != nil {
		return err
	}
	return nil
}

//get Messages
func (Message) GetMessages(db *gorm.DB, message *[]Message) (err error) {
	err = db.Preload("User").Find(message).Error
	if err != nil {
		return err
	}
	return nil
}

// //get Message by id
// func GetMessage(db *gorm.DB, Message *Message, id int64) (err error) {
// 	err = db.First(Message, id).Error
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// //update Message
// func UpdateMessage(db *gorm.DB, Message *Message) (err error) {
// 	db.Save(Message)
// 	return nil
// }
