package main

import (
	"go-auth-backend/config"
	"go-auth-backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	config.ConnectDB()
	//r := gin.Default()
	routes.AuthRoutes(r)

	r.Run(":8080")
}
