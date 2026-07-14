package services

import (
	"fmt"
	"ticketapp/domain"
	"time"
)

type AdminEventRepository interface {
	GetAllEventsAdmin() ([]domain.Event, error)
	GetEventByID(id uint) (*domain.Event, error)
	CreateEvent(event *domain.Event) error
	FullUpdateEvent(event *domain.Event) error
}

type AdminTicketRepository interface {
	GetTicketsByEventID(eventID uint) ([]domain.Ticket, error)
	CancelAllTicketsByEventID(eventID uint) error
}

type AdminEventService struct {
	eventDAO  AdminEventRepository
	ticketDAO AdminTicketRepository
}

func NewAdminEventService(eventDAO AdminEventRepository, ticketDAO AdminTicketRepository) *AdminEventService {
	return &AdminEventService{eventDAO: eventDAO, ticketDAO: ticketDAO}
}

type CreateEventInput struct {
	Titulo      string    `json:"titulo" binding:"required"`
	Descripcion string    `json:"descripcion"`
	Fecha       time.Time `json:"fecha" binding:"required"`
	Hora        string    `json:"hora" binding:"required"`
	Capacidad   int       `json:"capacidad" binding:"required,min=1"`
	Categoria   string    `json:"categoria"`
	Direccion   string    `json:"direccion"`
	ImagenURL   string    `json:"imagen_url"`
	Precio      int       `json:"precio"`
}

type UpdateEventInput struct {
	Titulo      string    `json:"titulo" binding:"required"`
	Descripcion string    `json:"descripcion"`
	Fecha       time.Time `json:"fecha" binding:"required"`
	Hora        string    `json:"hora" binding:"required"`
	Capacidad   int       `json:"capacidad" binding:"required,min=1"`
	Categoria   string    `json:"categoria"`
	Direccion   string    `json:"direccion"`
	ImagenURL   string    `json:"imagen_url"`
	Precio      int       `json:"precio"`
	Estado      string    `json:"estado"`
}

type BuyerInfo struct {
	UserID      uint      `json:"user_id"`
	Nombre      string    `json:"nombre"`
	Email       string    `json:"email"`
	TicketID    uint      `json:"ticket_id"`
	FechaCompra time.Time `json:"fecha_compra"`
	Estado      string    `json:"estado"`
	Origen      string    `json:"origen"`
}

type EventReport struct {
	EventID             uint        `json:"event_id"`
	Titulo              string      `json:"titulo"`
	Capacidad           int         `json:"capacidad"`
	CupoDisponible      int         `json:"cupo_disponible"`
	EntradasVendidas    int         `json:"entradas_vendidas"`
	PorcentajeOcupacion float64     `json:"porcentaje_ocupacion"`
	// Compradores: todos los que compraron originalmente una entrada (origen "compra"),
	// en cualquier estado — si la cancelaron o transfirieron, el campo Estado lo indica.
	Compradores []BuyerInfo `json:"compradores"`
	// TitularesActivos: quienes hoy tienen una entrada activa, ya sea porque la compraron
	// o porque se la transfirieron.
	TitularesActivos []BuyerInfo `json:"titulares_activos"`
}

func (s *AdminEventService) GetAllEvents() ([]domain.Event, error) {
	return s.eventDAO.GetAllEventsAdmin()
}

func (s *AdminEventService) CreateEvent(input CreateEventInput) (*domain.Event, error) {
	event := &domain.Event{
		Titulo:         input.Titulo,
		Descripcion:    input.Descripcion,
		Fecha:          input.Fecha,
		Hora:           input.Hora,
		Capacidad:      input.Capacidad,
		CupoDisponible: input.Capacidad,
		Categoria:      input.Categoria,
		Direccion:      input.Direccion,
		ImagenURL:      input.ImagenURL,
		Precio:         input.Precio,
		Estado:         "activo",
	}
	if err := s.eventDAO.CreateEvent(event); err != nil {
		return nil, fmt.Errorf("error al crear el evento")
	}
	return event, nil
}

func (s *AdminEventService) UpdateEvent(id uint, input UpdateEventInput) (*domain.Event, error) {
	event, err := s.eventDAO.GetEventByID(id)
	if err != nil {
		return nil, fmt.Errorf("evento no encontrado")
	}

	// Determinar nuevo estado (si no se envía, mantener el actual)
	newEstado := event.Estado
	if input.Estado == "activo" || input.Estado == "cancelado" {
		newEstado = input.Estado
	}

	var newCupo int
	switch {
	case event.Estado == "activo" && newEstado == "cancelado":
		// Desactivar: cancelar todas las entradas activas
		if err := s.ticketDAO.CancelAllTicketsByEventID(id); err != nil {
			return nil, fmt.Errorf("error al cancelar las entradas del evento")
		}
		newCupo = input.Capacidad
	case event.Estado == "cancelado" && newEstado == "activo":
		// Reactivar: los tickets ya estaban cancelados, arrancar con cupo completo
		newCupo = input.Capacidad
	default:
		// Sin cambio de estado: recalcular cupo según capacidad nueva
		ticketsSold := event.Capacidad - event.CupoDisponible
		newCupo = input.Capacidad - ticketsSold
		if newCupo < 0 {
			return nil, fmt.Errorf("la nueva capacidad (%d) es menor a las entradas ya vendidas (%d)", input.Capacidad, ticketsSold)
		}
	}

	event.Titulo = input.Titulo
	event.Descripcion = input.Descripcion
	event.Fecha = input.Fecha
	event.Hora = input.Hora
	event.Capacidad = input.Capacidad
	event.CupoDisponible = newCupo
	event.Categoria = input.Categoria
	event.Direccion = input.Direccion
	event.ImagenURL = input.ImagenURL
	event.Precio = input.Precio
	event.Estado = newEstado

	if err := s.eventDAO.FullUpdateEvent(event); err != nil {
		return nil, fmt.Errorf("error al actualizar el evento")
	}
	return event, nil
}

func (s *AdminEventService) CancelEvent(id uint) error {
	event, err := s.eventDAO.GetEventByID(id)
	if err != nil {
		return fmt.Errorf("evento no encontrado")
	}
	if event.Estado == "cancelado" {
		return fmt.Errorf("el evento ya está cancelado")
	}

	if err := s.ticketDAO.CancelAllTicketsByEventID(id); err != nil {
		return fmt.Errorf("error al cancelar las entradas del evento")
	}

	event.Estado = "cancelado"
	return s.eventDAO.FullUpdateEvent(event)
}

func (s *AdminEventService) GetEventReport(id uint) (*EventReport, error) {
	event, err := s.eventDAO.GetEventByID(id)
	if err != nil {
		return nil, fmt.Errorf("evento no encontrado")
	}

	tickets, err := s.ticketDAO.GetTicketsByEventID(id)
	if err != nil {
		return nil, fmt.Errorf("error al obtener las entradas")
	}

	compradores := make([]BuyerInfo, 0, len(tickets))
	activos := make([]BuyerInfo, 0, len(tickets))
	for _, t := range tickets {
		info := BuyerInfo{
			UserID:      t.IDUsers,
			Nombre:      t.User.Nombre,
			Email:       t.User.Email,
			TicketID:    t.IDTickets,
			FechaCompra: t.FechaCompra,
			Estado:      t.Estado,
			Origen:      t.Origen,
		}
		if t.Origen == "compra" || t.Origen == "" {
			compradores = append(compradores, info)
		}
		if t.Estado == "activo" {
			activos = append(activos, info)
		}
	}

	entradasVendidas := event.Capacidad - event.CupoDisponible
	var porcentaje float64
	if event.Capacidad > 0 {
		porcentaje = float64(entradasVendidas) / float64(event.Capacidad) * 100
	}

	return &EventReport{
		EventID:             event.IDEvents,
		Titulo:              event.Titulo,
		Capacidad:           event.Capacidad,
		CupoDisponible:      event.CupoDisponible,
		EntradasVendidas:    entradasVendidas,
		PorcentajeOcupacion: porcentaje,
		Compradores:         compradores,
		TitularesActivos:    activos,
	}, nil
}
