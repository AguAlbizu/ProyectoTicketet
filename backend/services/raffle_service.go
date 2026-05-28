package services

import (
	"ticketapp/clients"
	"ticketapp/dao"
	"ticketapp/domain"
)

// RaffleService handles business logic for raffle creation, chance purchasing, and drawing.
type RaffleService struct {
	raffleDAO   *dao.RaffleDAO
	eventDAO    *dao.EventDAO
	userDAO     *dao.UserDAO
	emailClient *clients.EmailClient
}

// NewRaffleService creates a new RaffleService with its required dependencies.
func NewRaffleService(
	raffleDAO *dao.RaffleDAO,
	eventDAO *dao.EventDAO,
	userDAO *dao.UserDAO,
	emailClient *clients.EmailClient,
) *RaffleService {
	return &RaffleService{
		raffleDAO:   raffleDAO,
		eventDAO:    eventDAO,
		userDAO:     userDAO,
		emailClient: emailClient,
	}
}

// RaffleInput holds the data required to create a new raffle.
type RaffleInput struct {
	EventID        uint
	Name           string
	PricePerChance float64
}

// Create creates a new raffle for an existing event (admin only).
func (s *RaffleService) Create(input RaffleInput) (*domain.Raffle, error) {
	// TODO: verify event exists and is active via eventDAO.FindByID
	// TODO: build domain.Raffle and call raffleDAO.CreateRaffle
	return nil, nil
}

// GetByEvent returns the raffle associated with a given event.
func (s *RaffleService) GetByEvent(eventID uint) (*domain.Raffle, error) {
	// TODO: delegate to raffleDAO.FindRaffleByEventID(eventID)
	return nil, nil
}

// BuyChances adds or increments a user's entry in a raffle.
func (s *RaffleService) BuyChances(userID, raffleID uint, quantity int) (*domain.RaffleEntry, error) {
	// TODO: verify raffle is pending via raffleDAO.FindRaffleByID
	// TODO: check if entry already exists via raffleDAO.FindEntryByUserAndRaffle
	// TODO: if exists, increment chances; if not, create new entry
	// TODO: persist via raffleDAO.CreateEntry or raffleDAO.UpdateEntry
	return nil, nil
}

// Draw executes the raffle: picks a random winner from all entries (weighted by chances).
// Then sends winner/loser notification emails to all participants.
func (s *RaffleService) Draw(raffleID uint) (*domain.User, error) {
	// TODO: verify raffle is pending, fetch all entries
	// TODO: build weighted pool (repeat userID by chances count)
	// TODO: pick random winner from pool
	// TODO: set raffle.WinnerUserID = winner, Status = done, call raffleDAO.UpdateRaffle
	// TODO: iterate all entries: send winner email to winner, loser email to others
	return nil, nil
}
