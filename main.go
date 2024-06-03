package main

import (
	generichandler "bootdev_blog_agg/pkg/generic"
	userhandler "bootdev_blog_agg/pkg/user/handlers"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	add := os.Getenv("HOST") + ":" + os.Getenv("PORT")
	m := http.NewServeMux()

	s := http.Server{
		Addr:    add,
		Handler: m,
	}

	m.HandleFunc("POST /v1/users", userhandler.CreateUserHandler)
	m.HandleFunc("GET /v1/healthz", generichandler.HealthzHandler)
	m.HandleFunc("GET /v1/err", generichandler.ErrHandler)

	log.Fatal(s.ListenAndServe())
}
