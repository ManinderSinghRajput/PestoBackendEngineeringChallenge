package routes

import (
    "ecommerce/user-service/controllers"
    "github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
    r.POST("/register", controllers.Register)
    r.POST("/login", controllers.Login)
    r.GET("/profile", controllers.Profile)
}
