package repository

import (
	"database/sql"

	"tourist-app/internal/models"
)

type TourRepository struct {
	db *sql.DB
}

func NewTourRepository(db *sql.DB) *TourRepository {
	return &TourRepository{db: db}
}

func (r *TourRepository) GetAll() ([]models.Tour, error) {
	rows, err := r.db.Query("SELECT id, name, description, price, start-date, end_date, capacity FROM tours")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tours []models.Tour
	for rows.Next() {
		var t models.Tour
		err := rows.Scan(&t.ID, &t.Name, &t.Description, &t.Price, &t.StartDate, &t.EndDate, &t.Capacity)
		if err != nil {
			return nil, err
		}
		tours = append(tours, t)
	}
	return tours, nil
}

func (r *TourRepository) GetByID(id int) (*models.Tour, error) {
	var t models.Tour
	err := r.db.QueryRow("SELECT id, name, description, price, start_date, end_date, capacity FROM tours WHERE id = $1", id).
		Scan(&t.ID, &t.Name, &t.Description, &t.Price, &t.StartDate, &t.EndDate, &t.Capacity)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func (r *TourRepository) Create(tour *models.Tour) error {
	_, err := r.db.Exec("INSERT INTO tour (name, description, price, start_date, end_date, capacity) VALUES ($1, $2, $3, $4, $5, $6)",
		tour.Name, tour.Description, tour.Price, tour.StartDate, tour.EndDate, tour.Capacity, tour.ID)
	return err
}

func (r *TourRepository) Update(tour *models.Tour) error {
	_, err := r.db.Exec("UPDATE tours SET name = $1, description = $2, price = $3, start_date = $4, end_date = $5, capacity = $6 WHERE id = $7",
		tour.Name, tour.Description, tour.Price, tour.StartDate, tour.EndDate, tour.Capacity, tour.ID)
	return err
}

func (r *TourRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM tours WHERE id = $1", id)
	return err
}
