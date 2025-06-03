package main

import (
	"auth-golang-react/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
}

func main() {
	r := gin.Default()
	r.Run()
}
