package models

import "time"

type Booking struct {
	BookingID int       `json:"booking_id"`
	ClientID  int       `json:"client_id"`
	TourID    int       `json:"tour_id"`
	Date      time.Time `json:"date"`
	Amount    float64   `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
