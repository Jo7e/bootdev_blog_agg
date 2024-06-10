package handlers

import (
	"bootdev_blog_agg/internal"
	"bootdev_blog_agg/internal/database"
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type CreateFeedRequest struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type CreateFeedResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uuid.UUID `json:"user_id"`
}

func CreateFeedHandler(a *ApiConfig) authedHandler {
	return func(w http.ResponseWriter, r *http.Request, user database.User) {
		ctx := r.Context()

		feedRequest := CreateFeedRequest{}
		err := json.NewDecoder(r.Body).Decode(&feedRequest)
		if err != nil {
			internal.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}

		now := time.Now()

		createFeedParams := database.CreateFeedParams{
			ID:        uuid.New(),
			CreatedAt: now,
			UpdatedAt: now,
			Name:      feedRequest.Name,
			Url:       feedRequest.Url,
			UserID:    user.ID,
		}

		feed, err := a.DB.CreateFeed(ctx, createFeedParams)
		if err != nil {
			internal.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		feedResponse := CreateFeedResponse{
			ID:        feed.ID,
			Name:      feed.Name,
			Url:       feed.Url,
			CreatedAt: feed.CreatedAt,
			UpdatedAt: feed.UpdatedAt,
			UserID:    feed.UserID,
		}

		internal.RespondWithJSON(w, http.StatusCreated, feedResponse)
	}
}
