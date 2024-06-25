package sqlconfig

import (
	"bootdev_blog_agg/internal"
	"bootdev_blog_agg/internal/auth"
	"bootdev_blog_agg/internal/database"
	"database/sql"
	"log"
	"net/http"
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

type AuthedHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *ApiConfig) MiddlewareAuth(handler AuthedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		auth, err := auth.GetAuthApikey(r)
		if err != nil {
			internal.RespondWithError(w, http.StatusUnauthorized, err.Error())
			return
		}

		user, err := cfg.DB.GetUserByApikey(ctx, auth)
		if err != nil {
			internal.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		handler(w, r, user)
	}
}
