package controllers

import (
	"net/http"

	"food/config"
	"food/models"
	"github.com/gin-gonic/gin"
)

// CreateJobPost creates a new job post
func CreateJobPost(c *gin.Context) {
	var input models.JobPost
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	jobPost := models.JobPost{
		Title:       input.Title,
		Description: input.Description,
		BusinessID:  input.BusinessID,
	}

	if err := models.DB.Create(&jobPost).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create job post"})
		return
	}

	c.JSON(http.StatusOK, jobPost)
}

// GetJobPost retrieves a job post by ID
func GetJobPost(c *gin.Context) {
	var jobPost models.JobPost
	if err := models.DB.First(&jobPost, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Job post not found"})
		return
	}

	c.JSON(http.StatusOK, jobPost)
}

// UpdateJobPost updates a job post by ID
func UpdateJobPost(c *gin.Context) {
	var jobPost models.JobPost
	if err := models.DB.First(&jobPost, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Job post not found"})
		return
	}

	var input models.JobPost
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	jobPost.Title = input.Title
	jobPost.Description = input.Description
	jobPost.BusinessID = input.BusinessID

	if err := models.DB.Save(&jobPost).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update job post"})
		return
	}

	c.JSON(http.StatusOK, jobPost)
}

// DeleteJobPost deletes a job post by ID
func DeleteJobPost(c *gin.Context) {
	var jobPost models.JobPost
	if err := models.DB.First(&jobPost, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Job post not found"})
		return
	}

	if err := models.DB.Delete(&jobPost).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete job post"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Job post deleted"})
}

// ListJobPosts retrieves a list of all job posts
func ListJobPosts(c *gin.Context) {
	var jobPosts []models.JobPost
	if err := models.DB.Find(&jobPosts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get job posts"})
		return
	}

	c.JSON(http.StatusOK, jobPosts)
}
