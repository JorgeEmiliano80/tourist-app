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

func (s *BookingService) GetAllBookings() ([]models.Booking, error) {
	return s.Repo.GetAll()
}

func (s *BookingService) GetBookingByID(id int) (*models.Booking, error) {
	return s.Repo.GetByID(id)
}

func (s *BookingService) CreateBooking(booking *models.Booking) error {
	return s.Repo.Create(booking)
}

func (s *BookingService) UpdateBooking(booking *models.Booking) error {
	return s.Repo.Update(booking)
}

func (s *BookingService) DeleteBooking(id int) error {
	return s.Repo.Delete(id)
}
