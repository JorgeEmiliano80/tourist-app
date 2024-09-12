package repository

import (
	"database/sql"
	"tourist-app/internal/models"
)

type ClientRepository struct {
	db *sql.DB
}

func NewClientRepository(db *sql.DB) *ClientRepository {
	return &ClientRepository{db: db}
}

func (r *ClientRepository) GetAll() ([]models.Client, error) {
	rows, err := r.db.Query("SELECT id, first_name, last_name, email, phone FROM clients")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var clients []models.Client
	for rows.Next() {
		var c models.Client
		err := rows.Scan(&c.ID, &c.FirstName, &c.LastName, &c.Email, &c.Phone)
		if err != nil {
			return nil, err
		}
		clients = append(clients, c)
	}
	return clients, nil
}

func (r *ClientRepository) GetByID(id int) (*models.Client, error) {
	var c models.Client
	err := r.db.QueryRow("SELECT id, first_name, last_name, email, phone FROM clients WHERE id = $1", id).
		Scan(&c.ID, &c.FirstName, &c.LastName, &c.Email, &c.Phone)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *ClientRepository) Create(client *models.Client) error {
	_, err := r.db.Exec("INSERT INTO clients (first_name, last_name, email, phone) VALUES ($1, $2, $3, $4)",
		client.FirstName, client.LastName, client.Email, client.Phone)
	return err
}

func (r *ClientRepository) Update(client *models.Client) error {
	_, err := r.db.Exec("UPDATE clients SET first_name = $1, last_name = $2, email = $3, phone = $4, WHERE id = $5",
		client.FirstName, client.LastName, client.Email, client.Phone, client.ID)
	return err
}

func (r *ClientRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM clients WHERE id = $1", id)
	return err
}
