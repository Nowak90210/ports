package app

import "github.com/Nowak90210/ports/internal/domain"

type PortRepository interface {
	Get(string) (*domain.Port, error)
	Upsert(*domain.Port) error
}
