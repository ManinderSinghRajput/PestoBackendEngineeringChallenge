package controllers

import (
    "ecommerce/order-service/models"
    "github.com/gin-gonic/gin"
    "net/http"
    "strconv"
)

var orders = []models.Order{
    {ID: 1, UserID: 1, ProductID: 1, Quantity: 2, TotalPrice: 20.0, Status: "Pending"},
}

func GetOrders(c *gin.Context) {
    c.JSON(http.StatusOK, orders)
}

func CreateOrder(c *gin.Context) {
    var order models.Order
    if err := c.ShouldBindJSON(&order); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    order.ID = uint(len(orders) + 1)
    orders = append(orders, order)
    c.JSON(http.StatusCreated, order)
}

func GetOrder(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    for _, order := range orders {
        if order.ID == uint(id) {
            c.JSON(http.StatusOK, order)
            return
        }
    }
    c.JSON(http.StatusNotFound, gin.H{"error": "order not found"})
}
