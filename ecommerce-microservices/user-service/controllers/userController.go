package controllers

import (
    "ecommerce/user-service/models"
    "ecommerce/user-service/utils"
    "github.com/gin-gonic/gin"
    "net/http"
)

func Register(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    // Hash password and save user to DB (mocked here)
    user.PasswordHash = utils.HashPassword(user.Password)
    // Assume user is saved successfully
    c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func Login(c *gin.Context) {
    var loginData models.Login
    if err := c.ShouldBindJSON(&loginData); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    // Mock user retrieval and password check
    if loginData.Username == "test" && loginData.Password == "password" {
        token, _ := utils.GenerateToken(loginData.Username)
        c.JSON(http.StatusOK, gin.H{"token": token})
        return
    }
    c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials"})
}

func Profile(c *gin.Context) {
    // Mocked profile response
    c.JSON(http.StatusOK, gin.H{"username": "test", "email": "test@example.com"})
}
