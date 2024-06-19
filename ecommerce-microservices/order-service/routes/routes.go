package routes

import (
    "ecommerce/order-service/controllers"
    "github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
    r.GET("/orders", controllers.GetOrders)
    r.POST("/orders", controllers.CreateOrder)
    r.GET("/orders/:id", controllers.GetOrder)
}
