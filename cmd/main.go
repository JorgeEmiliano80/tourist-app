package main

import (
	"log"
	"net/http"

	"tourist-app/config"
	"tourist-app/database"
	"tourist-app/internal/handlers"
	"tourist-app/internal/repository"
	"tourist-app/internal/services"

	"github.com/gorilla/mux"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Initialize the database
	db, err := database.InitDB(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close() // Asegúrate de cerrar la conexión cuando el programa termine

	// Create repositories
	tourRepo := repository.NewTourRepository(db)
	bookingRepo := repository.NewBookingRepository(db)
	paymentRepo := repository.NewPaymentRepository(db)
	clientRepo := repository.NewClientRepository(db)

	// Create services
	tourService := services.NewTourService(tourRepo)
	bookingService := services.NewBookingService(bookingRepo)
	paymentService := services.NewPaymentService(paymentRepo)
	clientService := services.NewClientService(clientRepo)

	// Create handlers
	tourHandler := handlers.NewTourHandler(tourService)
	bookingHandler := handlers.NewBookingHandler(bookingService)
	paymentHandler := handlers.NewPaymentHandler(paymentService)
	clientHandler := handlers.NewClientHandler(clientService)

	// Setup router
	r := mux.NewRouter()

	// Tours routes
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
	r.HandleFunc("payments/{id}", paymentHandler.UpdatePayment).Methods("PUT")
	r.HandleFunc("/payments/{id}", paymentHandler.DeletePayment).Methods("DELETE")

	// Client routes
	r.HandleFunc("/clients", clientHandler.GetAllClients).Methods("GET")
	r.HandleFunc("/clients/{id}", clientHandler.GetAllClients).Methods("GET")
	r.HandleFunc("/clients", clientHandler.CreateClient).Methods("POST")
	r.HandleFunc("/clients/{id}", clientHandler.UpdateClient).Methods("PUT")
	r.HandleFunc("/clients/{id}", clientHandler.DeleteClient).Methods("DELETE")

	// Start server
	log.Printf("Server starting on port %s", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, r))
}
