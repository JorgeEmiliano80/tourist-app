package handlers

import (
	"net/http"

	"tourist-app/internal/db"
	"tourist-app/internal/models"

	"github.com/gin-gonic/gin"
)

func CreateClient(c *gin.Context) {
	var client models.Client
	if err := c.ShouldBindJSON(&client); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.DB.Create(&client).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create client"})
		return
	}

	c.JSON(http.StatusCreated, client)
}

func GetClients(c *gin.Context) {
	var clients []models.Client
	if err := db.DB.Find(&clients).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch clients"})
		return
	}

	c.JSON(http.StatusOK, clients)
}

func GetClient(c *gin.Context) {
	id := c.Param("id")
	var client models.Client
	if err := db.DB.First(&client, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Clinet not found"})
		return
	}

	c.JSON(http.StatusOK, client)
}

func UpdateClient(c *gin.Context) {
	id := c.Param("id")
	var client models.Client
	if err := db.DB.First(&client, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Client not found"})
		return
	}

	if err := c.ShouldBindJSON(&client); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.DB.Save(&client)
	c.JSON(http.StatusOK, client)
}

func DeleteClient(c *gin.Context) {
	id := c.Param("id")
	if err := db.DB.Delete(&models.Client{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete client"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Client deleted successfully"})
}
