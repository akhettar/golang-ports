package repository

import (
	m "ports-writer-service/model"
	"reflect"
	"testing"
)

func TestInMemoryRepository_AddPort(t *testing.T) {

	rep := NewInMemoryRepository()

	// Add port
	given := m.Port{PortCode: "AEAJM", Name: "Ajman"}
	rep.AddPort(given)

	// Query Port
	got, err := rep.GetPort("AEAJM")

	if err != nil {
		t.Errorf("error should not have been returned")
	}

	if !reflect.DeepEqual(given, got) {
		t.Errorf("InMemoryRepository.AddPort() got = %v, want %v", got, given)
	}
}
