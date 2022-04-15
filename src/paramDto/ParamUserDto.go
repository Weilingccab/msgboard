package paramDto

type ParamCreateUserDto struct {
	Account     string `json:"Account"`
	Password    string `json:"Password"`
	IsAuthorize bool   `json:"IsAuthorize"`
}

type ParamUpdateUserDto struct {
	IsAuthorize bool `json:"IsAuthorize"`
}
