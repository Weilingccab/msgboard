package controller

import (
	"errors"
	"fmt"
	"msgboard/db"
	"msgboard/src/model"
	"msgboard/src/paramDto"
	"net/http"
	"strconv"

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
	userLoginModel := model.NewUserLoginModel()

	var paramCreateMessageDto paramDto.ParamCreateMessageDto
	c.BindJSON(&paramCreateMessageDto)

	//檢查是否Token是否存在
	var userLogin model.UserLogin
	err := userLoginModel.GetUserLoginToken(repository.Db, &userLogin, paramCreateMessageDto.UserLoginTokenId)
	if err != nil {
		//	找不到則可登入
		if errors.Is(err, gorm.ErrRecordNotFound) {
			errmsg := "userId not login:" + strconv.FormatInt(int64(paramCreateMessageDto.UserId), 10)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": errmsg})
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	messageModel := model.NewMessageModel()

	//送進DB前的資料處理
	var message model.Message
	message.MessageContent = paramCreateMessageDto.MessageContent
	fmt.Println(message.MessageContent)
	message.UserId = paramCreateMessageDto.UserId
	message.IsReplyType = paramCreateMessageDto.IsReplyType
	message.IsLockReply = false
	message.IsHide = false
	if paramCreateMessageDto.IsReplyType {
		var messageReply model.MessageReply
		messageReplyModel := model.NewMessageReplyModel()

		var dbMessageReply model.MessageReply
		//找回覆紀錄表取得首筆留言Id
		err := messageReplyModel.GetMessageReply(repository.Db, &dbMessageReply, paramCreateMessageDto.PreviousMessageId)
		if err != nil {
			//找不到，表示該為第一筆回覆
			if errors.Is(err, gorm.ErrRecordNotFound) {
				messageReply.MainMessageId = paramCreateMessageDto.PreviousMessageId
			} else {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
				return
			}
		} else {
			messageReply.MainMessageId = dbMessageReply.MainMessageId
		}
		messageReply.PreviousMessageId = paramCreateMessageDto.PreviousMessageId
		message.MessageReply = &messageReply
	}

	err = messageModel.CreateMessage(repository.Db, &message)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "message created successfully"})
}
