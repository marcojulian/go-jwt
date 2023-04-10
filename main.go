package main

import (
	"github.com/gin-gonic/gin"
	"github.com/marcojulian/go-jwt/controllers"
	"github.com/marcojulian/go-jwt/initializers"
	"github.com/marcojulian/go-jwt/middleware"
)

func init() {
	initializers.LoanEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()
	r.POST("/singup", controllers.Singup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)
	r.Run()
}
