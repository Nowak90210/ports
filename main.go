package main

import (
	"fmt"

	"github.com/Nowak90210/ports/internal/app"
	"github.com/Nowak90210/ports/internal/infrastructure/file"
	"github.com/Nowak90210/ports/internal/infrastructure/storage"
)

func main() {
	repo := storage.NewInMemoryRepo()
	fileReader := file.NewJsonFileReader("json_files/")
	service := app.NewService(repo, fileReader)

	counter, err := service.SavePortsFromFile("ports.json")
	if err != nil {
		fmt.Println("err:", err)
	}

	fmt.Printf("Saved %d rows", counter)
	port, err := service.Get("ZWUTA")
	fmt.Printf("port: %+v", port)
}
