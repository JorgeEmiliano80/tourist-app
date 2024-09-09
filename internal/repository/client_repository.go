package repository

import (
	"database/sql"
	"tourist-app/internal/models"
)

type ClientRepository struct {
	DB *sql.DB
}

func NewClientRepository(db *sql.DB) *ClientRepository {
	return &ClientRepository{DB: db}
}

func (r *ClientRepository) GetClients() ([]models.Client, error) {
	rows, err := r.DB.Query("SELECT * FROM public.clients")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var clients []models.Client
	for rows.Next() {
		var client models.Client
		if err := rows.Scan(&client.ClientID, &client.FirstName, &client.LastName, &client.Email, &client.Phone, &client.CreatedAt, &client.UpdatedAt); err != nil {
			return nil, err
		}
		clients = append(clients, client)
	}
	return clients, nil
}

func (r *ClientRepository) GetClient(id int) (*models.Client, error) {
	row := r.DB.QueryRow("SELECT * FROM public.clients WHERE client_id = $1", id)

	var client models.Client
	if err := row.Scan(&client.ClientID, &client.FirstName, &client.LastName, &client.Email, &client.Phone, &client.CreatedAt, &client.UpdatedAt); err != nil {
		return nil, err
	}
	return &client, nil
}

func (r *ClientRepository) CreateClient(client *models.Client) error {
	_, err := r.DB.Exec("INSERT INTO public.clients (first_name, last_name, email, phone,created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)",
		client.FirstName, client.LastName, client.Email, client.Phone, client.CreatedAt, client.UpdatedAt)
	return err
}

func (r *ClientRepository) UpdateClient(client *models.Client) error {
	_, err := r.DB.Exec("UPDATE public.clients SET first_name = $1, last_name = $2, email = $3, phone = $4, updated_at = $5, WHERE client_id = $6",
		client.FirstName, client.LastName, client.Email, client.Phone, client.UpdatedAt, client.ClientID)
	return err
}

func (r *ClientRepository) DeleteClient(id int) error {
	_, err := r.DB.Exec("DELETE FROM public.clients WHERE client_id = $1", id)
	return err
}
