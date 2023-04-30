package controllers

import (
	"net/http"
	"strconv"

	"food/config"
	"food/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CreateJobTitle creates a new job title
func CreateJobTitle(c *gin.Context) {
	var jobTitle models.JobTitle

	if err := c.ShouldBindJSON(&jobTitle); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.DB.Create(&jobTitle).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create job title"})
		return
	}

	c.JSON(http.StatusCreated, jobTitle)
}

// GetJobTitle retrieves a single job title
func GetJobTitle(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid job title ID"})
		return
	}

	var jobTitle models.JobTitle
	if err := models.DB.First(&jobTitle, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Job title not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get job title"})
		return
	}

	c.JSON(http.StatusOK, jobTitle)
}

// UpdateJobTitle updates an existing job title
func UpdateJobTitle(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid job title ID"})
		return
	}

	var jobTitle models.JobTitle
	if err := models.DB.First(&jobTitle, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Job title not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update job title"})
		return
	}

	if err := c.ShouldBindJSON(&jobTitle); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.DB.Save(&jobTitle).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update job title"})
		return
	}

	c.JSON(http.StatusOK, jobTitle)
}

// DeleteJobTitle deletes an existing job title
func DeleteJobTitle(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid job title ID"})
		return
	}

	var jobTitle models.JobTitle
	if err := models.DB.First(&jobTitle, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Job title not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete job title"})
		return
	}

	if err := models.DB.Delete(&jobTitle).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete job title"})
		return
		c.JSON(http.StatusOK, gin.H{"message": "Job title deleted"})
	}
}

// ListJobTitles retrieves a list of all job titles
func ListJobTitles(c *gin.Context) {
	var jobTitles []models.JobTitle
	if err := models.DB.Find(&jobTitles).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get job titles"})
		return
	}

	c.JSON(http.StatusOK, jobTitles)
}
