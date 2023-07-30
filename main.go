package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Nowak90210/ports/internal/app"
	"github.com/Nowak90210/ports/internal/infrastructure/file"
	transport "github.com/Nowak90210/ports/internal/infrastructure/http"
	"github.com/Nowak90210/ports/internal/infrastructure/storage"
)

func main() {
	var wait time.Duration
	var jsonFilesFolderPath string
	var port string

	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s")
	flag.StringVar(&jsonFilesFolderPath, "folder_path", "json_files/", "the path to folder containing .json files with ports data - e.g. json_files/")
	flag.StringVar(&port, "port", "8080", "the path to folder containing .json files with ports data - e.g. 8080")
	flag.Parse()

	repo := storage.NewInMemoryRepo()
	fileReader := file.NewJsonFileReader(jsonFilesFolderPath)
	service := app.NewService(repo, fileReader)
	handler := transport.NewHandler(service)
	router := handler.GetRouter()
	srv := &http.Server{
		Addr:         fmt.Sprintf("0.0.0.0:%s", port),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		Handler:      router,
	}

	log.Println("REST API server is running on port", port)

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	srv.Shutdown(ctx)

	log.Println("shutting down")
	os.Exit(0)
}
