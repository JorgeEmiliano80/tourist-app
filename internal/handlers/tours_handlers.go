package handlers

import (
	"net/http"
	"strconv"
	"tourist-app/internal/models"
	"tourist-app/internal/services"

	"github.com/gin-gonic/gin"
)

type TourHandler struct {
	Service *services.TourService
}

func NewTourHandler(service *services.TourService) *TourHandler {
	return &TourHandler{Service: service}
}

func (h *TourHandler) GetTours(c *gin.Context) {
	tours, err := h.Service.GetTours()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tours)
}

func (h *TourHandler) GetTour(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	tour, err := h.Service.GetTour(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tour not found"})
		return
	}
	c.JSON(http.StatusOK, tour)
}

func (h *TourHandler) CreateTour(c *gin.Context) {
	var tour models.Tour
	if err := c.ShouldBindJSON(&tour); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.Service.CreateTour(&tour); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, tour)
}

func (h *TourHandler) UpdateTour(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var tour models.Tour
	if err := c.ShouldBindJSON(&tour); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tour.TourID = id
	if err := h.Service.UpdateTour(&tour); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tour)
}

func (h *TourHandler) DeleteTour(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.Service.DeleteTour(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
