package main

import (
	"crud-simple/controller"
	"crud-simple/initializers"
	"crud-simple/middleware"
	"crud-simple/models"
	"github.com/gin-gonic/gin"
	"log"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()
	initializers.Validator()

	initializers.DB.AutoMigrate(&models.User{})

	initializers.DB.AutoMigrate(&models.Todo{})

}

func main() {

	r := gin.Default()

	r.POST("/signup", controller.Register)
	r.POST("/login", controller.Login)

	r.Use(middleware.VerifyToken())
	r.GET("/todo", controller.GetTodo)
	r.POST("/todo", controller.CreateTodo)
	r.PUT("/todo", controller.UpdateTodo)
	r.DELETE("/todo/:id", controller.RemoveTodo)

	// listen and serve in server
	err := r.Run()
	if err != nil {
		log.Panicln("fail to serve")
		return
	}
}
