package main

import (
	"auth-golang-react/controllers"
	"auth-golang-react/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
	initializers.AutoMigrate()
}

func main() {
	r := gin.Default()

	r.POST("/register", controllers.Register)

	r.Run()
}
