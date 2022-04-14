package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	//gorm為model的tag標籤，v2版的auto_increment要放在type裡面，v1版是放獨立定義
	UserId         int64     `gorm:"column:UserId;type:auto_increment;primary_key;" json:"UserId,omitempty"`
	Account        string    `gorm:"column:Account;" json:"Account,omitempty"`
	Password       string    `gorm:"column:Password;" json:"Password,omitempty"`
	IsAuthorize    bool      `gorm:"column:IsAuthorize; type:boolean;" json:"IsAuthorize,omitempty"`
	CreateDateTime time.Time `gorm:"column:CreateDateTime; autoCreateTime:milli" json:"CreateDateTime,omitempty"`
	UpdateDateTime time.Time `gorm:"column:UpdateDateTime; autoUpdateTime:milli" json:"UpdateDateTime,omitempty"`
}

type Tabler interface {
	TableName() string
}

// TableName 会将 User 的表名重写为 `profiles`
func (User) TableName() string {
	return "User"
}

//create a user
func CreateUser(db *gorm.DB, User *User) (err error) {
	err = db.Create(User).Error
	if err != nil {
		return err
	}
	return nil
}

//get users
func GetUsers(db *gorm.DB, User *[]User) (err error) {
	err = db.Find(User).Error
	if err != nil {
		return err
	}
	return nil
}

//get user by id
func GetUser(db *gorm.DB, User *User, id int64) (err error) {
	err = db.First(User, id).Error
	if err != nil {
		return err
	}
	return nil
}

//update user
func UpdateUser(db *gorm.DB, User *User) (err error) {
	db.Save(User)
	return nil
}

//delete user
func DeleteUser(db *gorm.DB, User *User, id int64) (err error) {
	db.First(User, id).Delete(User)

	return nil
}
