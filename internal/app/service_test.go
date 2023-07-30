package app_test

import (
	"testing"

	"github.com/Nowak90210/ports/internal/app"
	"github.com/Nowak90210/ports/internal/domain"
	"github.com/Nowak90210/ports/internal/infrastructure/storage"
	"github.com/stretchr/testify/assert"
)

func TestSavePort(t *testing.T) {
	repo := storage.NewInMemoryRepo()
	servce := app.NewService(repo)

	id := "testID"
	port := domain.NewPort(id, "test name", "city", "province", "country", "timezone", "52000", nil, nil, nil, []float64{15.23, 123.634})
	updatedPort := domain.NewPort(id, "edited name", "different city", "other province", "not same country", "timezone", "52001", []string{"test", "test1"}, nil, nil, []float64{32, 12.4})

	// save should work and return no error
	err := servce.SavePort(port)
	assert.Nil(t, err)

	// port with ID "testID" should be equal to port variable
	p, err := servce.Get(id)
	assert.Nil(t, err)
	assert.Equal(t, port, p)

	// now let's update port, there should be no error
	err = servce.SavePort(updatedPort)
	assert.Nil(t, err)

	// port with ID "testID" should be equal to updatedPort variable
	p, err = servce.Get(id)
	assert.Nil(t, err)
	assert.Equal(t, updatedPort, p)
}
