package file_test

import (
	"testing"

	"github.com/Nowak90210/ports/internal/domain"
	"github.com/Nowak90210/ports/internal/infrastructure/storage/file"
	"github.com/stretchr/testify/assert"
)

func TestReadJsonFile(t *testing.T) {
	portsChan := make(chan *domain.Port)
	ports := make([]*domain.Port, 0)
	fileReader := file.NewJsonFileReader()

	go fileReader.ReadPortsFromFile("test_file.json", portsChan)

	func() {
		for port := range portsChan {
			ports = append(ports, port)
		}
	}()

	port1 := domain.NewPort("AEAJM", "Ajman", "Ajman", "Ajman", "United Arab Emirates", "Asia/Dubai", "52000", []string{}, []string{}, []string{"AEAJM"}, []float64{55.5136433, 25.4052165})
	port2 := domain.NewPort("AEAUH", "Abu Dhabi", "Abu Dhabi", "Abu Z¸aby [Abu Dhabi]", "United Arab Emirates", "Asia/Dubai", "52001", []string{}, []string{}, []string{"AEAUH"}, []float64{54.37, 24.47})

	assert.Equal(t, 2, len(ports))
	assert.Equal(t, port1, ports[0])
	assert.Equal(t, port2, ports[1])
}
