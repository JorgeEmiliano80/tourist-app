package database

import (
	"time"
)

// User represents a user of the tourist application
type User struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"` // The dash indicates this field won't be serialized to JSON
}

// Destination represents a tourist location
type Destination struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Location    string  `json:"location"`
	Price       float64 `json:"price"`
}

// Booking represents a reservation made by a user for a destination
type Booking struct {
	ID            int64     `json:"id"`
	UserID        int64     `json:"user_id"`
	DestinationID int64     `json:"destination_id"`
	StartDate     time.Time `json:"start_date"`
	EndDate       time.Time `json:"end_date"`
	Status        string    `json:"status"` // For example: "confirmed", "pending", "cancelled"
}

// Review represents a review left by a user for a destination
type Review struct {
	ID            int64     `json:"id"`
	UserID        int64     `json:"user_id"`
	DestinationID int64     `json:"destination_id"`
	Content       string    `json:"content"`
	Rating        int       `json:"rating"` // For example, 1 to 5 stars
	Date          time.Time `json:"date"`
}

// ValidateCredentials checks if the provided email and password match the user's credentials
func (u *User) ValidateCredentials(email, password string) bool {
	return u.Email == email && u.Password == password
}

// UpdatePrice changes the price of a destination
func (d *Destination) UpdatePrice(newPrice float64) {
	d.Price = newPrice
}

// CancelBooking changes the status of a booking to "cancelled"
func (b *Booking) CancelBooking() {
	b.Status = "cancelled"
}

// ConfirmBooking changes the status of a booking to "confirmed"
func (b *Booking) ConfirmBooking() {
	b.Status = "confirmed"
}

// IsValid checks if a review has a valid rating and non-empty content
func (r *Review) IsValid() bool {
	return r.Rating >= 1 && r.Rating <= 5 && r.Content != ""
}
