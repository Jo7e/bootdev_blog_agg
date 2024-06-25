package users

import (
	"bootdev_blog_agg/internal"
	"bootdev_blog_agg/internal/database"
	sqlconfig "bootdev_blog_agg/pkg/config"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type GetUserByApikeyResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Apikey    string    `json:"apikey"`
}

func GetUserByApikeyHandler(a *sqlconfig.ApiConfig) sqlconfig.AuthedHandler {
	return func(w http.ResponseWriter, r *http.Request, user database.User) {
		userResponse := internal.DatabaseUserToUser(user)

		internal.RespondWithJSON(w, http.StatusCreated, userResponse)
	}
}
