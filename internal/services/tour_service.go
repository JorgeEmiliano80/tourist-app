package services

import (
	"tourist-app/internal/models"
	"tourist-app/internal/repository"
)

type TourService struct {
	repo *repository.TourRepository
}

func NewTourService(repo *repository.TourRepository) *TourService {
	return &TourService{repo: repo}
}

func (s *TourService) GetAllTours() ([]models.Tour, error) {
	return s.repo.GetAll()
}

func (s *TourService) GetTourByID(id int) (*models.Tour, error) {
	return s.repo.GetByID(id)
}

func (s *TourService) CreateTour(tour *models.Tour) error {
	return s.repo.Create(tour)
}

func (s *TourService) UpdateTour(tour *models.Tour) error {
	return s.repo.Update(tour)
}

func (s *TourService) DeleteTour(id int) error {
	return s.repo.Delete(id)
}
