package paramDto

type ParamCreateMessageDto struct {
	UserId            int64  `json:"UserId"`
	MessageContent    string `json:"MessageContent"`
	IsReplyType       bool   `json:"IsReplyType"`
	PreviousMessageId int64  `json:"PreviousMessageId,omitempty"`
}
