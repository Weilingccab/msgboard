package model

type MessageReply struct {
	MessageReplySrno  int64 `gorm:"column:MessageReplySrno;type:auto_increment;primary_key;" json:"MessageReplySrno,omitempty"`
	MessageId         int64 `gorm:"column:MessageId;" json:"MessageId,omitempty"`
	MainMessageId     int64 `gorm:"column:MainMessageId;" json:"MainMessageId,omitempty"`
	PreviousMessageId int64 `gorm:"column:PreviousMessageId;" json:"PreviousMessageId,omitempty"`
}

func (MessageReply) TableName() string {
	return "MessageReply"
}
