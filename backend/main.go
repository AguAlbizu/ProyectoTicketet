package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"ticketapp/controllers"
	"ticketapp/dao"
	"ticketapp/domain"
	"ticketapp/middleware"
	"ticketapp/services"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, usando variables de entorno del sistema")
	}

	// Construir DSN desde variables de entorno
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error al conectar con la base de datos:", err)
	}

	// AutoMigrate crea/actualiza las tablas automáticamente
	if err := db.AutoMigrate(&domain.User{}, &domain.Event{}, &domain.Ticket{}); err != nil {
		log.Fatal("AutoMigrate falló:", err)
	}

	// Instanciar DAOs
	userDAO := dao.NewUserDAO(db)
	eventDAO := dao.NewEventDAO(db)
	ticketDAO := dao.NewTicketDAO(db)

	// Instanciar services
	authService := services.NewAuthService(userDAO)
	eventService := services.NewEventService(eventDAO)
	ticketService := services.NewTicketService(ticketDAO, eventDAO, userDAO)

	// Instanciar controllers
	authController := controllers.NewAuthController(authService)
	eventController := controllers.NewEventController(eventService)
	ticketController := controllers.NewTicketController(ticketService)

	r := gin.Default()

	// CORS: permitir peticiones desde el frontend en localhost:5173
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "http://localhost:5173")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	})

	api := r.Group("/api")

	// Rutas públicas
	api.POST("/auth/register", authController.Register)
	api.POST("/auth/login", authController.Login)
	api.GET("/events", eventController.GetEvents)
	api.GET("/events/:id", eventController.GetEventByID)

	// Rutas protegidas — requieren JWT válido
	protected := api.Group("")
	protected.Use(middleware.AuthMiddleware())
	protected.POST("/tickets", ticketController.BuyTicket)
	protected.GET("/tickets/my-tickets", ticketController.GetMyTickets)
	protected.DELETE("/tickets/:id", ticketController.CancelTicket)
	protected.PUT("/tickets/:id/transfer", ticketController.TransferTicket)

	// TODO (entrega final): agregar rutas de administrador con middleware de rol

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Servidor corriendo en http://localhost:%s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal("Error al iniciar el servidor:", err)
	}
}
