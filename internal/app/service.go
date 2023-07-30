package app

import (
	"fmt"

	"github.com/Nowak90210/ports/internal/domain"
)

type Service struct {
	repo       PortRepository
	fileReader FileReader
}

func NewService(repo PortRepository, fileReader FileReader) *Service {
	return &Service{
		repo:       repo,
		fileReader: fileReader,
	}
}

func (s *Service) Get(id string) (*domain.Port, error) {
	return s.repo.Get(id)
}

func (s *Service) SavePortsFromFile(fileName string) (int, error) {
	portsChan := make(chan *domain.Port)
	errs := make(chan error)
	counterChan := make(chan int)

	go func() {
		err := s.fileReader.ReadPortsFromFile(fileName, portsChan)
		if err != nil {
			errs <- err
		}
	}()

	go func() {
		counter, err := s.savePortsInBatch(portsChan)
		if err != nil {
			errs <- fmt.Errorf("SavePortsInBatch, err = %w", err)
		}
		counterChan <- counter
	}()

	select {
	case err := <-errs:
		return 0, err
	case c := <-counterChan:
		return c, nil
	}
}

func (s *Service) savePortsInBatch(portsChan <-chan *domain.Port) (int, error) {
	counter := 0
	for port := range portsChan {
		err := s.SavePort(port)
		if err != nil {
			return 0, fmt.Errorf("save port, err = %w", err)
		}
		counter++
	}

	return counter, nil
}

func (s *Service) SavePort(port *domain.Port) error {
	err := s.repo.Upsert(port)
	if err != nil {
		return fmt.Errorf("upsert port id `%s`: %w", port.ID, err)
	}

	return nil
}
