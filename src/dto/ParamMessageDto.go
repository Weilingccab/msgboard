package dto

import (
	"time"
)

type MessageDto struct {
	User              *UserDto  `json:"User"`
	MessageContent    string    `json:"MessageContent"`
	MessageDateTime   time.Time `json:"MessageDateTime"`
	IsHide            bool      `json:"IsHide"`
	IsReplyType       bool      `json:"IsReplyType"`
	IsLockReply       bool      `json:"IsLockReply"`
	PreviousMessageId int64     `json:"PreviousMessageId,omitempty"`
}
