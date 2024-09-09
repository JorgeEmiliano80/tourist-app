package main

import (
	"log"
	"tourist-app/database"
	"tourist-app/internal/handlers"
	"tourist-app/internal/repository"
	"tourist-app/internal/services"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the database
	db, err := database.InitDB()
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
	router := gin.Default()

	// Tours routes
	router.GET("/tours", tourHandler.GetTours)
	router.GET("/tours/:id", tourHandler.GetTour)
	router.POST("/tours", tourHandler.CreateTour)
	router.PUT("/tours/:id", tourHandler.UpdateTour)
	router.DELETE("/tours/:id", tourHandler.DeleteTour)

	// Bookings routes
	router.GET("/bookings", bookingHandler.GetBookings)
	router.GET("/bookings/:id", bookingHandler.GetBooking)
	router.POST("/bookings", bookingHandler.CreateBooking)
	router.PUT("/bookings/:id", bookingHandler.UpdateBooking)
	router.DELETE("/bookings/:id", bookingHandler.DeleteBooking)

	// Payments routes
	router.GET("/payments", paymentHandler.GetPayments)
	router.GET("/payments/:id", paymentHandler.GetPayment)
	router.POST("/payments", paymentHandler.CreatePayment)
	router.PUT("/payments/:id", paymentHandler.UpdatePayment)
	router.DELETE("/payments/:id", paymentHandler.DeletePayment)

	// Clients routes
	router.GET("/clients", clientHandler.GetClients)
	router.GET("/clients/:id", clientHandler.GetClient)
	router.POST("/clients", clientHandler.CreateClient)
	router.PUT("/clients/:id", clientHandler.UpdateClient)
	router.DELETE("/clients/:id", clientHandler.DeleteClient)

	// Start server
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
