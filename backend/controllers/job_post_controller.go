package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"foodnetwork/config"
	"foodnetwork/models"
)

func CreateJobPost(c *gin.Context) {
	db := config.InitDB()

	var jobPost models.JobPost
	if err := c.ShouldBindJSON(&jobPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Create(&jobPost)

	c.JSON(http.StatusCreated, gin.H{"data": jobPost})
}

func GetJobPost(c *gin.Context) {
	db := config.InitDB()

	var jobPost models.JobPost
	if err := db.Where("id = ?", c.Param("id")).First(&jobPost).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": jobPost})
}

func UpdateJobPost(c *gin.Context) {
	db := config.InitDB()

	var jobPost models.JobPost
	if err := db.Where("id = ?", c.Param("id")).First(&jobPost).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	if err := c.ShouldBindJSON(&jobPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Save(&jobPost)

	c.JSON(http.StatusOK, gin.H{"data": jobPost})
}

func DeleteJobPost(c *gin.Context) {
	db := config.InitDB()

	var jobPost models.JobPost
	if err := db.Where("id = ?", c.Param("id")).First(&jobPost).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	db.Delete(&jobPost)

	c.JSON(http.StatusOK, gin.H{"data": "Record deleted successfully"})
}

func GetJobPosts(c *gin.Context) {
	db := config.InitDB()

	var jobPosts []models.JobPost
	if err := db.Find(&jobPosts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": jobPosts})
}
