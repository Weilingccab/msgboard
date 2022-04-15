package paramDto

type ParamCreateMessageDto struct {
	UserLoginTokenId  string  `json:"UserLoginTokenId"`
	UserId            int64   `json:"UserId"`
	MessageContent    *string `json:"MessageContent"`
	IsReplyType       bool    `json:"IsReplyType"`
	PreviousMessageId int64   `json:"PreviousMessageId,omitempty"`
}
