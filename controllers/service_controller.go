package controllers

import (
	"net/http"

	"food/config"
	"food/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type ServiceController struct {
	DB *gorm.DB
}

func (sc *ServiceController) GetAllServices(c *gin.Context) {
	var services []models.Service
	sc.DB.Find(&services)
	c.JSON(http.StatusOK, services)
}

func (sc *ServiceController) GetServiceByID(c *gin.Context) {
	id := c.Param("id")
	var service models.Service
	if err := sc.DB.Where("id = ?", id).First(&service).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, service)
}

func (sc *ServiceController) CreateService(c *gin.Context) {
	var service models.Service
	if err := c.ShouldBindJSON(&service); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	sc.DB.Create(&service)
	c.JSON(http.StatusCreated, service)
}

func (sc *ServiceController) UpdateService(c *gin.Context) {
	id := c.Param("id")
	var service models.Service
	if err := sc.DB.Where("id = ?", id).First(&service).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	if err := c.ShouldBindJSON(&service); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	sc.DB.Save(&service)
	c.JSON(http.StatusOK, service)
}

func (sc *ServiceController) DeleteService(c *gin.Context) {
	id := c.Param("id")
	var service models.Service
	if err := sc.DB.Where("id = ?", id).First(&service).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	sc.DB.Delete(&service)
	c.Status(http.StatusNoContent)
}
