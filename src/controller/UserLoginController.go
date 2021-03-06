package controller

import (
	"errors"
	"msgboard/db"
	"msgboard/src/model"
	"msgboard/src/paramDto"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserLoginRepo struct {
	Db *gorm.DB
}

func NewUserLoginRepo() *UserLoginRepo {
	db := db.InitDb()
	return &UserLoginRepo{Db: db}
}

// @Summary 使用者登入
// @Tags 使用者登入相關
// @Accept  json
// @Produce  json
// @Param user body paramDto.ParamUserLoginDto true "欲送出的使用者資料"
// @Success 200 object model.UserLogin "使用者登入回傳資料"
// @Failure 400 string string  "{"error": errInfo}"
// @Router /msgboard/userLoginInfo/userLogin [post]
func (repository *UserLoginRepo) CreateUserLogin(c *gin.Context) {
	userLoginModel := model.NewUserLoginModel()
	userModel := model.NewUserModel()
	var paramUserLoginDto paramDto.ParamUserLoginDto
	c.BindJSON(&paramUserLoginDto)
	var user model.User
	err := userModel.GetUser(repository.Db, &user, paramUserLoginDto.UserId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			errmsg := "userId not found:" + strconv.FormatInt(int64(paramUserLoginDto.UserId), 10)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": errmsg})
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	if user.IsAuthorize {
		//check is login, if login return token
		var userLogin model.UserLogin
		err := userLoginModel.GetUserLogin(repository.Db, &userLogin, paramUserLoginDto.UserId)

		if err != nil {
			//	找不到則可登入
			if errors.Is(err, gorm.ErrRecordNotFound) {
				//建立登入資訊
				userLogin.UserLoginTokenId = uuid.New().String()
				userLogin.UserId = user.UserId

				err = userLoginModel.CreateUserLogin(repository.Db, &userLogin)
				if err != nil {
					c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
					return
				}
				c.JSON(http.StatusOK, userLogin)
				return
			}

			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		//找到登入資料直接回傳
		c.JSON(http.StatusOK, userLogin)

	} else {
		errmsg := "userId not authrize:" + strconv.FormatInt(int64(paramUserLoginDto.UserId), 10)
		c.JSON(http.StatusOK, gin.H{"error": errmsg})

	}

}

// @Summary 檢查使用者登入狀態
// @Tags 使用者登入相關
// @Accept  json
// @Produce  json
// @Success 200 object model.UserLogin "使用者登入資料"
// @Failure 400 string string  "{"error": errInfo}"
// @Router /msgboard/userInfo/userLogin/{UserId} [put]
func (repository *UserLoginRepo) CheckUserLogin(c *gin.Context) {
	userLoginModel := model.NewUserLoginModel()
	id, _ := c.Params.Get("UserId")
	var userLogin model.UserLogin
	userId, _ := strconv.ParseInt(id, 10, 64)
	err := userLoginModel.GetUserLogin(repository.Db, &userLogin, userId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			errmsg := "userId not login:" + id
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": errmsg})
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, userLogin)

}
