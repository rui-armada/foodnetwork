package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"foodnetwork/config"
	"foodnetwork/models"
)

func CreateService(c *gin.Context) {
	db := config.InitDB()

	var service models.Service
	if err := c.ShouldBindJSON(&service); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Create(&service)

	c.JSON(http.StatusCreated, gin.H{"data": service})
}

func GetService(c *gin.Context) {
	db := config.InitDB()

	var service models.Service
	if err := db.Where("id = ?", c.Param("id")).First(&service).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": service})
}

func UpdateService(c *gin.Context) {
	db := config.InitDB()

	var service models.Service
	if err := db.Where("id = ?", c.Param("id")).First(&service).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	if err := c.ShouldBindJSON(&service); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Save(&service)

	c.JSON(http.StatusOK, gin.H{"data": service})
}

func DeleteService(c *gin.Context) {
	db := config.InitDB()

	var service models.Service
	if err := db.Where("id = ?", c.Param("id")).First(&service).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	db.Delete(&service)

	c.JSON(http.StatusOK, gin.H{"data": "Record deleted successfully"})
}
