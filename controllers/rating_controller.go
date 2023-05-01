package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"foodnetwork/config"
	"foodnetwork/models"
)

func CreateRating(c *gin.Context) {
	db := config.InitDB()

	var rating models.Rating
	if err := c.ShouldBindJSON(&rating); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Create(&rating)

	c.JSON(http.StatusCreated, gin.H{"data": rating})
}

func GetRating(c *gin.Context) {
	db := config.InitDB()

	var rating models.Rating
	if err := db.Where("id = ?", c.Param("id")).First(&rating).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": rating})
}

func UpdateRating(c *gin.Context) {
	db := config.InitDB()

	var rating models.Rating
	if err := db.Where("id = ?", c.Param("id")).First(&rating).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	if err := c.ShouldBindJSON(&rating); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Save(&rating)

	c.JSON(http.StatusOK, gin.H{"data": rating})
}

func DeleteRating(c *gin.Context) {
	db := config.InitDB()

	var rating models.Rating
	if err := db.Where("id = ?", c.Param("id")).First(&rating).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	db.Delete(&rating)

	c.JSON(http.StatusOK, gin.H{"data": "Record deleted successfully"})
}
