package main

import (
    "ecommerce/product-service/routes"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    routes.SetupRoutes(r)
    r.Run(":8081") // Run on port 8081
}
