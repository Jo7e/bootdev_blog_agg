package main

import (
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	add := os.Getenv("HOST") + os.Getenv("PORT")
	m := http.NewServeMux()

	s := http.Server{
		Addr:    add,
		Handler: m,
	}

	m.HandleFunc("GET /v1/healthz", HealthzHandler)
	m.HandleFunc("GET /v1/err", ErrHandler)

	s.ListenAndServe()
}
