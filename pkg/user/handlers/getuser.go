package userhandler

import (
	"bootdev_blog_agg/internal"
	"bootdev_blog_agg/internal/auth"
	"bootdev_blog_agg/sql/config"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type GetUserByApikeyResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Apikey    string    `json:"api_key"`
}

type GetUserCfg config.ApiConfig

func GetUserByApikeyHandler(w http.ResponseWriter, r *http.Request) {
	a := GetUserCfg(*config.GetApiConfig())
	a.GetUserByApikey(w, r)
}

func (a *GetUserCfg) GetUserByApikey(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	auth, err := auth.GetAuthApikey(r)
	if err != nil {
		internal.RespondWithError(w, http.StatusUnauthorized, err.Error())
		return
	}

	user, err := a.DB.GetUserByApikey(ctx, auth)
	if err != nil {
		internal.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	userResponse := GetUserByApikeyResponse{
		ID:        user.ID,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Apikey:    user.Apikey,
	}

	internal.RespondWithJSON(w, http.StatusCreated, userResponse)
}
