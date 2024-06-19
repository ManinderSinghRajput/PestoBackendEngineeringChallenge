package routes

import (
    "ecommerce/product-service/controllers"
    "github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
    r.GET("/products", controllers.GetProducts)
    r.POST("/products", controllers.CreateProduct)
    r.PUT("/products/:id", controllers.UpdateProduct)
    r.DELETE("/products/:id", controllers.DeleteProduct)
}
