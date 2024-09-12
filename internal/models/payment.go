package models

import "time"

type Payment struct {
	ID        int       `json:"id"`
	BookingID int       `json:"booking_id"`
	Amount    float64   `json:"amount"`
	Method    string    `json:"method"`
	Status    string    `json:"status"`
	PaidAt    time.Time `json:"paid_at"`
}
