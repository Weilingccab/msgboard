package dto

type UserDto struct {
	UserId      int64  `json:"UserId"`
	Account     string `json:"Account"`
	IsAuthorize bool   `json:"IsAuthorize"`
}
