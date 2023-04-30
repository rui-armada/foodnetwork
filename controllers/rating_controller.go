// controllers/ratings.go
package controllers

import (
	"net/http"
	"strconv"

	"foodnetwork/config"
	"foodnetwork/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// CreateRating creates a new rating
func CreateRating(c *gin.Context) {
	var rating models.Rating
	if err := c.ShouldBindJSON(&rating); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := models.DB.Create(&rating).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, rating)
}

// GetRating retrieves a rating by ID
func GetRating(c *gin.Context) {
	rating, err := findRatingByID(c)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, rating)
}

// UpdateRating updates an existing rating
func UpdateRating(c *gin.Context) {
	rating, err := findRatingByID(c)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	if err := c.ShouldBindJSON(&rating); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := models.DB.Save(&rating).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, rating)
}

// DeleteRating deletes a rating by ID
func DeleteRating(c *gin.Context) {
	rating, err := findRatingByID(c)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	if err := models.DB.Delete(&rating).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Rating deleted successfully"})
}

// ListRatings lists all ratings
func ListRatings(c *gin.Context) {
	var ratings []models.Rating
	if err := models.DB.Find(&ratings).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, ratings)
}

func findRatingByID(c *gin.Context) (*models.Rating, error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return nil, err
	}
	var rating models.Rating
	if err := models.DB.First(&rating, id).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, err
		}
		return nil, utils.NewInternalServerError(err.Error())
	}
	return &rating, nil
}
