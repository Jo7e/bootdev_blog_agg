package main

import (
	"bootdev_blog_agg/internal/database"
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	dbURL := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", dbURL)

	if err != nil {
		log.Fatal(err)
	}

	dbQueries := database.New(db)
	apiConfig := &apiConfig{
		DB: dbQueries,
	}

	add := os.Getenv("HOST") + ":" + os.Getenv("PORT")
	m := http.NewServeMux()

	s := http.Server{
		Addr:    add,
		Handler: m,
	}

	m.HandleFunc("GET /v1/users", apiConfig.CreateUserHandler)
	m.HandleFunc("GET /v1/healthz", HealthzHandler)
	m.HandleFunc("GET /v1/err", ErrHandler)

	log.Fatal(s.ListenAndServe())
}
