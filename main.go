package main

import (
	"bootdev_blog_agg/pkg/routes"
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

	routes.BuildRoutes(m)

	log.Fatal(s.ListenAndServe())
}
