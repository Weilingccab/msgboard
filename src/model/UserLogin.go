package model

import (
	"time"

	"gorm.io/gorm"
)

type UserLogin struct {
	//gorm為model的tag標籤，v2版的auto_increment要放在type裡面，v1版是放獨立定義
	UserLoginTokenId string    `gorm:"column:UserLoginTokenId;type:primary_key;" json:"UserLoginTokenId,omitempty"`
	UserId           int64     `gorm:"column:UserId;" json:"UserId,omitempty"`
	LoginDateTime    time.Time `gorm:"column:LoginDateTime; autoCreateTime:milli" json:"LoginDateTime,omitempty"`
}

func (UserLogin) TableName() string {
	return "UserLogin"
}

//create a user login
func CreateUserLogin(db *gorm.DB, UserLogin *UserLogin) (err error) {
	err = db.Create(UserLogin).Error
	if err != nil {
		return err
	}
	return nil
}

//check UserLogin by userId
func CheckUserLogin(db *gorm.DB, userLogin *UserLogin, userId int64) (err error) {
	err = db.First(userLogin, userId).Error
	if err != nil {
		return err
	}
	return nil
}

//delete UserLogin by UserId
func DeleteUserLogin(db *gorm.DB, userLogin *UserLogin, userId int64) (err error) {
	db.First(userLogin, userId).Delete(userLogin)

	return nil
}
