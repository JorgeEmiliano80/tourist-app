package services

import (
	"tourist-app/internal/models"
	"tourist-app/internal/repository"
)

type PaymentService struct {
	Repo *repository.PaymentRepository
}

func NewPaymentService(repo *repository.PaymentRepository) *PaymentService {
	return &PaymentService{Repo: repo}
}

func (s *PaymentService) GetPayments() ([]models.Payment, error) {
	return s.Repo.GetPayments()
}

func (s *PaymentService) GetPayment(id int) (*models.Payment, error) {
	return s.Repo.GetPayment(id)
}

func (s *PaymentService) CreatePayment(payment *models.Payment) error {
	return s.Repo.CreatePayment(payment)
}

func (s *PaymentService) UpdatePayment(payment *models.Payment) error {
	return s.Repo.UpdatePayment(payment)
}

func (s *PaymentService) DeletePayment(id int) error {
	return s.Repo.DeletePayment(id)
}
