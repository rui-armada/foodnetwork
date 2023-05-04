package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"foodnetwork/config"
	"foodnetwork/models"
)

func CreateBusinessType(c *gin.Context) {
	db := config.InitDB()

	var businessType models.BusinessType
	if err := c.ShouldBindJSON(&businessType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Create(&businessType)

	c.JSON(http.StatusCreated, gin.H{"data": businessType})
}

func GetBusinessType(c *gin.Context) {
	db := config.InitDB()

	var businessType models.BusinessType
	if err := db.Where("id = ?", c.Param("id")).First(&businessType).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": businessType})
}

func UpdateBusinessType(c *gin.Context) {
	db := config.InitDB()

	var businessType models.BusinessType
	if err := db.Where("id = ?", c.Param("id")).First(&businessType).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	if err := c.ShouldBindJSON(&businessType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Save(&businessType)

	c.JSON(http.StatusOK, gin.H{"data": businessType})
}

func DeleteBusinessType(c *gin.Context) {
	db := config.InitDB()

	var businessType models.BusinessType
	if err := db.Where("id = ?", c.Param("id")).First(&businessType).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	db.Delete(&businessType)

	c.JSON(http.StatusOK, gin.H{"data": "Record deleted successfully"})
}
