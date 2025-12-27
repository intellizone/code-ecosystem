package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"readinglist/internal/models"
	"strconv"
	"time"
)

type config struct {
	Addr     int
	Endpoint string
}

type application struct {
	config      config
	logger      *log.Logger
	readinglist *models.ReadinglistModel
}

func main() {
	cfg := config{}
	flag.IntVar(&cfg.Addr, "port", 8080, "Application serve port, eg: 8080")
	url := flag.String("endpoint", "http://localhost:8081", "url of readinglist app")
	flag.Parse()

	if addr := os.Getenv("PORT"); addr != "" {
		port, err := strconv.Atoi(addr)
		if err != nil {
			log.Fatalln(err)
		}
		if port <= 0 || port > 65535 {
			log.Fatal("invalid port number")
		}
		cfg.Addr = port
	}

	if len(os.Getenv("URL")) > 0 {
		*url = os.Getenv("URL")
	}
	cfg.Endpoint = fmt.Sprintf("%v/api/v1/books", *url)

	addr := fmt.Sprintf(":%v", cfg.Addr)
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := &application{
		config:      cfg,
		logger:      logger,
		readinglist: &models.ReadinglistModel{Endpoint: cfg.Endpoint},
	}

	srv := &http.Server{
		Addr:         addr,
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	app.logger.Printf("Application started in port %v\n", app.config.Addr)
	err := srv.ListenAndServe()
	if err != nil {
		app.logger.Fatalln(err)
	}
}
