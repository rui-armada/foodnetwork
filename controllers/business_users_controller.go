package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"food/config"
	"food/models"
)

func CreateBusinessUser(c *gin.Context) {
	var input models.BusinessUser
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := database.GetDB()
	if err := db.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create business user"})
		return
	}

	c.JSON(http.StatusCreated, input)
}

func GetBusinessUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var businessUser models.BusinessUser
	db := database.GetDB()
	if err := db.First(&businessUser, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Business user not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch business user"})
		return
	}

	c.JSON(http.StatusOK, businessUser)
}

func UpdateBusinessUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var input models.BusinessUser
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := database.GetDB()
	var businessUser models.BusinessUser
	if err := db.First(&businessUser, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Business user not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch business user"})
		return
	}

	businessUser.UserID = input.UserID
	businessUser.BusinessID = input.BusinessID
	businessUser.JobTitleID = input.JobTitleID

	if err := db.Save(&businessUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update business user"})
		return
	}

	c.JSON(http.StatusOK, businessUser)
}

func DeleteBusinessUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	db := database.GetDB()
	if err := db.Delete(&models.BusinessUser{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete business user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Business user deleted"})
}
