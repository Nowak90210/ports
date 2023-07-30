package storage

import (
	"fmt"

	"github.com/Nowak90210/ports/internal/domain"
)

type InMemoryRepo struct {
	ports map[string]domain.Port
}

func NewInMemoryRepo() *InMemoryRepo {
	ports := make(map[string]domain.Port)

	return &InMemoryRepo{ports: ports}
}

func (r *InMemoryRepo) Get(id string) (*domain.Port, error) {
	p, ok := r.ports[id]
	if !ok {
		return nil, fmt.Errorf("port id: `%s` not found", id)
	}

	return &p, nil
}

func (r *InMemoryRepo) Upsert(port *domain.Port) error {
	r.ports[port.ID] = *port

	return nil
}
