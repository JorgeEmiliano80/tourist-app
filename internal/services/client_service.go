package services

import (
	"tourist-app/internal/models"
	"tourist-app/internal/repository"
)

type ClientService struct {
	repo *repository.ClientRepository
}

func NewClientService(repo *repository.ClientRepository) *ClientService {
	return &ClientService{repo: repo}
}

func (s *ClientService) GetAllClients() ([]models.Client, error) {
	return s.repo.GetAll()
}

func (s *ClientService) GetClientByID(id int) (*models.Client, error) {
	return s.repo.GetByID(id)
}

func (s *ClientService) CreateClient(client *models.Client) error {
	return s.repo.Create(client)
}

func (s *ClientService) UpdateClient(client *models.Client) error {
	return s.repo.Update(client)
}

func (s *ClientService) DeleteClient(id int) error {
	return s.repo.Delete(id)
}
