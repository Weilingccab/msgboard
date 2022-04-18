package dto

type MessageReplyDto struct {
	MainMessageId     int64 `json:"MainMessageId,omitempty"`
	PreviousMessageId int64 `json:"PreviousMessageId,omitempty"`
}
