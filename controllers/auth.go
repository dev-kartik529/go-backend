package controllers

import (
    "context"
    "net/http"
    "os"
    "time"

    "go-auth-backend/config"
    "go-auth-backend/models"
    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v4"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "golang.org/x/crypto/bcrypt"
)

var userCol *mongo.Collection = config.DB.Collection("users")

func Signup(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
    user.Password = string(hashedPassword)

    _, err := userCol.InsertOne(context.TODO(), user)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "User created"})
}

func Login(c *gin.Context) {
    var input models.User
    var dbUser models.User

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    err := userCol.FindOne(context.TODO(), bson.M{"email": input.Email}).Decode(&dbUser)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
        return
    }

    err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(input.Password))
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
        return
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "email": dbUser.Email,
        "exp":   time.Now().Add(time.Hour * 24).Unix(),
    })

    tokenString, _ := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
    c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func Protected(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"message": "This is protected"})
}