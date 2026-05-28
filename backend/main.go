package main

import (
	"log"
	"os"
	"ticketapp/clients"
	"ticketapp/controllers"
	"ticketapp/dao"
	"ticketapp/domain"
	"ticketapp/services"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// Load .env file (ignore error in production where env vars are injected directly)
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	// Connect to database
	// TODO: build DSN string from DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME env vars
	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", ...)
	db, err := gorm.Open(mysql.Open(""), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto-migrate domain models
	// TODO: uncomment after DSN is configured
	_ = db.AutoMigrate(
		&domain.User{},
		&domain.Event{},
		&domain.Ticket{},
		&domain.Raffle{},
		&domain.RaffleEntry{},
	)

	// Instantiate DAOs
	userDAO := dao.NewUserDAO(db)
	eventDAO := dao.NewEventDAO(db)
	ticketDAO := dao.NewTicketDAO(db)
	raffleDAO := dao.NewRaffleDAO(db)

	// Instantiate services
	authService := services.NewAuthService(userDAO)
	eventService := services.NewEventService(eventDAO)
	ticketService := services.NewTicketService(ticketDAO, eventDAO, userDAO)
	emailClient := clients.NewEmailClient()
	raffleService := services.NewRaffleService(raffleDAO, eventDAO, userDAO, emailClient)

	// Instantiate controllers
	authController := controllers.NewAuthController(authService)
	eventController := controllers.NewEventController(eventService)
	ticketController := controllers.NewTicketController(ticketService)
	raffleController := controllers.NewRaffleController(raffleService)

	// Initialize Gin router
	r := gin.Default()

	// TODO: add CORS middleware

	api := r.Group("/api")

	// TODO: create JWT middleware (authMiddleware) using utils.ValidateToken
	// TODO: create admin middleware (adminMiddleware) that checks role == "admin"
	var authMiddleware gin.HandlerFunc  // TODO: replace with actual middleware
	var adminMiddleware gin.HandlerFunc // TODO: replace with actual middleware

	// Register routes
	authController.RegisterRoutes(api)
	eventController.RegisterRoutes(api, authMiddleware, adminMiddleware)
	ticketController.RegisterRoutes(api, authMiddleware)
	raffleController.RegisterRoutes(api, authMiddleware, adminMiddleware)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := r.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
