package main

import (
	"log"
	"net/http"

	"github.com/Nowak90210/ports/internal/app"
	"github.com/Nowak90210/ports/internal/infrastructure/file"
	transport "github.com/Nowak90210/ports/internal/infrastructure/http"
	"github.com/Nowak90210/ports/internal/infrastructure/storage"
)

func main() {
	repo := storage.NewInMemoryRepo()
	fileReader := file.NewJsonFileReader("json_files/")
	service := app.NewService(repo, fileReader)

	handler := transport.NewHandler(service)
	router := handler.GetRouter()

	log.Println("JSON API server running on port: ", "8080")

	http.ListenAndServe(":8080", router)
}
