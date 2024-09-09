package services

import (
	"tourist-app/internal/models"
	"tourist-app/internal/repository"
)

type TourService struct {
	Repo *repository.TourRepository
}

func NewTourService(repo *repository.TourRepository) *TourService {
	return &TourService{Repo: repo}
}

func (s *TourService) GetTours() ([]models.Tour, error) {
	return s.Repo.GetTours()
}

func (s *TourService) GetTour(id int) (*models.Tour, error) {
	return s.Repo.GetTour(id)
}

func (s *TourService) CreateTour(tour *models.Tour) error {
	return s.Repo.CreateTour(tour)
}

func (s *TourService) UpdateTour(tour *models.Tour) error {
	return s.Repo.UpdateTour(tour)
}

func (s *TourService) DeleteTour(id int) error {
	return s.Repo.DeleteTour(id)
}
