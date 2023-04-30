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

func CreateProfessionalExperience(c *gin.Context) {
	var input models.ProfessionalExperience
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := c.MustGet("db").(*gorm.DB)
	professionalExperience := models.ProfessionalExperience{
		UserID:      input.UserID,
		JobTitleID:  input.JobTitleID,
		BusinessID:  input.BusinessID,
		StartDate:   input.StartDate,
		EndDate:     input.EndDate,
		Description: input.Description,
	}

	db.Create(&professionalExperience)
	c.JSON(http.StatusOK, gin.H{"data": professionalExperience})
}

func GetProfessionalExperience(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid professional experience id"})
		return
	}

	db := c.MustGet("db").(*gorm.DB)
	var professionalExperience models.ProfessionalExperience
	if err := db.Where("id = ?", id).First(&professionalExperience).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Professional experience not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": professionalExperience})
}

func UpdateProfessionalExperience(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid professional experience id"})
		return
	}

	db := c.MustGet("db").(*gorm.DB)
	var professionalExperience models.ProfessionalExperience
	if err := db.Where("id = ?", id).First(&professionalExperience).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Professional experience not found"})
		return
	}

	var input models.ProfessionalExperience
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Model(&professionalExperience).Updates(models.ProfessionalExperience{
		UserID:      input.UserID,
		JobTitleID:  input.JobTitleID,
		BusinessID:  input.BusinessID,
		StartDate:   input.StartDate,
		EndDate:     input.EndDate,
		Description: input.Description,
	})

	c.JSON(http.StatusOK, gin.H{"data": professionalExperience})
}

func DeleteProfessionalExperience(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid professional experience id"})
		return
	}

	db := c.MustGet("db").(*gorm.DB)
	var professionalExperience models.ProfessionalExperience
	if err := db.Where("id = ?", id).First(&professionalExperience).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Professional experience not found"})
		return
	}

	db.Delete(&professionalExperience)

	c.JSON(http.StatusOK, gin.H{"data": fmt.Sprintf("Professional experience with ID %d has been deleted", id)})
}

func ListProfessionalExperiences(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var professionalExperiences []models.ProfessionalExperience
	if err := db.Find(&professionalExperiences).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": professionalExperiences})
}
