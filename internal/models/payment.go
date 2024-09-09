package models

import "time"

type Payment struct {
	PaymentID int       `json:"payment_id"`
	BookingID int       `json:"booking_id"`
	Amount    float64   `json:"amount"`
	Date      time.Time `json:"date"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
