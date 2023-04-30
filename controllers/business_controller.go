package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"food/config"
	"food/models"
)

func CreateBusiness(c *gin.Context) {
	db := config.InitDB()
	defer func() {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}()

	var business models.Business
	if err := c.ShouldBindJSON(&business); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Create(&business)

	c.JSON(http.StatusCreated, gin.H{"data": business})
}

func GetBusiness(c *gin.Context) {
	db := config.InitDB()
	defer func() {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}()

	var business models.Business
	if err := db.Where("id = ?", c.Param("id")).First(&business).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": business})
}

func UpdateBusiness(c *gin.Context) {
	db := config.InitDB()
	defer func() {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}()

	var business models.Business
	if err := db.Where("id = ?", c.Param("id")).First(&business).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	if err := c.ShouldBindJSON(&business); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Save(&business)

	c.JSON(http.StatusOK, gin.H{"data": business})
}

func DeleteBusiness(c *gin.Context) {
	db := config.InitDB()
	defer func() {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}()

	var business models.Business
	if err := db.Where("id = ?", c.Param("id")).First(&business).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	db.Delete(&business)

	c.JSON(http.StatusOK, gin.H{"data": "Record deleted successfully"})
}
