package controllers

import (
	"auth-golang-react/initializers"
	"auth-golang-react/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
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

func Login(c *gin.Context) {
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

	var user models.User
	initializers.DB.First(&user, "Name = ?", x.Name)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid email or password",
		})

		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(x.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid email or password",
		})

		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create token",
		})
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successfully!",
	})
}

func Validate(c *gin.Context) {
	user, _ := c.Get("user")

	c.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}
