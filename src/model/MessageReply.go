package model

import "gorm.io/gorm"

type MessageReply struct {
	MessageReplySrno  int64 `gorm:"column:MessageReplySrno;type:auto_increment;primary_key;" json:"MessageReplySrno"`
	MessageId         int64 `gorm:"column:MessageId;" json:"MessageId"`
	MainMessageId     int64 `gorm:"column:MainMessageId;" json:"MainMessageId"`
	PreviousMessageId int64 `gorm:"column:PreviousMessageId;" json:"PreviousMessageId"`
}

func (MessageReply) TableName() string {
	return "MessageReply"
}

func NewMessageReplyModel() *MessageReply {
	return &MessageReply{}
}

func (*MessageReply) GetMessageReply(db *gorm.DB, messageReply *MessageReply, messageId int64) (err error) {
	err = db.Where("\"MessageId\" = ?", messageId).First(messageReply).Error
	if err != nil {
		return err
	}
	return nil
}
