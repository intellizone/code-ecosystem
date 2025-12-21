package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

type config struct {
	port    int
	env     string
	version string
}

var cfg config

func main() {

	flag.IntVar(&cfg.port, "port", 8080, "Application serve port, eg: 8080")
	flag.StringVar(&cfg.env, "env", "dev", "environment: dev|eng|prod")
	flag.StringVar(&cfg.version, "version", "v0.0.0", "Application version")
	flag.Parse()

	// if env has VERSION use it and save to cfg.version else use version v0.0.0
	if version := os.Getenv("VERSION"); version != "" {
		cfg.version = version
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/api/v1/healthcheck", healthcheck)

	log.Printf("Starting server in port :%d...", cfg.port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.port), mux)
	if err != nil {
		log.Println(err)
	}
}

func healthcheck(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Write([]byte("status: ok\n"))
		fmt.Fprintf(w, "Environment: %s\n", cfg.env)
		fmt.Fprintf(w, "Version: %s\n", cfg.version)
	} else {
		fmt.Fprintf(w, "Unsupported method: %v\n", r.Method)
	}
}
