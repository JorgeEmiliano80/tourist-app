package models

import "time"

type Booking struct {
	ID        int       `json:"id"`
	ClientID  int       `json:"client_id"`
	TourID    int       `json:"tour_id"`
	BookedAt  time.Time `json:booked_at`
	Status    string    `json:"status"`
	Quantity  int       `json:"quantity"`
	TotalCost float64   `json:"total_cost"`
}
