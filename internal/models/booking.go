package models

import "time"

type Booking struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	ClientID    uint      `json:"client_id"`
	TourID      uint      `json:"tour_id"`
	BookingDate time.Time `json:"booking_date"`
	Status      string    `json:"status"`
	Lodging     string    `json:"lodging"`
	Comment     string    `json:"comment"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
