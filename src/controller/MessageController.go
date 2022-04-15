package controller

import (
	"msgboard/db"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type MessageRepo struct {
	Db *gorm.DB
}

func NewMessageRepo() *MessageRepo {
	db := db.InitDb()
	return &MessageRepo{Db: db}
}

//create user
func (repository *MessageRepo) CreateMessage(c *gin.Context) {
	// var paramUserDto paramDto.ParamCreateUserDto
	// c.BindJSON(&paramUserDto)

	// //送進DB前的資料處理
	// var user model.User
	// user.Account = paramUserDto.Account
	// user.Password = paramUserDto.Password
	// user.IsAuthorize = paramUserDto.IsAuthorize

	// err := model.CreateUser(repository.Db, &user)
	// if err != nil {
	// 	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
	// 	return
	// }
	// c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}
