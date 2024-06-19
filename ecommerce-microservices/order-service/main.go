package main

import (
    "ecommerce/order-service/routes"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    routes.SetupRoutes(r)
    r.Run(":8082") // Run on port 8082
}
