package main

import (
	"fmt"
	"msgboard/src/controller"

	_ "msgboard/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func setupRouter(r *gin.Engine) *gin.Engine {

	testRouter := r.Group("/msgboard/test")
	{
		testRouter.GET("/ping", func(c *gin.Context) {
			c.String(200, "pong")
		})

	}

	//使用者註冊相關
	userRouter := r.Group("/msgboard/userInfo")
	{
		userRepo := controller.NewUserRepo()
		userRouter.POST("/user", userRepo.CreateUser)                               //建立使用者資訊
		userRouter.GET("/users", userRepo.GetUsers)                                 //取得所有使用者資訊
		userRouter.GET("/user/:UserId", userRepo.GetUser)                           //取得單一使用者資訊
		userRouter.DELETE("/user/:UserId", userRepo.DeleteUser)                     //刪除使用者資訊
		userRouter.PUT("/user/isAuthorize/:UserId", userRepo.UpdateUserIsAuthorize) //管理者使用者停權

	}

	//使用者登入相關
	userLoginRouter := r.Group("/msgboard/userLoginInfo")
	{
		userLoginRepo := controller.NewUserLoginRepo()
		userLoginRouter.POST("/userLogin", userLoginRepo.CreateUserLogin)       //使用者有授權者可登入
		userLoginRouter.PUT("/userLogin/:UserId", userLoginRepo.CheckUserLogin) //使用者是否登入,已登入者回傳資訊

	}

	//使用者留言相關
	messageRouter := r.Group("/msgboard/messageInfo")
	{
		messageRepo := controller.NewMessageRepo()
		messageRouter.POST("/message", messageRepo.CreateMessage)                                //有登入才可建立留言，回覆留言需判斷是否為回覆留言
		messageRouter.GET("/messages", messageRepo.GetMessages)                                  //瀏覽所有留言
		messageRouter.POST("/messages/flexibleSearch", messageRepo.GetMessagesFlexibleSearch)    //管理者查看所有留言並可彈性搜尋
		messageRouter.PUT("/message/isHide/:MessageId", messageRepo.UpdateMessageHide)           //管理者隱藏留言
		messageRouter.PUT("/message/isLockReply/:MessageId", messageRepo.UpdateMessageLockReply) //管理者鎖定留言不可回復

	}

	return r
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// @title Msgboard Demo
// @version 1.0
// @description Swagger API.
// @host 35.78.97.207:8080
func main() {
	r := gin.Default()
	r.Use(CORSMiddleware())

	setupRouter(r)
	ginPort := fmt.Sprintf(":%s", "8080")

	pgConnString := fmt.Sprintf("http://35.78.97.207%s/swagger/doc.json",
		ginPort)
	fmt.Println(pgConnString)
	url := ginSwagger.URL(pgConnString) // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	fmt.Println(ginPort)
	r.Run(ginPort)

}
