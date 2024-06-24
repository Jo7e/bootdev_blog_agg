package handlers

import (
	"bootdev_blog_agg/internal"
	"bootdev_blog_agg/internal/database"
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

func GetUserByApikeyHandler(a *ApiConfig) authedHandler {
	return func(w http.ResponseWriter, r *http.Request, user database.User) {
		userResponse := GetUserByApikeyResponse{
			ID:        user.ID,
			Name:      user.Name,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			Apikey:    user.Apikey,
		}

		internal.RespondWithJSON(w, http.StatusCreated, userResponse)
	}
}
