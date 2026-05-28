package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestCreateRaffle_Success verifies that an admin can create a raffle for an active event.
func TestCreateRaffle_Success(t *testing.T) {
	// TODO: seed active event
	// TODO: call service.Create(RaffleInput{EventID, Name, PricePerChance})
	// TODO: assert no error and raffle status == pending
	assert.True(t, true, "placeholder — implement test body")
}

// TestBuyChances_NewEntry verifies that a user's first chance purchase creates a new entry.
func TestBuyChances_NewEntry(t *testing.T) {
	// TODO: seed pending raffle, seed user
	// TODO: call service.BuyChances(userID, raffleID, 3)
	// TODO: assert entry is created with chances == 3
	assert.True(t, true, "placeholder — implement test body")
}

// TestBuyChances_IncrementExisting verifies that a second purchase increments existing chances.
func TestBuyChances_IncrementExisting(t *testing.T) {
	// TODO: seed raffle entry with chances = 2
	// TODO: call service.BuyChances(userID, raffleID, 3)
	// TODO: assert entry.Chances == 5
	assert.True(t, true, "placeholder — implement test body")
}

// TestDraw_PicksWinner verifies that Draw selects a winner and updates raffle status to done.
func TestDraw_PicksWinner(t *testing.T) {
	// TODO: seed raffle with multiple entries
	// TODO: call service.Draw(raffleID)
	// TODO: assert no error, returned user is one of the participants, raffle status == done
	assert.True(t, true, "placeholder — implement test body")
}
