package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var version string

func main() {

	version = os.Getenv("VERSION")
	if len(version) == 0 {
		flag.StringVar(&version, "version", "", "version of the app")
		flag.Parse()
	}

	if len(version) == 0 {
		log.Fatal("please provide a valid version")
	}

	app := &http.Server{
		Addr:        ":8080",
		ReadTimeout: 3 * time.Second,
	}

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello Suriya")
	})

	http.HandleFunc("/metrics", metrics)

	http.HandleFunc("/api/v1/healthcheck", healthCheck)

	go func() {
		err := app.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()
	log.Printf("Starting application in port %s\n", app.Addr)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down gracefully")
	app.Shutdown(context.Background())
	log.Println("App stopped gracefully")

}

func metrics(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "my_metrics 2")
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"version": version,
		"status":  "ok",
	}

	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Println(err)
	}
	w.Header().Add("content-type", "application/json")
	w.Write(jsonData)
}
