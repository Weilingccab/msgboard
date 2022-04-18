package paramDto

type ParamCreateMessageDto struct {
	UserLoginTokenId  string  `json:"UserLoginTokenId"`
	UserId            int64   `json:"UserId"`
	MessageContent    *string `json:"MessageContent"`
	IsReplyType       bool    `json:"IsReplyType"`
	PreviousMessageId int64   `json:"PreviousMessageId,omitempty"`
}

type ParamUpdateMessageIsHideDto struct {
	IsHide bool `json:"IsHide"`
}

type ParamUpdateMessageIsLockReplyDto struct {
	IsLockReply bool `json:"IsLockReply"`
}

type ParamQueryMessageDto struct {
	UserId               *int64  `json:"UserId"`
	MessageContent       *string `json:"MessageContent"`
	MessageDateTimeStart *string `json:"MessageDateTimeStart"`
	MessageDateTimeTo    *string `json:"MessageDateTimeTo"`
	IsHide               *bool   `json:"IsHide"`
	IsReplyType          *bool   `json:"IsReplyType"`
	IsLockReply          *bool   `json:"IsLockReply"`
}
