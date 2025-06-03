package main

import (
	"auth-golang-react/controllers"
	"auth-golang-react/initializers"
	"auth-golang-react/middleware"

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
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.HandleAuth, controllers.Validate)

	r.Run()
}
