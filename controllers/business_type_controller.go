package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"food/config"
	"food/models"
)

// CreateBusinessType creates a new BusinessType
func CreateBusinessType(c *gin.Context) {
	db := config.InitDB()
	defer db.Close()

	var businessType models.BusinessType
	if err := c.ShouldBindJSON(&businessType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Create(&businessType)

	c.JSON(http.StatusCreated, gin.H{"data": businessType})
}

// GetBusinessType gets a BusinessType by ID
func GetBusinessType(c *gin.Context) {
	db := config.InitDB()
	defer db.Close()

	var businessType models.BusinessType
	if err := db.Where("id = ?", c.Param("id")).First(&businessType).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": businessType})
}

// UpdateBusinessType updates a BusinessType by ID
func UpdateBusinessType(c *gin.Context) {
	db := config.InitDB()
	defer db.Close()

	var businessType models.BusinessType
	if err := db.Where("id = ?", c.Param("id")).First(&businessType).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	if err := c.ShouldBindJSON(&businessType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Save(&businessType)

	c.JSON(http.StatusOK, gin.H{"data": businessType})
}

// DeleteBusinessType deletes a BusinessType by ID
func DeleteBusinessType(c *gin.Context) {
	db := config.InitDB()
	defer db.Close()

	var businessType models.BusinessType
	if err := db.Where("id = ?", c.Param("id")).First(&businessType).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	db.Delete(&businessType)

	c.JSON(http.StatusOK, gin.H{"data": "Record deleted successfully"})
}

func ListBusinessTypes(c *gin.Context) {
	db := config.InitDB()
	defer db.Close()

	var businessTypes []models.BusinessType
	if err := db.Find(&businessTypes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": businessTypes})
}
