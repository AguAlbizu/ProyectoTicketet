package services

import (
	"fmt"
	"math/rand"
	"time"

	"ticketapp/clients"
	"ticketapp/domain"
)

// SorteoDAOPort define los métodos de persistencia de sorteos requeridos por SorteoService.
type SorteoDAOPort interface {
	CreateSorteo(sorteo *domain.Sorteo) error
	GetSorteoByID(id uint) (*domain.Sorteo, error)
	GetSorteoByEventID(eventID uint) (*domain.Sorteo, error)
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

type SorteoService struct {
	sorteoDAO   SorteoDAOPort
	chanceDAO   ChanceDAOPort
	eventDAO    SorteoEventDAOPort
	ticketDAO   SorteoTicketDAOPort
	userDAO     SorteoUserDAOPort
	emailClient clients.EmailClient
}

func NewSorteoService(
	sorteoDAO SorteoDAOPort,
	chanceDAO ChanceDAOPort,
	eventDAO SorteoEventDAOPort,
	ticketDAO SorteoTicketDAOPort,
	userDAO SorteoUserDAOPort,
	emailClient clients.EmailClient,
) *SorteoService {
	return &SorteoService{
		sorteoDAO:   sorteoDAO,
		chanceDAO:   chanceDAO,
		eventDAO:    eventDAO,
		ticketDAO:   ticketDAO,
		userDAO:     userDAO,
		emailClient: emailClient,
	}
}

// CreateSorteo crea el sorteo para un evento. Un evento admite un único sorteo activo.
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
	if _, err := s.sorteoDAO.GetSorteoByEventID(eventID); err == nil {
		return nil, fmt.Errorf("el evento ya tiene un sorteo cargado")
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

	s.notifyParticipants(chances, winner.IDUsers, sorteo.Nombre)

	return winner, nil
}

// notifyParticipants envía un email a cada participante único: felicitación al ganador,
// aviso de resultado a los demás. Los errores de envío no interrumpen el sorteo (best-effort).
func (s *SorteoService) notifyParticipants(chances []domain.Chance, winnerID uint, sorteoNombre string) {
	notified := map[uint]bool{}
	for _, chance := range chances {
		if notified[chance.IDUsers] {
			continue
		}
		notified[chance.IDUsers] = true

		participant, err := s.userDAO.GetUserByID(chance.IDUsers)
		if err != nil || participant.Email == "" {
			continue
		}

		if chance.IDUsers == winnerID {
			_ = s.emailClient.SendEmail(participant.Email, "¡Ganaste el sorteo!",
				fmt.Sprintf("Felicitaciones %s, resultaste ganador/a del sorteo \"%s\".", participant.Nombre, sorteoNombre))
		} else {
			_ = s.emailClient.SendEmail(participant.Email, "Resultado del sorteo",
				fmt.Sprintf("Hola %s, el sorteo \"%s\" ya se realizó y en esta oportunidad no resultaste ganador/a. ¡Gracias por participar!", participant.Nombre, sorteoNombre))
		}
	}
}
