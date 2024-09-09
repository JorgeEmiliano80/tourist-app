package repository

import (
	"database/sql"
	"tourist-app/internal/models"
)

type TourRepository struct {
	DB *sql.DB
}

func NewTourRepository(db *sql.DB) *TourRepository {
	return &TourRepository{DB: db}
}

func (r *TourRepository) GetTours() ([]models.Tour, error) {
	rows, err := r.DB.Query("SELECT * FROM public.tours")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tours []models.Tour
	for rows.Next() {
		var tour models.Tour
		if err := rows.Scan(&tour.TourID, &tour.Name, &tour.Description, &tour.Price, &tour.CreatedAt, &tour.UpdatedAt); err != nil {
			return nil, err
		}
		tours = append(tours, tour)
	}
	return tours, nil
}

func (r *TourRepository) GetTour(id int) (*models.Tour, error) {
	row := r.DB.QueryRow("SELECT * FROM public.tours WHERE tour_id = $1", id)

	var tour models.Tour
	if err := row.Scan(&tour.TourID, &tour.Name, &tour.Description, &tour.Price, &tour.CreatedAt, &tour.UpdatedAt); err != nil {
		return nil, err
	}
	return &tour, nil
}

func (r *TourRepository) CreateTour(tour *models.Tour) error {
	_, err := r.DB.Exec("INSERT INTO public.tours (name, description, price, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)",
		tour.Name, tour.Description, tour.Price, tour.CreatedAt, tour.UpdatedAt)
	return err
}

func (r *TourRepository) UpdateTour(tour *models.Tour) error {
	_, err := r.DB.Exec("UPDATE public.tours SET name = $1, description = $2, price = $3, updated_at = $4 WHERE tour_id = $5",
		tour.Name, tour.Description, tour.Price, tour.UpdatedAt, tour.TourID)
	return err
}

func (r *TourRepository) DeleteTour(id int) error {
	_, err := r.DB.Exec("DELETE FROM public.tours WHERE tour_id = $1", id)
	return err
}
