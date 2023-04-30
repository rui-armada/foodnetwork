package controllers

import (
	"net/http"

	"foodnetwork/config"
	"foodnetwork/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type ProductController struct {
	DB *gorm.DB
}

func (pc *ProductController) GetAllProducts(c *gin.Context) {
	var products []models.Product
	pc.DB.Find(&products)
	c.JSON(http.StatusOK, products)
}

func (pc *ProductController) GetProductByID(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	if err := pc.DB.Where("id = ?", id).First(&product).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, product)
}

func (pc *ProductController) CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	pc.DB.Create(&product)
	c.JSON(http.StatusCreated, product)
}

func (pc *ProductController) UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	if err := pc.DB.Where("id = ?", id).First(&product).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	pc.DB.Save(&product)
	c.JSON(http.StatusOK, product)
}

func (pc *ProductController) DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	if err := pc.DB.Where("id = ?", id).First(&product).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	pc.DB.Delete(&product)
	c.Status(http.StatusNoContent)
}
