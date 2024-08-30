package handlers

import (
	"net/http"

	"tourist-app/internal/db"
	"tourist-app/internal/models"

	"github.com/gin-gonic/gin"
)

func CreateBooking(c *gin.Context) {
	var booking models.Booking
	if err := c.ShouldBindJSON(&booking); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.DB.Create(&booking).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create booking"})
		return
	}

	c.JSON(http.StatusCreated, booking)
}

func GetBookings(c *gin.Context) {
	var bookings []models.Booking
	if err := db.DB.Find(&bookings).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch bookings"})
		return
	}

	c.JSON(http.StatusOK, bookings)
}

func GetBooking(c *gin.Context) {
	id := c.Param("id")
	var booking models.Booking
	if err := db.DB.First(&booking, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Booking not Found"})
		return
	}

	c.JSON(http.StatusOK, booking)
}

func UpdateBooking(c *gin.Context) {
	id := c.Param("id")
	var booking models.Booking
	if err := db.DB.First(&booking, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	db.DB.Save(&booking)
	c.JSON(http.StatusOK, booking)
}

func DeleteBooking(c *gin.Context) {
	id := c.Param("id")
	if err := db.DB.Delete(&models.Booking{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete booking"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Booking deleted successfully"})
}
