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
	Feed       feedResponse       `json:"feed"`
	FeedFollow feedFollowResponse `json:"feed_follow"`
}

type feedResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uuid.UUID `json:"user_id"`
}

type feedFollowResponse struct {
	ID        uuid.UUID `json:"id"`
	FeedID    uuid.UUID `json:"feed_id"`
	UserID    uuid.UUID `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
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

		feedResponse := feedResponse{
			ID:        feed.ID,
			Name:      feed.Name,
			Url:       feed.Url,
			CreatedAt: feed.CreatedAt,
			UpdatedAt: feed.UpdatedAt,
			UserID:    feed.UserID,
		}

		createFeedFollowParams := database.CreateFeedFollowParams{
			ID:        uuid.New(),
			FeedID:    feed.ID,
			UserID:    user.ID,
			CreatedAt: now,
			UpdatedAt: now,
		}

		feedFollow, err := a.DB.CreateFeedFollow(ctx, createFeedFollowParams)
		if err != nil {
			internal.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		feedFollowResponse := feedFollowResponse{
			ID:        feedFollow.ID,
			FeedID:    feedFollow.FeedID,
			UserID:    feedFollow.UserID,
			CreatedAt: feedFollow.CreatedAt,
			UpdatedAt: feedFollow.UpdatedAt,
		}

		createFeedResponse := CreateFeedResponse{
			Feed:       feedResponse,
			FeedFollow: feedFollowResponse,
		}

		internal.RespondWithJSON(w, http.StatusCreated, createFeedResponse)
	}
}
