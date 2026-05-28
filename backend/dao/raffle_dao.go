package dao

import (
	"ticketapp/domain"

	"gorm.io/gorm"
)

// RaffleDAO handles all database operations for Raffle and RaffleEntry models.
type RaffleDAO struct {
	db *gorm.DB
}

// NewRaffleDAO creates a new RaffleDAO with the provided GORM instance.
func NewRaffleDAO(db *gorm.DB) *RaffleDAO {
	return &RaffleDAO{db: db}
}

// CreateRaffle persists a new raffle record.
func (d *RaffleDAO) CreateRaffle(raffle *domain.Raffle) error {
	// TODO: d.db.Create(raffle)
	return nil
}

// FindRaffleByEventID returns the raffle associated with a given event.
func (d *RaffleDAO) FindRaffleByEventID(eventID uint) (*domain.Raffle, error) {
	// TODO: d.db.Where("event_id = ?", eventID).First(&raffle)
	return nil, nil
}

// FindRaffleByID returns a single raffle by its ID.
func (d *RaffleDAO) FindRaffleByID(id uint) (*domain.Raffle, error) {
	// TODO: d.db.Preload("WinnerUser").First(&raffle, id)
	return nil, nil
}

// UpdateRaffle saves changes to an existing raffle (e.g. set winner, change status).
func (d *RaffleDAO) UpdateRaffle(raffle *domain.Raffle) error {
	// TODO: d.db.Save(raffle)
	return nil
}

// CreateEntry persists a new raffle entry (chance purchase).
func (d *RaffleDAO) CreateEntry(entry *domain.RaffleEntry) error {
	// TODO: d.db.Create(entry)
	return nil
}

// FindEntriesByRaffleID returns all entries for a given raffle.
func (d *RaffleDAO) FindEntriesByRaffleID(raffleID uint) ([]domain.RaffleEntry, error) {
	// TODO: d.db.Preload("User").Where("raffle_id = ?", raffleID).Find(&entries)
	return nil, nil
}

// FindEntryByUserAndRaffle returns the entry of a specific user in a specific raffle.
func (d *RaffleDAO) FindEntryByUserAndRaffle(userID, raffleID uint) (*domain.RaffleEntry, error) {
	// TODO: d.db.Where("user_id = ? AND raffle_id = ?", userID, raffleID).First(&entry)
	return nil, nil
}
