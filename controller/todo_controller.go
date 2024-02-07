package controller

import (
	"crud-simple/helper"
	"crud-simple/initializers"
	"crud-simple/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetTodo(c *gin.Context) {
	var todo []models.Todo

	// find in database
	result := initializers.DB.Find(&todo)
	if result.Error != nil {
		response := helper.ApiResponseFailure("failed", http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiResponseSuccess("success", http.StatusOK, todo)
	c.JSON(http.StatusOK, response)
	return
}

func GetTodoByID(c *gin.Context) {
	id := c.Param("id")

	var todo models.Todo
	//
	result := initializers.DB.First(&todo, id)

	if result != nil {
		response := helper.ApiResponseFailure("can't find todo", http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, response)
		return
	}

}

func CreateTodo(c *gin.Context) {
	var body struct {
		Title    string `json:"title,"`
		Desc     string `json:"desc,"`
		Deadline string `json:"deadline,"`
		Status   string `json:"status,"`
		UserID   uint   `json:"userID,"`
	}
	// read parameter
	if err := c.ShouldBindJSON(&body); err != nil {
		response := helper.ApiResponseFailure("can't read parameter", http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	todo := models.Todo{
		Title:    body.Title,
		Desc:     body.Desc,
		Deadline: body.Deadline,
		Status:   body.Status,
		UserID:   body.UserID,
	}
	// save to database
	result := initializers.DB.Create(&todo)
	if result.Error != nil {
		response := helper.ApiResponseFailure("Failed to Create Todo", http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponseSuccess("Success", http.StatusCreated, todo)
	c.JSON(http.StatusCreated, response)

}

func UpdateTodo(c *gin.Context) {
	var body struct {
		TodoID   uint   `json:"todoID"`
		Title    string `json:"title"`
		Desc     string `json:"desc"`
		Deadline string `json:"deadline"`
		Status   string `json:"status"`
		UserID   uint   `json:"userID"`
	}
	// read parameter
	if err := c.ShouldBindJSON(&body); err != nil {
		response := helper.ApiResponseFailure("can't read parameter", http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	var todo models.Todo
	// search todo by id
	result := initializers.DB.First(&todo, body.TodoID)
	if result.Error != nil {
		response := helper.ApiResponseFailure("can't find Todo", http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	todo.Title = body.Title
	todo.Desc = body.Desc
	todo.Deadline = body.Deadline
	todo.Status = body.Status
	todo.UserID = body.UserID

	// update in database
	r := initializers.DB.Save(&todo)
	if r.Error != nil {
		response := helper.ApiResponseFailure("failed to update Todo", http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiResponseSuccess("Success", http.StatusOK, todo)
	c.JSON(http.StatusOK, response)

}

func RemoveTodo(c *gin.Context) {
	id := c.Param("id")

	var todo models.Todo
	result := initializers.DB.First(&todo, id)
	if result.Error != nil {
		response := helper.ApiResponseFailure("Todo not found", http.StatusNotFound)
		c.JSON(http.StatusNotFound, response)
		return
	}

	r := initializers.DB.Unscoped().Delete(&todo)
	if r.Error != nil {
		response := helper.ApiResponseFailure("failed to delete Todo", http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiResponseSuccess("Success", http.StatusOK, gin.H{})
	c.JSON(http.StatusOK, response)
}
