package main

import (
	"crud-simple/controller"
	"crud-simple/initializers"
	"crud-simple/middleware"
	"crud-simple/models"

	"github.com/gin-gonic/gin"
)

func init() {
	//load env variable
	initializers.LoadEnv()
	initializers.ConnectToDB()
	initializers.Validator()
	err := initializers.DB.AutoMigrate(&models.User{})
	if err != nil {
		return
	}

}

func main() {

	r := gin.Default()

	r.POST("/signup", controller.Register)
	r.POST("/login", controller.Login)

	r.Use(middleware.VerifyToken())
	r.POST("/home", controller.Test)

	// listen and serve in server
	r.Run()
}
