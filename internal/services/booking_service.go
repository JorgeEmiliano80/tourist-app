package services

import (
	"tourist-app/internal/models"
	"tourist-app/internal/repository"
)

type BookingService struct {
	Repo *repository.BookingRepository
}

func NewBookingService(repo *repository.BookingRepository) *BookingService {
	return &BookingService{Repo: repo}
}

func (s *BookingService) GetBookings() ([]models.Booking, error) {
	return s.Repo.GetBookings()
}

func (s *BookingService) GetBooking(id int) (*models.Booking, error) {
	return s.Repo.GetBooking(id)
}

func (s *BookingService) CreateBooking(booking *models.Booking) error {
	return s.Repo.CreateBooking(booking)
}

func (s *BookingService) UpdateBooking(booking *models.Booking) error {
	return s.Repo.UpdateBooking(booking)
}

func (s *BookingService) DeleteBooking(id int) error {
	return s.Repo.DeleteBooking(id)
}
