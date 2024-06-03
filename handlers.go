package main

import (
	"bootdev_blog_agg/internal/database"
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func HealthzHandler(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func ErrHandler(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, http.StatusInternalServerError, "Internal Server Error")
}

type CreateUserRequest struct {
	Name string `json:"name"`
}

func (a *apiConfig) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userRequest := CreateUserRequest{}
	err := json.NewDecoder(r.Body).Decode(&userRequest)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	now := time.Now()

	user := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
		Name:      userRequest.Name,
	}

	userResponse, err := a.DB.CreateUser(ctx, user)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, userResponse)
}
