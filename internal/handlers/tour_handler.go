package handlers

import (
	"net/http"

	"tourist-app/internal/db"
	"tourist-app/internal/models"

	"github.com/gin-gonic/gin"
)

func Createtour(c *gin.Context) {
	var tour models.Tour
	if err := c.ShouldBindJSON(&tour); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.DB.Create(&tour).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create tour"})
		return
	}

	c.JSON(http.StatusCreated, tour)
}

func GetTours(c *gin.Context) {
	var tours []models.Tour
	if err := db.DB.Find(&tours).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tours"})
		return
	}

	c.JSON(http.StatusOK, tours)
}

func GetTour(c *gin.Context) {
	id := c.Param("id")
	var tour models.Tour
	if err := db.DB.First(&tour, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tour not found"})
		return
	}

	c.JSON(http.StatusOK, tour)
}

func UpdateTour(c *gin.Context) {
	id := c.Param("id")
	var tour models.Tour
	if err := db.DB.First(&tour, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tour not found"})
		return
	}

	if err := c.ShouldBindJSON(&tour); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.DB.Save(&tour)
	c.JSON(http.StatusOK, tour)
}

func DeleteTour(c *gin.Context) {
	id := c.Param("id")
	if err := db.DB.Delete(&models.Tour{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete tour"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Tour deleted successfully"})
}
