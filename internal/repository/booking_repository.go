package repository

import (
	"database/sql"

	"tourist-app/internal/models"
)

type BookingRepository struct {
	db *sql.DB
}

func NewBookingRepository(db *sql.DB) *BookingRepository {
	return &BookingRepository{db: db}
}

func (r *BookingRepository) GetAll() ([]models.Booking, error) {
	rows, err := r.db.Query("SELECT id, tour_id, client_id, booked_at, status, quantity, total_cost FROM bookings")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bookings []models.Booking
	for rows.Next() {
		var b models.Booking
		err := rows.Scan(&b.ID, &b.TourID, &b.ClientID, &b.BookedAt, &b.Status, &b.Quantity, &b.TotalCost)
		if err != nil {
			return nil, err
		}
		bookings = append(bookings, b)
	}
	return bookings, nil
}

func (r *BookingRepository) GetByID(id int) (*models.Booking, error) {
	var b models.Booking
	err := r.db.QueryRow("SELECT id, tour_id, client_id, booked_at, status, quantity, total_cost FROM bookings WHERE id = $1", id).
		Scan(&b.ID, &b.TourID, &b.ClientID, &b.BookedAt, &b.Status, &b.Quantity, &b.TotalCost)
	if err != nil {
		return nil, err
	}
	return &b, nil
}

func (r *BookingRepository) Create(booking *models.Booking) error {
	_, err := r.db.Exec("INSERT INTO bookings (tour_id, client_id, booked_at, status, quantity, total_cost) VALUES ($1, $2, $3, $4, $5, $6)",
		booking.TourID, booking.ClientID, booking.BookedAt, booking.Status, booking.Quantity, booking.TotalCost)
	return err
}

func (r *BookingRepository) Update(booking *models.Booking) error {
	_, err := r.db.Exec("UPDATE bookings SET tour_id = $1, client_id = $2, booked_at = $3, status = $4, quantity = $5, total_cost = $6 WHERE id = $7",
		booking.TourID, booking.ClientID, booking.BookedAt, booking.Status, booking.Quantity, booking.TotalCost, booking.ID)
	return err
}

func (r *BookingRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM bookings WHERE id = $1", id)
	return err
}
