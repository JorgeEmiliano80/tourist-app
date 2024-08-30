package handlers

import (
	"net/http"

	"tourist-app/internal/db"
	"tourist-app/internal/models"

	"github.com/gin-gonic/gin"
)

func CreatePayment(c *gin.Context) {
	var payment models.Payment
	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos de pago inválidos"})
		return
	}

	if err := db.DB.Create(&payment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear el pago"})
		return
	}

	c.JSON(http.StatusCreated, payment)
}

func GetPayments(c *gin.Context) {
	var payments []models.Payment
	if err := db.DB.Find(&payments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los pagos"})
		return
	}

	c.JSON(http.StatusOK, payments)
}

func GetPayment(c *gin.Context) {
	id := c.Param("id")
	var payment models.Payment
	if err := db.DB.First(&payment, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pago no encontrado"})
		return
	}

	c.JSON(http.StatusOK, payment)
}

func UpdatePayment(c *gin.Context) {
	id := c.Param("id")
	var payment models.Payment
	if err := db.DB.First(&payment, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pago no encontrado"})
		return
	}

	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos de pago inválidos"})
		return
	}

	if err := db.DB.Save(&payment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar el pago"})
		return
	}

	c.JSON(http.StatusOK, payment)
}

func DeletePayment(c *gin.Context) {
	id := c.Param("id")
	if err := db.DB.Delete(&models.Payment{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar el pago"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"mensaje": "Pago eliminado exitosamente"})
}
