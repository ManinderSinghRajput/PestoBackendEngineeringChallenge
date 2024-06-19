package controllers

import (
    "ecommerce/product-service/models"
    "github.com/gin-gonic/gin"
    "net/http"
    "strconv"
)

var products = []models.Product{
    {ID: 1, Name: "Product1", Description: "Description1", Price: 10.0, Stock: 100, Version: 1},
}

func GetProducts(c *gin.Context) {
    c.JSON(http.StatusOK, products)
}

func CreateProduct(c *gin.Context) {
    var product models.Product
    if err := c.ShouldBindJSON(&product); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    product.ID = uint(len(products) + 1)
    product.Version = 1
    products = append(products, product)
    c.JSON(http.StatusCreated, product)
}

func UpdateProduct(c *gin.Context) {
    var product models.Product
    if err := c.ShouldBindJSON(&product); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    id, _ := strconv.Atoi(c.Param("id"))
    for i, p := range products {
        if p.ID == uint(id) {
            if p.Version != product.Version {
                c.JSON(http.StatusConflict, gin.H{"error": "version mismatch"})
                return
            }
            product.Version++
            products[i] = product
            c.JSON(http.StatusOK, product)
            return
        }
    }
    c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
}

func DeleteProduct(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    for i, p := range products {
        if p.ID == uint(id) {
            products = append(products[:i], products[i+1:]...)
            c.JSON(http.StatusOK, gin.H{"message": "product deleted"})
            return
        }
    }
    c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
}
