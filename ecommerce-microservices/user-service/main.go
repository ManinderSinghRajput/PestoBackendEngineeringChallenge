package main

import (
    "ecommerce/user-service/routes"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    routes.SetupRoutes(r)
    r.Run(":8080") // Run on port 8080
}
