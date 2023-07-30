package app

import (
	"fmt"

	"github.com/Nowak90210/ports/internal/domain"
)

type Service struct {
	repo PortRepository
}

func NewService(repo PortRepository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) Get(id string) (*domain.Port, error) {
	return s.repo.Get(id)
}

func (s *Service) SavePort(port *domain.Port) error {
	err := s.repo.Upsert(port)
	if err != nil {
		return fmt.Errorf("upsert port id `%s`, error: %w", port.ID, err)
	}

	return nil
}
