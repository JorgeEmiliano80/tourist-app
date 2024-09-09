package services

import (
	"tourist-app/internal/models"
	"tourist-app/internal/repository"
)

type ClientService struct {
	Repo *repository.ClientRepository
}

func NewClientService(repo *repository.ClientRepository) *ClientService {
	return &ClientService{Repo: repo}
}

func (s *ClientService) GetClients() ([]models.Client, error) {
	return s.Repo.GetClients()
}

func (s *ClientService) GetClient(id int) (*models.Client, error) {
	return s.Repo.GetClient(id)
}

func (s *ClientService) CreateClient(client *models.Client) error {
	return s.Repo.CreateClient(client)
}

func (s *ClientService) UpdateClient(client *models.Client) error {
	return s.Repo.UpdateClient(client)
}

func (s *ClientService) DeleteClient(id int) error {
	return s.Repo.DeleteClient(id)
}
