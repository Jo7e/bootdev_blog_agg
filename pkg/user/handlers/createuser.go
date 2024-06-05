package userhandler

import (
	"bootdev_blog_agg/internal"
	"bootdev_blog_agg/internal/database"
	"bootdev_blog_agg/sql/config"
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type CreateUserRequest struct {
	Name string `json:"name"`
}

type CreateUserCfg config.ApiConfig

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	a := CreateUserCfg(*config.GetApiConfig())
	a.CreateUser(w, r)
}

func (a *CreateUserCfg) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userRequest := CreateUserRequest{}
	err := json.NewDecoder(r.Body).Decode(&userRequest)
	if err != nil {
		internal.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
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
		internal.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	internal.RespondWithJSON(w, http.StatusCreated, userResponse)
}
