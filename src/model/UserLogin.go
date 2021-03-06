package model

import (
	"time"

	"gorm.io/gorm"
)

type UserLogin struct {
	UserLoginTokenId string    `gorm:"column:UserLoginTokenId;type:primary_key;" json:"UserLoginTokenId"`
	UserId           int64     `gorm:"column:UserId;" json:"UserId"`
	LoginDateTime    time.Time `gorm:"column:LoginDateTime; autoCreateTime:milli" json:"LoginDateTime"`
}

func (UserLogin) TableName() string {
	return "UserLogin"
}

func NewUserLoginModel() *UserLogin {
	return &UserLogin{}
}

//create a user login
func (*UserLogin) CreateUserLogin(db *gorm.DB, userLogin *UserLogin) (err error) {
	err = db.Create(userLogin).Error
	if err != nil {
		return err
	}
	return nil
}

//check UserLogin by userId
func (*UserLogin) GetUserLogin(db *gorm.DB, userLogin *UserLogin, userId int64) (err error) {
	err = db.Where("\"UserId\" = ?", userId).First(userLogin).Error
	if err != nil {
		return err
	}
	return nil
}


//check UserLogin by userId
func (*UserLogin) GetUserLoginToken(db *gorm.DB, userLogin *UserLogin, userLoginTokenId string) (err error) {
	err = db.Where("\"UserLoginTokenId\" = ?", userLoginTokenId).First(userLogin).Error
	if err != nil {
		return err
	}
	return nil
}
