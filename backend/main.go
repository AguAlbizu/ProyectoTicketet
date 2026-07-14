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

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		log.Fatal("Error al conectar con la base de datos:", err)
	}

	// Quitar el índice único anterior en sorteos.id_events: ahora un evento puede tener
	// varios sorteos en su historial (uno solo activo a la vez, validado en el service).
	// MySQL no permite borrar un índice que sostiene una FK, así que primero se quita la FK
	// (se vuelve a crear más abajo). Idempotente: si no existen, los errores se ignoran.
	db.Exec("ALTER TABLE sorteos DROP FOREIGN KEY fk_sorteos_events")
	db.Exec("ALTER TABLE sorteos DROP INDEX idx_sorteos_id_events")

	// AutoMigrate crea/actualiza las tablas automáticamente
	if err := db.AutoMigrate(&domain.User{}, &domain.Event{}, &domain.Ticket{}, &domain.Sorteo{}, &domain.Chance{}, &domain.Notification{}); err != nil {
		log.Fatal("AutoMigrate falló:", err)
	}

	// Agregar claves foráneas manualmente (idempotente: ignora error si ya existen)
	db.Exec("ALTER TABLE tickets ADD CONSTRAINT fk_tickets_users FOREIGN KEY (id_users) REFERENCES users(id_users) ON DELETE RESTRICT ON UPDATE CASCADE")
	db.Exec("ALTER TABLE tickets ADD CONSTRAINT fk_tickets_events FOREIGN KEY (id_events) REFERENCES events(id_events) ON DELETE RESTRICT ON UPDATE CASCADE")
	db.Exec("ALTER TABLE sorteos ADD CONSTRAINT fk_sorteos_events FOREIGN KEY (id_events) REFERENCES events(id_events) ON DELETE CASCADE ON UPDATE CASCADE")
	db.Exec("ALTER TABLE sorteos ADD CONSTRAINT fk_sorteos_ganador FOREIGN KEY (id_ganador) REFERENCES users(id_users) ON DELETE SET NULL ON UPDATE CASCADE")
	db.Exec("ALTER TABLE chances ADD CONSTRAINT fk_chances_sorteos FOREIGN KEY (id_sorteo) REFERENCES sorteos(id_sorteo) ON DELETE CASCADE ON UPDATE CASCADE")
	db.Exec("ALTER TABLE chances ADD CONSTRAINT fk_chances_users FOREIGN KEY (id_users) REFERENCES users(id_users) ON DELETE RESTRICT ON UPDATE CASCADE")
	db.Exec("ALTER TABLE notifications ADD CONSTRAINT fk_notifications_users FOREIGN KEY (id_users) REFERENCES users(id_users) ON DELETE CASCADE ON UPDATE CASCADE")
	db.Exec("ALTER TABLE notifications ADD CONSTRAINT fk_notifications_sorteos FOREIGN KEY (id_sorteo) REFERENCES sorteos(id_sorteo) ON DELETE CASCADE ON UPDATE CASCADE")

	// Instanciar DAOs
	userDAO := dao.NewUserDAO(db)
	eventDAO := dao.NewEventDAO(db)
	ticketDAO := dao.NewTicketDAO(db)
	sorteoDAO := dao.NewSorteoDAO(db)
	chanceDAO := dao.NewChanceDAO(db)
	notificationDAO := dao.NewNotificationDAO(db)

	// Instanciar services
	authService := services.NewAuthService(userDAO)
	eventService := services.NewEventService(eventDAO)
	ticketService := services.NewTicketService(ticketDAO, eventDAO, userDAO)
	sorteoService := services.NewSorteoService(sorteoDAO, chanceDAO, eventDAO, ticketDAO, userDAO, notificationDAO)
	adminEventService := services.NewAdminEventService(eventDAO, ticketDAO)
	notificationService := services.NewNotificationService(notificationDAO)

	// Instanciar controllers
	authController := controllers.NewAuthController(authService)
	eventController := controllers.NewEventController(eventService)
	ticketController := controllers.NewTicketController(ticketService)
	sorteoController := controllers.NewSorteoController(sorteoService)
	adminController := controllers.NewAdminController(adminEventService, authService)
	notificationController := controllers.NewNotificationController(notificationService)

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
	protected.GET("/notifications", notificationController.GetMyNotifications)
	protected.PUT("/notifications/:id/read", notificationController.MarkAsRead)
	protected.PUT("/notifications/read-all", notificationController.MarkAllAsRead)

	// Rutas de administrador — requieren JWT válido + rol administrador
	admin := api.Group("/admin")
	admin.Use(middleware.AuthMiddleware(), middleware.RequireRole("administrador"))
	admin.GET("/events", adminController.GetAllEvents)
	admin.POST("/events", adminController.CreateEvent)
	admin.PUT("/events/:id", adminController.UpdateEvent)
	admin.DELETE("/events/:id", adminController.CancelEvent)
	admin.GET("/events/:id/report", adminController.GetEventReport)
	admin.POST("/users", adminController.CreateAdmin)
	admin.PUT("/users/promote", adminController.PromoteToAdmin)
	admin.POST("/events/:id/sorteo", sorteoController.CreateSorteo)
	admin.GET("/events/:id/sorteos", sorteoController.GetSorteosByEvent)
	admin.GET("/sorteos", sorteoController.ListSorteosAdmin)
	admin.POST("/sorteos/:id/draw", sorteoController.RunDraw)
	admin.GET("/sorteos/:id/chances", sorteoController.GetChanceSummary)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Servidor corriendo en http://localhost:%s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal("Error al iniciar el servidor:", err)
	}
}
