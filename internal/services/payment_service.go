package services

import (
	"tourist-app/internal/models"
	"tourist-app/internal/repository"
)

type PaymentService struct {
	repo *repository.PaymentRepository
}

func NewPaymentService(repo *repository.PaymentRepository) *PaymentService {
	return &PaymentService{repo: repo}
}

func (s *PaymentService) GetAllPayments() ([]models.Payment, error) {
	return s.repo.GetAll()
}

func (s *PaymentService) GetPaymentByID(id int) (*models.Payment, error) {
	return s.repo.GetByID(id)
}

func (s *PaymentService) CreatePayment(payment *models.Payment) error {
	return s.repo.Create(payment)
}

func (s *PaymentService) UpdatePayment(payment *models.Payment) error {
	return s.repo.Update(payment)
}

func (s *PaymentService) DeletePayment(id int) error {
	return s.repo.Delete(id)
}
