package repository

import (
	"log"
	m "ports-manager-service/model"
)

type Repository interface {
	AddPort(p m.Port) error
	GetPort(code string) (*m.Port, error)
}

// SQLite Repository
type SQLiteRepository struct {
	// some DB instnace
}

// InMemoryRepository Repository
type InMemoryRepository struct {
	DB map[string]*m.Port
}

// NewSQLiteRepository return a sequence of repository
func NewInMemoryRepository() Repository {
	return InMemoryRepository{DB: map[string]*m.Port{}}
}

func (r InMemoryRepository) AddPort(p m.Port) error {
	log.Printf("Creating port with code %s", p.PortCode)
	r.DB[p.PortCode] = &p
	return nil
}

func (r InMemoryRepository) GetPort(code string) (*m.Port, error) {
	log.Printf("querying port for given code %s\n", code)
	found := r.DB[code]
	return found, nil
}
