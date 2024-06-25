package handlers

import (
	"bootdev_blog_agg/internal"
	"bootdev_blog_agg/internal/database"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type GetFeedFollowsResponse struct {
	ID        uuid.UUID `json:"id"`
	FeedID    uuid.UUID `json:"feed_id"`
	UserID    uuid.UUID `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func GetFeedFollowsHandler(a *ApiConfig) authedHandler {
	return func(w http.ResponseWriter, r *http.Request, user database.User) {
		ctx := r.Context()

		feedFollows, err := a.DB.GetFeedFollows(ctx, user.ID)
		if err != nil {
			internal.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		feedFollowsResponse := make([]GetFeedFollowsResponse, 0, len(feedFollows))

		for _, feedFollow := range feedFollows {
			feedFollowResponse := GetFeedFollowsResponse{
				ID:        feedFollow.ID,
				FeedID:    feedFollow.FeedID,
				UserID:    feedFollow.UserID,
				CreatedAt: feedFollow.CreatedAt,
				UpdatedAt: feedFollow.UpdatedAt,
			}
			feedFollowsResponse = append(feedFollowsResponse, feedFollowResponse)
		}

		internal.RespondWithJSON(w, http.StatusCreated, feedFollowsResponse)
	}
}
