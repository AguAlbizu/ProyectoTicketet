package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestCreateEvent_Success verifies that an admin can create a valid event.
func TestCreateEvent_Success(t *testing.T) {
	// TODO: set up mock or in-memory EventDAO
	// TODO: build valid EventInput (title, date, capacity > 0)
	// TODO: call service.Create(input)
	// TODO: assert no error and event has AvailableTickets == Capacity
	assert.True(t, true, "placeholder — implement test body")
}

// TestCreateEvent_InvalidCapacity verifies that capacity <= 0 returns an error.
func TestCreateEvent_InvalidCapacity(t *testing.T) {
	// TODO: call service.Create with capacity = 0
	// TODO: assert validation error is returned
	assert.True(t, true, "placeholder — implement test body")
}

// TestGetAllEvents verifies that active events are returned and filtered by category.
func TestGetAllEvents(t *testing.T) {
	// TODO: seed events of different categories
	// TODO: call service.GetAll("concierto")
	// TODO: assert only events with matching category are returned
	assert.True(t, true, "placeholder — implement test body")
}

// TestCancelEvent verifies that cancelling an event changes its status.
func TestCancelEvent(t *testing.T) {
	// TODO: seed an active event
	// TODO: call service.Cancel(event.ID)
	// TODO: fetch event again and assert status == cancelled
	assert.True(t, true, "placeholder — implement test body")
}
