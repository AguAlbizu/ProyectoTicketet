package tests

// Objetivo de cobertura para la entrega parcial: >= 40% en servicios y controladores.
// Correr con: go test ./tests/... -v -cover

import (
	"testing"
	"ticketapp/domain"
	"ticketapp/services"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

// mockEventRepository implementa services.EventRepository para tests sin base de datos.
type mockEventRepository struct {
	events  []domain.Event
	findErr error
}

func (m *mockEventRepository) GetAllEvents(categoria string) ([]domain.Event, error) {
	if categoria == "" {
		return m.events, m.findErr
	}
	var filtered []domain.Event
	for _, e := range m.events {
		if e.Categoria == categoria {
			filtered = append(filtered, e)
		}
	}
	return filtered, m.findErr
}

func (m *mockEventRepository) GetEventByID(id uint) (*domain.Event, error) {
	if m.findErr != nil {
		return nil, m.findErr
	}
	for _, e := range m.events {
		if e.IDEvents == id {
			return &e, nil
		}
	}
	return nil, gorm.ErrRecordNotFound
}

// TestGetEventByID_NotFound verifica que buscar un ID inexistente retorna error.
func TestGetEventByID_NotFound(t *testing.T) {
	mockDAO := &mockEventRepository{findErr: gorm.ErrRecordNotFound}
	svc := services.NewEventService(mockDAO)

	_, err := svc.GetEventByID(9999)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "no encontrado")
}

// TestGetEvents_WithFilter verifica que el filtro por categoría retorna solo los eventos correctos.
func TestGetEvents_WithFilter(t *testing.T) {
	mockDAO := &mockEventRepository{
		events: []domain.Event{
			{IDEvents: 1, Titulo: "Concierto A", Categoria: "Música", Estado: "activo"},
			{IDEvents: 2, Titulo: "Obra de teatro", Categoria: "Teatro", Estado: "activo"},
			{IDEvents: 3, Titulo: "Concierto B", Categoria: "Música", Estado: "activo"},
		},
	}
	svc := services.NewEventService(mockDAO)

	eventos, err := svc.GetEvents("Música")
	assert.NoError(t, err)
	assert.Len(t, eventos, 2)
	for _, e := range eventos {
		assert.Equal(t, "Música", e.Categoria)
	}
}
