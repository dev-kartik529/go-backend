package main

import (
    "github.com/gin-gonic/gin"
    "go-auth-backend/config"
    "go-auth-backend/routes"
)

func main() {
    config.ConnectDB()
    r := gin.Default()
    routes.AuthRoutes(r)
    r.Run(":8080")
}