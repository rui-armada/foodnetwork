package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"foodnetwork/config"
	"foodnetwork/models"
)

func CreateJobTitle(c *gin.Context) {
	db := config.InitDB()

	var jobTitle models.JobTitle
	if err := c.ShouldBindJSON(&jobTitle); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Create(&jobTitle)

	c.JSON(http.StatusCreated, gin.H{"data": jobTitle})
}

func GetJobTitle(c *gin.Context) {
	db := config.InitDB()

	var jobTitle models.JobTitle
	if err := db.Where("id = ?", c.Param("id")).First(&jobTitle).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": jobTitle})
}

func UpdateJobTitle(c *gin.Context) {
	db := config.InitDB()

	var jobTitle models.JobTitle
	if err := db.Where("id = ?", c.Param("id")).First(&jobTitle).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	if err := c.ShouldBindJSON(&jobTitle); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Save(&jobTitle)

	c.JSON(http.StatusOK, gin.H{"data": jobTitle})
}

func DeleteJobTitle(c *gin.Context) {
	db := config.InitDB()
	var jobTitle models.JobTitle
	if err := db.Where("id = ?", c.Param("id")).First(&jobTitle).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	db.Delete(&jobTitle)

	c.JSON(http.StatusOK, gin.H{"data": "Record deleted successfully"})
}

func GetAllJobTitles(c *gin.Context) {
	db := config.InitDB()

	var jobTitles []models.JobTitle
	if err := db.Find(&jobTitles).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": jobTitles})
}
