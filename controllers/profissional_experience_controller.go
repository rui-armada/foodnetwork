package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"foodnetwork/config"
	"foodnetwork/models"
)

func CreateProfessionalExperience(c *gin.Context) {
	db := config.InitDB()

	var experience models.ProfessionalExperience
	if err := c.ShouldBindJSON(&experience); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Create(&experience)

	c.JSON(http.StatusCreated, gin.H{"data": experience})
}

func GetProfessionalExperience(c *gin.Context) {
	db := config.InitDB()

	var experience models.ProfessionalExperience
	if err := db.Where("id = ?", c.Param("id")).First(&experience).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": experience})
}

func UpdateProfessionalExperience(c *gin.Context) {
	db := config.InitDB()

	var experience models.ProfessionalExperience
	if err := db.Where("id = ?", c.Param("id")).First(&experience).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	if err := c.ShouldBindJSON(&experience); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Save(&experience)

	c.JSON(http.StatusOK, gin.H{"data": experience})
}

func DeleteProfessionalExperience(c *gin.Context) {
	db := config.InitDB()

	var experience models.ProfessionalExperience
	if err := db.Where("id = ?", c.Param("id")).First(&experience).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	db.Delete(&experience)

	c.JSON(http.StatusOK, gin.H{"data": "Record deleted successfully"})
}
