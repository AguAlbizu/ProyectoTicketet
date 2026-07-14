package services

import (
	"fmt"
	"math/rand"
	"time"

	"ticketapp/domain"
)

// SorteoDAOPort define los métodos de persistencia de sorteos requeridos por SorteoService.
type SorteoDAOPort interface {
	CreateSorteo(sorteo *domain.Sorteo) error
	GetSorteoByID(id uint) (*domain.Sorteo, error)
	GetSorteoByEventID(eventID uint) (*domain.Sorteo, error)
	GetActiveSorteoByEventID(eventID uint) (*domain.Sorteo, error)
	GetSorteosByEventID(eventID uint) ([]domain.Sorteo, error)
	UpdateSorteo(sorteo *domain.Sorteo) error
	GetSorteosConEvento() ([]domain.Sorteo, error)
}

// ChanceDAOPort define los métodos de persistencia de chances requeridos por SorteoService.
type ChanceDAOPort interface {
	CreateChance(chance *domain.Chance) error
	GetChancesBySorteoID(sorteoID uint) ([]domain.Chance, error)
	CountChancesByUserAndSorteo(userID, sorteoID uint) (int64, error)
}

// SorteoEventDAOPort define el método de eventos requerido por SorteoService.
type SorteoEventDAOPort interface {
	GetEventByID(id uint) (*domain.Event, error)
}

// SorteoTicketDAOPort verifica elegibilidad: solo puede comprar chances quien tiene
// una entrada activa para el evento del sorteo.
type SorteoTicketDAOPort interface {
	GetActiveTicketByUserAndEvent(userID, eventID uint) (*domain.Ticket, error)
}

// SorteoUserDAOPort define el método de usuarios requerido por SorteoService.
type SorteoUserDAOPort interface {
	GetUserByID(id uint) (*domain.User, error)
}

// SorteoNotificationDAOPort define el método de notificaciones requerido por SorteoService.
type SorteoNotificationDAOPort interface {
	CreateNotification(n *domain.Notification) error
}

type SorteoService struct {
	sorteoDAO       SorteoDAOPort
	chanceDAO       ChanceDAOPort
	eventDAO        SorteoEventDAOPort
	ticketDAO       SorteoTicketDAOPort
	userDAO         SorteoUserDAOPort
	notificationDAO SorteoNotificationDAOPort
}

func NewSorteoService(
	sorteoDAO SorteoDAOPort,
	chanceDAO ChanceDAOPort,
	eventDAO SorteoEventDAOPort,
	ticketDAO SorteoTicketDAOPort,
	userDAO SorteoUserDAOPort,
	notificationDAO SorteoNotificationDAOPort,
) *SorteoService {
	return &SorteoService{
		sorteoDAO:       sorteoDAO,
		chanceDAO:       chanceDAO,
		eventDAO:        eventDAO,
		ticketDAO:       ticketDAO,
		userDAO:         userDAO,
		notificationDAO: notificationDAO,
	}
}

// CreateSorteo crea un nuevo sorteo para un evento. Un evento admite un único sorteo ACTIVO
// a la vez, pero una vez realizado (o cancelado) se puede cargar otro: el historial completo
// queda disponible vía GetSorteosByEventID.
func (s *SorteoService) CreateSorteo(eventID uint, nombre string, valorChance int) (*domain.Sorteo, error) {
	if nombre == "" {
		return nil, fmt.Errorf("el nombre del sorteo es requerido")
	}
	if valorChance <= 0 {
		return nil, fmt.Errorf("el valor de la chance debe ser mayor a 0")
	}
	if _, err := s.eventDAO.GetEventByID(eventID); err != nil {
		return nil, fmt.Errorf("evento no encontrado")
	}
	if _, err := s.sorteoDAO.GetActiveSorteoByEventID(eventID); err == nil {
		return nil, fmt.Errorf("el evento ya tiene un sorteo activo")
	}

	sorteo := &domain.Sorteo{
		IDEvents:    eventID,
		Nombre:      nombre,
		ValorChance: valorChance,
		Estado:      "activo",
	}
	if err := s.sorteoDAO.CreateSorteo(sorteo); err != nil {
		return nil, fmt.Errorf("error al crear el sorteo: %w", err)
	}
	return sorteo, nil
}

// GetSorteoByEventID retorna el sorteo del evento, si tiene uno cargado.
func (s *SorteoService) GetSorteoByEventID(eventID uint) (*domain.Sorteo, error) {
	sorteo, err := s.sorteoDAO.GetSorteoByEventID(eventID)
	if err != nil {
		return nil, fmt.Errorf("este evento no tiene sorteo")
	}
	return sorteo, nil
}

// BuyChances registra "cantidad" chances para el usuario en el sorteo dado.
// Requiere que el usuario tenga una entrada activa para el evento del sorteo.
func (s *SorteoService) BuyChances(userID, sorteoID uint, cantidad int) ([]domain.Chance, error) {
	if cantidad < 1 {
		return nil, fmt.Errorf("la cantidad de chances debe ser al menos 1")
	}

	sorteo, err := s.sorteoDAO.GetSorteoByID(sorteoID)
	if err != nil {
		return nil, fmt.Errorf("sorteo no encontrado")
	}
	if sorteo.Estado != "activo" {
		return nil, fmt.Errorf("el sorteo ya no admite nuevas chances")
	}
	if _, err := s.ticketDAO.GetActiveTicketByUserAndEvent(userID, sorteo.IDEvents); err != nil {
		return nil, fmt.Errorf("necesitás una entrada activa para este evento para participar del sorteo")
	}

	chances := make([]domain.Chance, 0, cantidad)
	for i := 0; i < cantidad; i++ {
		chance := &domain.Chance{IDSorteo: sorteoID, IDUsers: userID, FechaCompra: time.Now()}
		if err := s.chanceDAO.CreateChance(chance); err != nil {
			return nil, fmt.Errorf("error al registrar la chance: %w", err)
		}
		chances = append(chances, *chance)
	}

	_ = s.notificationDAO.CreateNotification(&domain.Notification{
		IDUsers:  userID,
		Tipo:     "chance_comprada",
		Titulo:   "Compraste tu participación en el sorteo",
		Mensaje:  fmt.Sprintf("Compraste %d chance(s) para el sorteo \"%s\". ¡Mucha suerte!", cantidad, sorteo.Nombre),
		IDSorteo: &sorteo.IDSorteo,
	})

	return chances, nil
}

// GetMyChancesCount retorna cuántas chances tiene el usuario en un sorteo puntual.
func (s *SorteoService) GetMyChancesCount(userID, sorteoID uint) (int64, error) {
	return s.chanceDAO.CountChancesByUserAndSorteo(userID, sorteoID)
}

// GetSorteosConEvento lista todos los sorteos con su evento, para el panel admin.
func (s *SorteoService) GetSorteosConEvento() ([]domain.Sorteo, error) {
	return s.sorteoDAO.GetSorteosConEvento()
}

// GetSorteosByEventID retorna el historial de sorteos de un evento, más recientes primero.
func (s *SorteoService) GetSorteosByEventID(eventID uint) ([]domain.Sorteo, error) {
	return s.sorteoDAO.GetSorteosByEventID(eventID)
}

// ChanceSummary agrupa las chances de un sorteo por usuario participante.
type ChanceSummary struct {
	UserID   uint   `json:"user_id"`
	Nombre   string `json:"nombre"`
	Email    string `json:"email"`
	Cantidad int    `json:"cantidad"`
}

// GetChanceSummary retorna, para un sorteo, cada usuario participante con la cantidad
// de chances que compró. Se usa en el panel admin para ver quiénes están participando.
func (s *SorteoService) GetChanceSummary(sorteoID uint) ([]ChanceSummary, error) {
	chances, err := s.chanceDAO.GetChancesBySorteoID(sorteoID)
	if err != nil {
		return nil, fmt.Errorf("error al obtener las chances del sorteo")
	}

	order := make([]uint, 0)
	byUser := map[uint]*ChanceSummary{}
	for _, c := range chances {
		summary, ok := byUser[c.IDUsers]
		if !ok {
			summary = &ChanceSummary{UserID: c.IDUsers, Nombre: c.User.Nombre, Email: c.User.Email}
			byUser[c.IDUsers] = summary
			order = append(order, c.IDUsers)
		}
		summary.Cantidad++
	}

	result := make([]ChanceSummary, 0, len(order))
	for _, uid := range order {
		result = append(result, *byUser[uid])
	}
	return result, nil
}

// RunDraw selecciona un ganador al azar entre las chances cargadas, marca el sorteo como
// realizado y notifica por email a todos los participantes (ganador y perdedores).
func (s *SorteoService) RunDraw(sorteoID uint) (*domain.User, error) {
	sorteo, err := s.sorteoDAO.GetSorteoByID(sorteoID)
	if err != nil {
		return nil, fmt.Errorf("sorteo no encontrado")
	}
	if sorteo.Estado != "activo" {
		return nil, fmt.Errorf("el sorteo ya fue realizado")
	}

	chances, err := s.chanceDAO.GetChancesBySorteoID(sorteoID)
	if err != nil {
		return nil, fmt.Errorf("error al obtener las chances del sorteo: %w", err)
	}
	if len(chances) == 0 {
		return nil, fmt.Errorf("el sorteo no tiene participantes")
	}

	winnerChance := chances[rand.Intn(len(chances))]
	winner, err := s.userDAO.GetUserByID(winnerChance.IDUsers)
	if err != nil {
		return nil, fmt.Errorf("error al obtener el ganador")
	}

	now := time.Now()
	sorteo.Estado = "realizado"
	sorteo.IDGanador = &winner.IDUsers
	sorteo.FechaRealizado = &now
	if err := s.sorteoDAO.UpdateSorteo(sorteo); err != nil {
		return nil, fmt.Errorf("error al actualizar el sorteo: %w", err)
	}

	s.notifyParticipants(chances, winner.IDUsers, sorteo.IDSorteo, sorteo.Nombre)

	return winner, nil
}

// notifyParticipants crea una notificación in-app para cada participante único: felicitación
// al ganador, aviso de resultado a los demás. Los errores no interrumpen el sorteo (best-effort).
func (s *SorteoService) notifyParticipants(chances []domain.Chance, winnerID, sorteoID uint, sorteoNombre string) {
	notified := map[uint]bool{}
	for _, chance := range chances {
		if notified[chance.IDUsers] {
			continue
		}
		notified[chance.IDUsers] = true

		if chance.IDUsers == winnerID {
			_ = s.notificationDAO.CreateNotification(&domain.Notification{
				IDUsers:  chance.IDUsers,
				Tipo:     "sorteo_ganador",
				Titulo:   "¡Ganaste el sorteo!",
				Mensaje:  fmt.Sprintf("Felicitaciones, resultaste ganador/a del sorteo \"%s\".", sorteoNombre),
				IDSorteo: &sorteoID,
			})
		} else {
			_ = s.notificationDAO.CreateNotification(&domain.Notification{
				IDUsers:  chance.IDUsers,
				Tipo:     "sorteo_perdedor",
				Titulo:   "Resultado del sorteo",
				Mensaje:  fmt.Sprintf("El sorteo \"%s\" ya se realizó y en esta oportunidad no resultaste ganador/a. ¡Gracias por participar!", sorteoNombre),
				IDSorteo: &sorteoID,
			})
		}
	}
}
