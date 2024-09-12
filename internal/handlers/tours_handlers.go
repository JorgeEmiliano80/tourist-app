package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"tourist-app/internal/models"
	"tourist-app/internal/services"

	"github.com/gorilla/mux"
)

type TourHandler struct {
	service *services.TourService
}

func NewTourHandler(service *services.TourService) *TourHandler {
	return &TourHandler{service: service}
}

func (h *TourHandler) GetAllTours(w http.ResponseWriter, r *http.Request) {
	tours, err := h.service.GetAllTours()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(tours)
}

func (h *TourHandler) GetTour(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid tour ID", http.StatusBadRequest)
		return
	}

	tour, err := h.service.GetTourByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(tour)
}

func (h *TourHandler) CreateTour(w http.ResponseWriter, r *http.Request) {
	var tour models.Tour
	err := json.NewDecoder(r.Body).Decode(&tour)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.service.CreateTour(&tour)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *TourHandler) UpdateTour(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid tour ID", http.StatusBadRequest)
		return
	}

	var tour models.Tour
	err = json.NewDecoder(r.Body).Decode(&tour)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tour.ID = id
	err = h.service.UpdateTour(&tour)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *TourHandler) DeleteTour(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid tour ID", http.StatusBadRequest)
		return
	}

	err = h.service.DeleteTour(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
