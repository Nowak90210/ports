package app

import "github.com/Nowak90210/ports/internal/domain"

type FileReader interface {
	ReadPortsFromFile(string, chan *domain.Port) error
}
