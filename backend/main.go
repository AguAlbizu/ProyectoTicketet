package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"ticketapp/clients"
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

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		log.Fatal("Error al conectar con la base de datos:", err)
	}

	// AutoMigrate crea/actualiza las tablas automáticamente
	if err := db.AutoMigrate(&domain.User{}, &domain.Event{}, &domain.Ticket{}, &domain.Sorteo{}, &domain.Chance{}); err != nil {
		log.Fatal("AutoMigrate falló:", err)
	}

	// Agregar claves foráneas manualmente (idempotente: ignora error si ya existen)
	db.Exec("ALTER TABLE tickets ADD CONSTRAINT fk_tickets_users FOREIGN KEY (id_users) REFERENCES users(id_users) ON DELETE RESTRICT ON UPDATE CASCADE")
	db.Exec("ALTER TABLE tickets ADD CONSTRAINT fk_tickets_events FOREIGN KEY (id_events) REFERENCES events(id_events) ON DELETE RESTRICT ON UPDATE CASCADE")
	db.Exec("ALTER TABLE sorteos ADD CONSTRAINT fk_sorteos_events FOREIGN KEY (id_events) REFERENCES events(id_events) ON DELETE CASCADE ON UPDATE CASCADE")
	db.Exec("ALTER TABLE sorteos ADD CONSTRAINT fk_sorteos_ganador FOREIGN KEY (id_ganador) REFERENCES users(id_users) ON DELETE SET NULL ON UPDATE CASCADE")
	db.Exec("ALTER TABLE chances ADD CONSTRAINT fk_chances_sorteos FOREIGN KEY (id_sorteo) REFERENCES sorteos(id_sorteo) ON DELETE CASCADE ON UPDATE CASCADE")
	db.Exec("ALTER TABLE chances ADD CONSTRAINT fk_chances_users FOREIGN KEY (id_users) REFERENCES users(id_users) ON DELETE RESTRICT ON UPDATE CASCADE")

	// Instanciar DAOs
	userDAO := dao.NewUserDAO(db)
	eventDAO := dao.NewEventDAO(db)
	ticketDAO := dao.NewTicketDAO(db)
	sorteoDAO := dao.NewSorteoDAO(db)
	chanceDAO := dao.NewChanceDAO(db)

	// Cliente de email: usa la API HTTP configurada, o un no-op si no hay EMAIL_API_URL (desarrollo local).
	var emailClient clients.EmailClient
	if os.Getenv("EMAIL_API_URL") != "" {
		emailClient = clients.NewEmailClient()
	} else {
		emailClient = &clients.NoOpEmailClient{}
		log.Println("EMAIL_API_URL no configurado: los emails de sorteo se omiten (NoOpEmailClient)")
	}

	// Instanciar services
	authService := services.NewAuthService(userDAO)
	eventService := services.NewEventService(eventDAO)
	ticketService := services.NewTicketService(ticketDAO, eventDAO, userDAO)
	sorteoService := services.NewSorteoService(sorteoDAO, chanceDAO, eventDAO, ticketDAO, userDAO, emailClient)

	// Instanciar controllers
	authController := controllers.NewAuthController(authService)
	eventController := controllers.NewEventController(eventService)
	ticketController := controllers.NewTicketController(ticketService)
	sorteoController := controllers.NewSorteoController(sorteoService)

	r := gin.Default()

	// CORS: permitir peticiones desde el frontend en localhost:5173 o 5174
	r.Use(func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		if origin == "http://localhost:5173" || origin == "http://localhost:5174" {
			c.Header("Access-Control-Allow-Origin", origin)
		}
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
	api.GET("/events/:id/sorteo", sorteoController.GetSorteoByEvent)

	// Rutas protegidas — requieren JWT válido
	protected := api.Group("")
	protected.Use(middleware.AuthMiddleware())
	protected.POST("/tickets", ticketController.BuyTicket)
	protected.GET("/tickets/my-tickets", ticketController.GetMyTickets)
	protected.DELETE("/tickets/:id", ticketController.CancelTicket)
	protected.PUT("/tickets/:id/transfer", ticketController.TransferTicket)
	protected.POST("/sorteos/:id/chances", sorteoController.BuyChances)
	protected.GET("/sorteos/:id/my-chances", sorteoController.GetMyChances)

	// Rutas de administrador — requieren JWT válido + rol "admin"
	admin := api.Group("/admin")
	admin.Use(middleware.AuthMiddleware(), middleware.AdminOnly())
	admin.POST("/events/:id/sorteo", sorteoController.CreateSorteo)
	admin.GET("/sorteos", sorteoController.ListSorteosAdmin)
	admin.POST("/sorteos/:id/draw", sorteoController.RunDraw)

	// TODO (entrega final): agregar el resto de rutas de administrador (CRUD de eventos, reportes)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Servidor corriendo en http://localhost:%s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal("Error al iniciar el servidor:", err)
	}
}
