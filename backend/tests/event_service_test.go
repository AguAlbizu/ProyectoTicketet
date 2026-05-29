package tests

// Objetivo de cobertura para la entrega parcial: >= 40% en servicios y controladores.
// Correr con: go test ./tests/... -cover

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestGetAllEvents_ReturnsActiveEvents
// Caso de éxito: debe retornar solo los eventos con estado "activo".
func TestGetAllEvents_ReturnsActiveEvents(t *testing.T) {
	// TODO: seedear eventos activos y cancelados en la DB de test
	// TODO: llamar service.GetAll("")
	// TODO: assert que solo se retornen eventos activos
	// TODO: assert que la cantidad sea la esperada
	assert.True(t, true, "placeholder — implementar test body")
}

// TestGetAllEvents_FilterByCategory
// Caso de éxito: filtrar por categoría debe retornar solo eventos de esa categoría.
func TestGetAllEvents_FilterByCategory(t *testing.T) {
	// TODO: seedear eventos de distintas categorías
	// TODO: llamar service.GetAll("concierto")
	// TODO: assert que todos los eventos retornados tengan categoria == "concierto"
	assert.True(t, true, "placeholder — implementar test body")
}

// TestGetEventByID_Success
// Caso de éxito: buscar un evento existente por ID debe retornarlo sin error.
func TestGetEventByID_Success(t *testing.T) {
	// TODO: seedear un evento
	// TODO: llamar service.GetByID(event.ID)
	// TODO: assert que no haya error y que el evento retornado tenga el ID correcto
	assert.True(t, true, "placeholder — implementar test body")
}

// TestGetEventByID_NotFound
// Caso de error: buscar un ID inexistente debe retornar error.
func TestGetEventByID_NotFound(t *testing.T) {
	// TODO: llamar service.GetByID con un ID que no existe (ej: 99999)
	// TODO: assert que se retorne un error
	assert.True(t, true, "placeholder — implementar test body")
}
