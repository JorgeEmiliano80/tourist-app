package router

import (
	"tourist-app/internal/handlers"

	"github.com/gorilla/mux"
)

func SetupRouter(
	tourHandler *handlers.TourHandler,
	bookingHandler *handlers.BookingHandler,
	paymentHandler *handlers.PaymentHandler,
	clientHandler *handlers.ClientHandler,
) *mux.Router {
	r := mux.NewRouter()

	// Tour routes
	r.HandleFunc("/tours", tourHandler.GetAllTours).Methods("GET")
	r.HandleFunc("/tours/{id}", tourHandler.GetTour).Methods("GET")
	r.HandleFunc("/tours", tourHandler.CreateTour).Methods("POST")
	r.HandleFunc("/tours/{id}", tourHandler.UpdateTour).Methods("PUT")
	r.HandleFunc("/tours/{id}", tourHandler.DeleteTour).Methods("DELETE")

	// Booking routes
	r.HandleFunc("/bookings", bookingHandler.GetAllBookings).Methods("GET")
	r.HandleFunc("/bookings/{id}", bookingHandler.GetBooking).Methods("GET")
	r.HandleFunc("/bookings", bookingHandler.CreateBooking).Methods("POST")
	r.HandleFunc("/bookings/{id}", bookingHandler.UpdateBooking).Methods("PUT")
	r.HandleFunc("/bookings/{id}", bookingHandler.DeleteBooking).Methods("DELETE")

	// Payment routes
	r.HandleFunc("/payments", paymentHandler.GetAllPayments).Methods("GET")
	r.HandleFunc("/payments/{id}", paymentHandler.GetPayment).Methods("GET")
	r.HandleFunc("/payments", paymentHandler.CreatePayment).Methods("POST")
	r.HandleFunc("/payments/{id}", paymentHandler.UpdatePayment).Methods("PUT")
	r.HandleFunc("/payments/{id}", paymentHandler.DeletePayment).Methods("DELETE")

	// Client routes
	r.HandleFunc("/clients", clientHandler.GetAllClients).Methods("GET")
	r.HandleFunc("/clients/{id}", clientHandler.GetClient).Methods("GET")
	r.HandleFunc("/clients", clientHandler.CreateClient).Methods("POST")
	r.HandleFunc("/clients/{id}", clientHandler.UpdateClient).Methods("PUT")
	r.HandleFunc("/clients/{id}", clientHandler.DeleteClient).Methods("DELETE")

	return r
}
