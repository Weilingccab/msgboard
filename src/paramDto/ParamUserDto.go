package paramDto

type ParamCreateUserDto struct {
	Account     string `json:"Account,omitempty"`
	Password    string `json:"Password,omitempty"`
	IsAuthorize bool   `json:"IsAuthorize,omitempty"`
}

type ParamUpdateUserDto struct {
	IsAuthorize bool `json:"IsAuthorize,omitempty"`
}
