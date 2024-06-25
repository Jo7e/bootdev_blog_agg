package handlers

import (
	"bootdev_blog_agg/internal"
	"bootdev_blog_agg/internal/database"
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type CreateFeedFollowRequest struct {
	FeedId uuid.UUID `json:"feed_id"`
}

type CreateFeedFollowResponse struct {
	ID        uuid.UUID `json:"id"`
	FeedID    uuid.UUID `json:"feed_id"`
	UserID    uuid.UUID `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func CreateFeedFollowHandler(a *ApiConfig) authedHandler {
	return func(w http.ResponseWriter, r *http.Request, user database.User) {
		ctx := r.Context()

		feedRequest := CreateFeedFollowRequest{}
		err := json.NewDecoder(r.Body).Decode(&feedRequest)
		if err != nil {
			internal.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}

		now := time.Now()

		createFeedFollowParams := database.CreateFeedFollowParams{
			ID:        uuid.New(),
			FeedID:    uuid.UUID(feedRequest.FeedId),
			UserID:    user.ID,
			CreatedAt: now,
			UpdatedAt: now,
		}

		feed, err := a.DB.CreateFeedFollow(ctx, createFeedFollowParams)
		if err != nil {
			internal.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		feedFollowResponse := CreateFeedFollowResponse{
			ID:        feed.ID,
			FeedID:    feed.FeedID,
			UserID:    feed.UserID,
			CreatedAt: feed.CreatedAt,
			UpdatedAt: feed.UpdatedAt,
		}

		internal.RespondWithJSON(w, http.StatusCreated, feedFollowResponse)
	}
}
