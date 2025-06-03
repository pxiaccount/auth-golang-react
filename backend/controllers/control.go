package controllers

import (
	"auth-golang-react/initializers"
	"auth-golang-react/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var x struct {
		Name     string
		Password string
	}

	if c.Bind(&x) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to read body",
		})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(x.Password), 8)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to hash password",
		})
		return
	}

	user := models.User{Name: x.Name, Password: string(hash)}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to create user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User created!",
	})
}
