package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"food/config"
	"food/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func CreateReview(c *gin.Context) {
	var input models.Review
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := c.MustGet("db").(*gorm.DB)
	review := models.Review{
		Text:            input.Text,
		UserPublisherID: input.UserPublisherID,
		UserID:          input.UserID,
		ProductID:       input.ProductID,
		ServiceID:       input.ServiceID,
		BusinessID:      input.BusinessID,
	}

	db.Create(&review)
	c.JSON(http.StatusOK, gin.H{"data": review})
}

func GetReview(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid review id"})
		return
	}

	db := c.MustGet("db").(*gorm.DB)
	var review models.Review
	if err := db.Where("id = ?", id).First(&review).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Review not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": review})
}

func UpdateReview(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid review id"})
		return
	}

	db := c.MustGet("db").(*gorm.DB)
	var review models.Review
	if err := db.Where("id = ?", id).First(&review).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Review not found"})
		return
	}

	var input models.Review
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Model(&review).Updates(models.Review{
		Text:            input.Text,
		UserPublisherID: input.UserPublisherID,
		UserID:          input.UserID,
		ProductID:       input.ProductID,
		ServiceID:       input.ServiceID,
		BusinessID:      input.BusinessID,
	})

	c.JSON(http.StatusOK, gin.H{"data": review})
}

func DeleteReview(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid review id"})
		return
	}

	db := c.MustGet("db").(*gorm.DB)
	var review models.Review
	if err := db.Where("id = ?", id).First(&review).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Review not found"})
		return
	}

	db.Delete(&review)

	c.JSON(http.StatusOK, gin.H{"data": fmt.Sprintf("Review with ID %d has been deleted", id)})
}

func ListReviews(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var reviews []models.Review
	if err := db.Find(&reviews).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving reviews"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": reviews})
}
