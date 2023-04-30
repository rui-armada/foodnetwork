package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"foodnetwork/config"
	"foodnetwork/models"
)

func CreateBusinessUser(c *gin.Context) {
	db := config.InitDB()

	var businessUser models.BusinessUser
	if err := c.ShouldBindJSON(&businessUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Create(&businessUser)

	c.JSON(http.StatusCreated, gin.H{"data": businessUser})
}

func GetBusinessUser(c *gin.Context) {
	db := config.InitDB()

	var businessUser models.BusinessUser
	if err := db.Where("id = ?", c.Param("id")).First(&businessUser).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": businessUser})
}

func UpdateBusinessUser(c *gin.Context) {
	db := config.InitDB()

	var businessUser models.BusinessUser
	if err := db.Where("id = ?", c.Param("id")).First(&businessUser).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	if err := c.ShouldBindJSON(&businessUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Save(&businessUser)

	c.JSON(http.StatusOK, gin.H{"data": businessUser})
}

func DeleteBusinessUser(c *gin.Context) {
	db := config.InitDB()

	var businessUser models.BusinessUser
	if err := db.Where("id = ?", c.Param("id")).First(&businessUser).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	db.Delete(&businessUser)

	c.JSON(http.StatusOK, gin.H{"data": "Record deleted successfully"})

}
