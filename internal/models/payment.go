package models

import "time"

type Payment struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	BookingID uint      `json:"booking_id"`
	Amount    float64   `json:"amount"`
	Currency  string    `json:"currency"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
