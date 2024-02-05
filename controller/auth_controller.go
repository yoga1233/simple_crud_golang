package controller

import (
	"crud-simple/helper"
	"crud-simple/initializers"
	"crud-simple/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var body struct {
		Email    string
		Password string
	}
	err := c.ShouldBindJSON(&body)
	if err != nil {
		response := helper.ApiResponseFailure(err.Error(), http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		response := helper.ApiResponseFailure(err.Error(), http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	user := models.User{Email: body.Email, Password: string(hash)}

	result := initializers.DB.Create(&user)

	if result.Error != nil {
		response := helper.ApiResponseFailure("Failed to Sign Up", http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiResponseSuccess("Success", http.StatusOK, user)
	c.JSON(http.StatusOK, response)
	return

}

func Login(c *gin.Context) {
	var body struct {
		Email    string
		Password string
	}

	err := c.ShouldBindJSON(&body)
	if err != nil {
		response := helper.ApiResponseFailure(err.Error(), http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var user models.User

	initializers.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		response := helper.ApiResponseFailure("invalid Email or Password", http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	//decrypt password
	err2 := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err2 != nil {
		response := helper.ApiResponseFailure("invalid Email or Password", http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// //generate jwt token

	claims := jwt.MapClaims{
		"authorized": true,
		"user":       user.Email,
		"exp":        time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))
	if err != nil {

		return
	}

	c.JSON(http.StatusOK, helper.ApiResponseSuccess("Success", http.StatusOK, gin.H{"token": tokenString}))
}

func Test(c *gin.Context) {
	c.JSON(http.StatusOK, helper.ApiResponseSuccess("Success", http.StatusOK, gin.H{"message": "ini home page"}))

}
