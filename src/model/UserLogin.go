package model

import (
	"time"

	"gorm.io/gorm"
)

type UserLogin struct {
	UserLoginTokenId string    `gorm:"column:UserLoginTokenId;type:primary_key;" json:"UserLoginTokenId,omitempty"`
	UserId           int64     `gorm:"column:UserId;" json:"UserId,omitempty"`
	LoginDateTime    time.Time `gorm:"column:LoginDateTime; autoCreateTime:milli" json:"LoginDateTime,omitempty"`
}

func (UserLogin) TableName() string {
	return "UserLogin"
}

func NewUserLoginModel() *UserLogin {
	return &UserLogin{}
}

//create a user login
func (UserLogin) CreateUserLogin(db *gorm.DB, UserLogin *UserLogin) (err error) {
	err = db.Create(UserLogin).Error
	if err != nil {
		return err
	}
	return nil
}

//check UserLogin by userId
func (UserLogin) GetUserLogin(db *gorm.DB, userLogin *UserLogin, userId int64) (err error) {
	// err = db.First(userLogin, userId).Error
	err = db.Where("\"UserId\" = ?", userId).First(&userLogin).Error
	if err != nil {
		return err
	}
	return nil
}
