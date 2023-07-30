package file

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"github.com/Nowak90210/ports/internal/domain"
	"github.com/Nowak90210/ports/internal/infrastructure"
)

type JsonFileReader struct{}

func NewJsonFileReader() *JsonFileReader {
	return &JsonFileReader{}
}

func (r *JsonFileReader) ReadPortsFromFile(fileName string, stream chan (*domain.Port)) error {
	defer close(stream)

	file, err := os.Open(fileName)
	if err != nil {
		return fmt.Errorf("cannot open file: '%w'", err)
	}

	fileReader := bufio.NewReader(file)
	decoder := json.NewDecoder(fileReader)

	// strip first curly braces
	_, err = decoder.Token()
	if err != nil {
		return fmt.Errorf("decode first token: %w", err)
	}

	for decoder.More() {
		token, err := decoder.Token()
		if err != nil {
			return fmt.Errorf("decode second token: %w", err)
		}

		var dto infrastructure.PortDto
		err = decoder.Decode(&dto)
		if err != nil {
			return fmt.Errorf("decode dto: %w", err)
		}

		dto.ID = fmt.Sprint(token)
		stream <- dto.ToDomain()
	}

	return nil
}
