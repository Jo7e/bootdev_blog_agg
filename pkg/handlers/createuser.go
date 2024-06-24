package handlers

import (
	"bootdev_blog_agg/internal"
	"bootdev_blog_agg/internal/database"
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type CreateUserRequest struct {
	Name string `json:"name"`
}

type CreateUserResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Apikey    string    `json:"apikey"`
}

func CreateUserHandler(a *ApiConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		userRequest := CreateUserRequest{}
		err := json.NewDecoder(r.Body).Decode(&userRequest)
		if err != nil {
			internal.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}

		now := time.Now()

		createUserParams := database.CreateUserParams{
			ID:        uuid.New(),
			CreatedAt: now,
			UpdatedAt: now,
			Name:      userRequest.Name,
		}

		user, err := a.DB.CreateUser(ctx, createUserParams)
		if err != nil {
			internal.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		userResponse := CreateUserResponse{
			ID:        user.ID,
			Name:      user.Name,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			Apikey:    user.Apikey,
		}

		internal.RespondWithJSON(w, http.StatusCreated, userResponse)
	}
}