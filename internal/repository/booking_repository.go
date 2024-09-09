package repository

import (
	"database/sql"
	"time"
	"tourist-app/internal/models"
)

type BookingRepository struct {
	DB *sql.DB
}

func NewBookingRepository(db *sql.DB) *BookingRepository {
	return &BookingRepository{DB: db}
}

func (r *BookingRepository) GetBookings() ([]models.Booking, error) {
	rows, err := r.DB.Query("SELECT * FROM public.bookings") // Corregido "SELECTE" a "SELECT"
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bookings []models.Booking
	for rows.Next() {
		var booking models.Booking
		if err := rows.Scan(&booking.BookingID, &booking.ClientID, &booking.TourID, &booking.Date, &booking.Amount, &booking.CreatedAt, &booking.UpdatedAt); err != nil {
			return nil, err
		}
		bookings = append(bookings, booking)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return bookings, nil
}

func (r *BookingRepository) GetBooking(id int) (*models.Booking, error) {
	row := r.DB.QueryRow("SELECT * FROM public.bookings WHERE booking_id = $1", id)

	var booking models.Booking
	err := row.Scan(&booking.BookingID, &booking.ClientID, &booking.TourID, &booking.Date, &booking.Amount, &booking.CreatedAt, &booking.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Retorna nil si no se encuentra la reserva
		}
		return nil, err
	}
	return &booking, nil
}

func (r *BookingRepository) CreateBooking(booking *models.Booking) error {
	now := time.Now()
	booking.CreatedAt = now
	booking.UpdatedAt = now

	_, err := r.DB.Exec("INSERT INTO public.bookings (client_id, tour_id, date, amount, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)",
		booking.ClientID, booking.TourID, booking.Date, booking.Amount, booking.CreatedAt, booking.UpdatedAt)
	return err
}

func (r *BookingRepository) UpdateBooking(booking *models.Booking) error {
	booking.UpdatedAt = time.Now()

	_, err := r.DB.Exec("UPDATE public.bookings SET client_id = $1, tour_id = $2, date = $3, amount = $4, updated_at = $5 WHERE booking_id = $6",
		booking.ClientID, booking.TourID, booking.Date, booking.Amount, booking.UpdatedAt, booking.BookingID)
	return err
}

func (r *BookingRepository) DeleteBooking(id int) error {
	result, err := r.DB.Exec("DELETE FROM public.bookings WHERE booking_id = $1", id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows // Retorna un error si no se eliminó ninguna fila
	}

	return nil
}
