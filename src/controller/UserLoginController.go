package controller

import (
	"msgboard/db"
	"msgboard/src/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserLoginRepo struct {
	Db *gorm.DB
}

func NewUserLoginRepo() *UserLoginRepo {
	db := db.InitDb()
	return &UserLoginRepo{Db: db}
}

//create user
func (repository *UserRepo) CreateUserLogin(c *gin.Context) {
	var user model.User
	c.BindJSON(&user)
	err := model.CreateUser(repository.Db, &user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, user)
}

//get users
// func (repository *UserRepo) GetUsers(c *gin.Context) {
// 	var user []model.User
// 	err := model.GetUsers(repository.Db, &user)
// 	if err != nil {
// 		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
// 		return
// 	}
// 	c.JSON(http.StatusOK, user)
// }
