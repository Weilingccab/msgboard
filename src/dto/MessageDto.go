package dto

import (
	"time"
)

type MessageDto struct {
	MessageId            int64       `json:"messageId"`
	User                 *UserDto    `json:"User"`
	MessageContent       string      `json:"MessageContent"`
	MessageDateTime      time.Time   `json:"MessageDateTime"`
	IsHide               bool        `json:"IsHide"`
	IsReplyType          bool        `json:"IsReplyType"`
	IsLockReply          bool        `json:"IsLockReply"`
	PreviousMessageReply *MessageDto `json:"PreviousMessageReply,omitempty"`
}
