package handlers

import (
	"bootdev_blog_agg/internal/database"
	"database/sql"
	"log"
	"os"
)

type ApiConfig struct {
	DB *database.Queries
}

func GetApiConfig() *ApiConfig {
	dbURL := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}

	dbQueries := database.New(db)
	return &ApiConfig{
		DB: dbQueries,
	}
}
