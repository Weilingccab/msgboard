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
	userRepo := controller.NewUserRepo()
	r.POST("/user", userRepo.CreateUser)
	r.GET("/users", userRepo.GetUsers)
	r.GET("/user/:UserId", userRepo.GetUser)
	r.PUT("/user/:UserId", userRepo.UpdateUser)
	r.DELETE("/user/:UserId", userRepo.DeleteUser)
	return r
}

func main() {

	r := setupRouter()
	r.Run(":8082") // default localhost:8000

}
