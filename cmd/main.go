package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"tourist-app/config"
	"tourist-app/internal/handlers"
	"tourist-app/internal/models"

	"github.com/gin-gonic/gin"
)

func init() {
	// Configuración de la base de datos
	dbHost := config.DBHost
	dbUser := config.DBUser
	dbPassword := config.DBPassword
	dbName := config.DBName
	dbPort := config.DBPort

	// Configuración del servidor
	serverPort := config.ServerPort

	fmt.Printf("Configuración cargada:\nDB Host: %s\nDB User: %s\nDB Name: %s\nDB Port: %s\nServer Port: %s\n",
		dbHost, dbUser, dbName, dbPort, serverPort)

	// Variables para configurar conexion con base de datos
	dbURI := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)
	fmt.Printf("DB URI: %s\n", dbURI)
}

func toJSON(v interface{}) string {
	jsonData, err := json.Marshal(v)
	if err != nil {
		return ""
	}
	return string(jsonData)
}

func getPaginationInfo(c *gin.Context, totalItems int) (int, int, []int) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	itemsPerPage := 10 // Puedes ajustar esto según tus necesidades
	totalPages := (totalItems + itemsPerPage - 1) / itemsPerPage

	if page < 1 {
		page = 1
	} else if page > totalPages {
		page = totalPages
	}

	pages := make([]int, 0)
	for i := 1; i <= totalPages; i++ {
		pages = append(pages, i)
	}

	return page, totalPages, pages
}

func main() {
	r := gin.Default()
	r.Use(gin.Logger())

	// Definir funciones personalizadas antes de cargar las plantillas
	r.SetFuncMap(template.FuncMap{
		"toJSON": toJSON,
		"sub": func(a, b int) int {
			return a - b
		},
		"add": func(a, b int) int {
			return a + b
		},
	})

	// Cargar plantillas HTML después de definir las funciones personalizadas
	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "./static")

	// Ruta para la página de inicio
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title":         "Bienvenido a nuestra aplicación de tours",
			"featuredTours": []models.Tour{},
		})
	})

	// Rutas de Booking
	r.POST("/bookings", handlers.CreateBooking)
	r.GET("/bookings", func(c *gin.Context) {
		bookings := []models.Booking{} // Reemplaza esto con la obtención real de reservas
		totalBookings := len(bookings)
		currentPage, totalPages, pages := getPaginationInfo(c, totalBookings)

		c.HTML(http.StatusOK, "bookings.html", gin.H{
			"bookings":    bookings,
			"currentPage": currentPage,
			"pages":       pages,
			"totalPages":  totalPages,
		})
	})
	r.GET("/bookings/:id", handlers.GetBooking)
	r.PUT("/bookings/:id", handlers.UpdateBooking)
	r.DELETE("/bookings/:id", handlers.DeleteBooking)

	// Rutas de Payment
	r.POST("/payments", handlers.CreatePayment)
	r.GET("/payments", func(c *gin.Context) {
		payments := []models.Payment{} // Reemplaza esto con la obtención real de pagos
		totalPayments := len(payments)
		currentPage, totalPages, pages := getPaginationInfo(c, totalPayments)

		c.HTML(http.StatusOK, "payments.html", gin.H{
			"payments":    payments,
			"currentPage": currentPage,
			"pages":       pages,
			"totalPages":  totalPages,
		})
	})
	r.GET("/payments/:id", handlers.GetPayment)
	r.PUT("/payments/:id", handlers.UpdatePayment)
	r.DELETE("/payments/:id", handlers.DeletePayment)

	// Rutas de Client
	r.POST("/clients", handlers.CreateClient)
	r.GET("/clients", func(c *gin.Context) {
		clients := []models.Client{} // Reemplaza esto con la obtención real de clientes
		totalClients := len(clients)
		currentPage, totalPages, pages := getPaginationInfo(c, totalClients)

		c.HTML(http.StatusOK, "clients.html", gin.H{
			"clients":     clients,
			"currentPage": currentPage,
			"pages":       pages,
			"totalPages":  totalPages,
		})
	})
	r.GET("/clients/:id", handlers.GetClient)
	r.PUT("/clients/:id", handlers.UpdateClient)
	r.DELETE("/clients/:id", handlers.DeleteClient)

	// Rutas de Tour
	r.POST("/tours", handlers.Createtour)
	r.GET("/tours", func(c *gin.Context) {
		tours := []models.Tour{} // Reemplaza esto con la obtención real de tours
		totalTours := len(tours)
		currentPage, totalPages, pages := getPaginationInfo(c, totalTours)

		c.HTML(http.StatusOK, "tours.html", gin.H{
			"tours":       tours,
			"currentPage": currentPage,
			"pages":       pages,
			"totalPages":  totalPages,
		})
	})
	r.GET("/tours/:id", handlers.GetTour)
	r.PUT("/tours/:id", handlers.UpdateTour)
	r.DELETE("/tours/:id", handlers.DeleteTour)

	// Rutas de manejador
	r.GET("/bookings-panel", func(c *gin.Context) {
		recentBookings := []models.Booking{} // Reemplaza esto con la obtención real de reservas recientes
		totalBookings := 100                 // Reemplaza con el total real de reservas
		currentPage, totalPages, pages := getPaginationInfo(c, totalBookings)

		trendLabels := []string{"Ene", "Feb", "Mar", "Abr", "May"}
		trendData := []int{10, 15, 13, 17, 20}
		popularTourNames := []string{"Tour A", "Tour B", "Tour C"}
		popularTourBookings := []int{50, 30, 20}

		c.HTML(http.StatusOK, "bookings.html", gin.H{
			"title":                   "Panel de Reservas",
			"totalBookings":           totalBookings,
			"totalRevenue":            5000,
			"occupancyRate":           75,
			"pendingBookings":         5,
			"recentBookings":          recentBookings,
			"currentPage":             currentPage,
			"pages":                   pages,
			"totalPages":              totalPages,
			"trendLabels":             trendLabels,
			"trendData":               trendData,
			"popularTourNames":        popularTourNames,
			"popularTourBookings":     popularTourBookings,
			"trendLabelsJSON":         toJSON(trendLabels),
			"trendDataJSON":           toJSON(trendData),
			"popularTourNamesJSON":    toJSON(popularTourNames),
			"popularTourBookingsJSON": toJSON(popularTourBookings),
		})
	})

	// Iniciar el servidor
	serverPort := ":" + config.ServerPort
	if err := r.Run(serverPort); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
