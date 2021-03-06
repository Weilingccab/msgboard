package controller

import (
	"encoding/json"
	"errors"
	"msgboard/db"
	"msgboard/src/dto"
	"msgboard/src/model"
	"msgboard/src/paramDto"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserRepo struct {
	Db *gorm.DB
}

func NewUserRepo() *UserRepo {
	db := db.InitDb()
	return &UserRepo{Db: db}
}

// @Summary 建立使用者
// @Tags 使用者相關
// @Accept  json
// @Produce  json
// @Param user body paramDto.ParamCreateUserDto true "欲送出的使用者資料"
// @Success 200 string string  "{"message": "User created successfully"}"
// @Failure 400 string string  "{"error": errInfo}"
// @Router /msgboard/userInfo/user [post]
func (repository *UserRepo) CreateUser(c *gin.Context) {
	userModel := model.NewUserModel()

	var paramUserDto paramDto.ParamCreateUserDto
	c.BindJSON(&paramUserDto)

	//送進DB前的資料處理
	var user model.User
	user.Account = paramUserDto.Account
	user.Password = paramUserDto.Password
	user.IsAuthorize = paramUserDto.IsAuthorize

	err := userModel.CreateUser(repository.Db, &user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}

// @Summary 取得所有使用者資訊
// @Tags 使用者相關
// @Accept  json
// @Produce  json
// @Success 200 array dto.UserDto "使用者資料陣列"
// @Failure 400 string string  "{"error": errInfo}"
// @Router /msgboard/userInfo/users [get]
func (repository *UserRepo) GetUsers(c *gin.Context) {
	userModel := model.NewUserModel()

	var users []model.User
	err := userModel.GetUsers(repository.Db, &users)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	//假設送出無需特殊處理，可直接用json轉Dto
	jsondata, _ := json.Marshal(users)
	var userDtos []dto.UserDto
	json.Unmarshal(jsondata, &userDtos)
	c.JSON(http.StatusOK, userDtos)
}

// @Summary 取得單一使用者資訊
// @Tags 使用者相關
// @Accept  json
// @Produce  json
// @Param UserId path string true "使用者Id"
// @Success 200 object dto.UserDto  "使用者資料"
// @Failure 400 string string  "{"error": errInfo}"
// @Router /msgboard/userInfo/user/{UserId} [get]
func (repository *UserRepo) GetUser(c *gin.Context) {
	userModel := model.NewUserModel()

	id, _ := c.Params.Get("UserId")
	var user model.User
	userId, _ := strconv.ParseInt(id, 10, 64)
	err := userModel.GetUser(repository.Db, &user, userId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			errmsg := "userId not found" + id
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": errmsg})
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	//送至前台前如有需特殊處理，可用此寫法
	var userDto dto.UserDto
	userDto.UserId = user.UserId
	userDto.Account = user.Account
	userDto.IsAuthorize = user.IsAuthorize
	c.JSON(http.StatusOK, userDto)

}

// @Summary 更新使用者授權
// @Tags 使用者相關
// @Accept  json
// @Produce  json
// @Param UserId path string true "使用者Id"
// @Success 200 string string "{"message": "User updated successfully"}"
// @Failure 400 string string  "{"error": errInfo}"
// @Router /msgboard/userInfo/user/isAuthorize/{UserId} [put]
func (repository *UserRepo) UpdateUserIsAuthorize(c *gin.Context) {
	userModel := model.NewUserModel()

	id, _ := c.Params.Get("UserId")
	userId, _ := strconv.ParseInt(id, 10, 64)

	var user model.User
	err := userModel.GetUser(repository.Db, &user, userId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			errmsg := "userId not found" + id
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": errmsg})
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	var paramUpdateUserDto paramDto.ParamUpdateUserDto
	c.BindJSON(&paramUpdateUserDto)
	user.IsAuthorize = paramUpdateUserDto.IsAuthorize

	err = userModel.UpdateUser(repository.Db, &user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

// @Summary 刪除使用者授權
// @Tags 使用者相關
// @Accept  json
// @Produce  json
// @Param UserId path string true "使用者Id"
// @Success 200 string string "{"message": "User deleted successfully"}"
// @Failure 400 string string  "{"error": errInfo}"
// @Router /msgboard/userInfo/user/{UserId} [delete]
func (repository *UserRepo) DeleteUser(c *gin.Context) {
	userModel := model.NewUserModel()

	var user model.User
	id, _ := c.Params.Get("UserId")
	userId, _ := strconv.ParseInt(id, 10, 64)
	err := userModel.DeleteUser(repository.Db, &user, userId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
