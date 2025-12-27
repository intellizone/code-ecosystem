package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"readinglist/internal/data"
	"time"

	_ "github.com/lib/pq"
)

type config struct {
	env     string
	port    int
	version string
	dsn     string
}

type application struct {
	config config
	logger *log.Logger
	models data.Models
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 8080, "Application serve port, eg: 8080")
	flag.StringVar(&cfg.env, "env", os.Getenv("ENV"), "environment: dev|eng|prod")
	flag.StringVar(&cfg.version, "version", os.Getenv("VERSION"), "Application version")
	dbUser := flag.String("dbUser", os.Getenv("DB_USER"), "Database username")
	dbPass := flag.String("dbPass", os.Getenv("DB_PASS"), "Database password")
	dbHost := flag.String("dbHost", os.Getenv("DB_HOST"), "Database host")
	flag.Parse()

	if len(cfg.env) == 0 {
		log.Fatal("missing env flag")
	}
	if len(cfg.version) == 0 {
		log.Fatal("missing version flag")
	}
	if len(*dbUser) == 0 {
		log.Fatal("missing dbUser flag")
	}
	if len(*dbPass) == 0 {
		log.Fatal("missing dbPass flag")
	}
	if len(*dbHost) == 0 {
		log.Fatal("missing dbHost flag")
	}

	cfg.dsn = fmt.Sprintf("postgres://%v:%v@%v/readinglist?sslmode=disable", *dbUser, *dbPass, *dbHost)

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	db, err := sql.Open("postgres", cfg.dsn)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		logger.Fatal(err)
	}
	logger.Println("db connection established...")

	app := &application{
		config: cfg,
		logger: logger,
		models: data.NewModels(db),
	}
	addr := fmt.Sprintf(":%d", app.config.port)
	srv := &http.Server{
		Addr:         addr,
		Handler:      app.route(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	app.logger.Printf("Application started, running in port %v\n", addr)
	err = srv.ListenAndServe()
	if err != nil {
		app.logger.Fatal(err)
	}
}
