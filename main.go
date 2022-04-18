package main

import (
	"msgboard/src/controller"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	//使用者註冊相關
	userRepo := controller.NewUserRepo()
	r.POST("/user", userRepo.CreateUser)           //建立使用者資訊
	r.GET("/users", userRepo.GetUsers)             //取得所有使用者資訊
	r.GET("/user/:UserId", userRepo.GetUser)       //取得單一使用者資訊
	r.DELETE("/user/:UserId", userRepo.DeleteUser) //刪除使用者資訊

	//使用者登入相關
	userLoginRepo := controller.NewUserLoginRepo()
	r.POST("/userLogin", userLoginRepo.CreateUserLogin)       //使用者有授權者可登入
	r.PUT("/userLogin/:UserId", userLoginRepo.CheckUserLogin) //使用者是否登入,已登入者回傳資訊

	//使用者留言相關
	messageRepo := controller.NewMessageRepo()
	r.POST("/message", messageRepo.CreateMessage) //有登入才可建立留言，回覆留言需判斷是否為回覆留言
	r.GET("/messages", messageRepo.GetMessages)   //瀏覽所有留言

	//管理者相關
	r.POST("/messages/flexibleSearch", messageRepo.GetMessagesFlexibleSearch) //管理者查看所有留言並可彈性搜尋
	r.PUT("/user/isAuthorize/:UserId", userRepo.UpdateUserIsAuthorize)           //管理者使用者停權
	r.PUT("/message/isHide/:MessageId", messageRepo.UpdateMessageHide)           //管理者隱藏留言
	r.PUT("/message/isLockReply/:MessageId", messageRepo.UpdateMessageLockReply) //管理者鎖定留言不可回復

	return r
}

func main() {

	r := setupRouter()
	r.Run(":8082") // default localhost:8000

}
