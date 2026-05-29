package main

import (
	"fmt"
	"log"
	"os"
	"ticketapp/controllers"
	"ticketapp/dao"
	"ticketapp/domain"
	"ticketapp/services"
	"ticketapp/utils"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// Load .env file (ignorado en producción si las vars ya están inyectadas)
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	// Build DSN from environment variables
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	// Connect to database
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto-migrate domain models
	if err := db.AutoMigrate(&domain.User{}, &domain.Event{}, &domain.Ticket{}); err != nil {
		log.Fatal("AutoMigrate failed:", err)
	}

	// Instantiate DAOs
	userDAO := dao.NewUserDAO(db)
	eventDAO := dao.NewEventDAO(db)
	ticketDAO := dao.NewTicketDAO(db)

	// Instantiate services
	authService := services.NewAuthService(userDAO)
	eventService := services.NewEventService(eventDAO)
	ticketService := services.NewTicketService(ticketDAO, eventDAO, userDAO)

	// Instantiate controllers
	authController := controllers.NewAuthController(authService)
	eventController := controllers.NewEventController(eventService)
	ticketController := controllers.NewTicketController(ticketService)

	// Initialize Gin router
	r := gin.Default()

	// TODO (entrega final): agregar middleware CORS

	api := r.Group("/api")

	// JWT middleware — verifica firma y expiración, sin validación de rol
	// NOTE (entrega parcial): solo autentica, no autoriza por rol
	jwtMiddleware := func(c *gin.Context) {
		// TODO: extraer token del header Authorization: Bearer <token>
		// TODO: llamar utils.ValidateToken(tokenString)
		// TODO: si inválido, c.AbortWithStatusJSON(401, gin.H{"error": "unauthorized"})
		// TODO: guardar userID en contexto: c.Set("userID", claims.UserID)
		_ = utils.ValidateToken
		c.Next()
	}

	// Rutas públicas (sin autenticación)
	api.POST("/auth/register", authController.Register)
	api.POST("/auth/login", authController.Login)
	api.GET("/events", eventController.GetEvents)
	api.GET("/events/:id", eventController.GetEventByID)

	// Rutas protegidas (requieren JWT válido, sin validación de rol)
	protected := api.Group("")
	protected.Use(jwtMiddleware)
	protected.POST("/tickets", ticketController.BuyTicket)
	protected.GET("/tickets/my-tickets", ticketController.GetMyTickets)
	protected.DELETE("/tickets/:id", ticketController.CancelTicket)
	protected.PUT("/tickets/:id/transfer", ticketController.TransferTicket)

	// TODO (entrega final): agregar rutas de administrador con middleware de rol

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := r.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
