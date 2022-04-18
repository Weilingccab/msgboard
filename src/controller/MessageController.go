package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"msgboard/db"
	"msgboard/src/dto"
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

func (repository *MessageRepo) GetMessages(c *gin.Context) {
	messageModel := model.NewMessageModel()

	var messages []model.Message
	err := messageModel.GetMessages(repository.Db, &messages)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	//假設送出無需特殊處理，可直接用json轉Dto
	jsondata, _ := json.Marshal(messages)
	var messageDtos []dto.MessageDto
	json.Unmarshal(jsondata, &messageDtos)
	c.JSON(http.StatusOK, messageDtos)

}

// update message hide
func (repository *MessageRepo) UpdateMessageHide(c *gin.Context) {
	messageModel := model.NewMessageModel()

	id, _ := c.Params.Get("MessageId")
	messageId, _ := strconv.ParseInt(id, 10, 64)

	var message model.Message
	err := messageModel.GetMessage(repository.Db, &message, messageId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			errmsg := "messageId not found" + id
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": errmsg})
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	var paramUpdateMessageIsHideDto paramDto.ParamUpdateMessageIsHideDto
	c.BindJSON(&paramUpdateMessageIsHideDto)
	message.IsHide = paramUpdateMessageIsHideDto.IsHide

	err = messageModel.UpdateMessage(repository.Db, &message)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "message is hide successfully"})
}

// update message lock reply
func (repository *MessageRepo) UpdateMessageLockReply(c *gin.Context) {
	messageModel := model.NewMessageModel()

	id, _ := c.Params.Get("MessageId")
	messageId, _ := strconv.ParseInt(id, 10, 64)

	var message model.Message
	err := messageModel.GetMessage(repository.Db, &message, messageId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			errmsg := "messageId not found" + id
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": errmsg})
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	var paramUpdateMessageIsLockReplyDto paramDto.ParamUpdateMessageIsLockReplyDto
	c.BindJSON(&paramUpdateMessageIsLockReplyDto)
	message.IsLockReply = paramUpdateMessageIsLockReplyDto.IsLockReply

	err = messageModel.UpdateMessage(repository.Db, &message)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "message is lock reply successfully"})
}


func (repository *MessageRepo) GetMessagesFlexibleSearch(c *gin.Context) {
	messageModel := model.NewMessageModel()

	var paramQueryMessageDto paramDto.ParamQueryMessageDto
	c.BindJSON(&paramQueryMessageDto)


	var messages []model.Message
	err := messageModel.GetMessagesFlexibleSearch(repository.Db, &messages, &paramQueryMessageDto)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	//假設送出無需特殊處理，可直接用json轉Dto
	jsondata, _ := json.Marshal(messages)
	var messageDtos []dto.MessageDto
	json.Unmarshal(jsondata, &messageDtos)
	c.JSON(http.StatusOK, messageDtos)

}


