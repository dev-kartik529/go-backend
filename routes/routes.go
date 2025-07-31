package routes

import (
	"go-auth-backend/controllers"
	"go-auth-backend/middleware"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine) {
	r.POST("/signup", controllers.RegisterUser)
	r.POST("/login", controllers.LoginUser)

	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware())
	protected.GET("/protected", controllers.Protected)
}
