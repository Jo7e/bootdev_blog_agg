package feedfollows

import (
	"bootdev_blog_agg/internal"
	"bootdev_blog_agg/internal/database"
	sqlconfig "bootdev_blog_agg/pkg/config"
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type CreateFeedFollowRequest struct {
	FeedId uuid.UUID `json:"feed_id"`
}

func CreateFeedFollowHandler(a *sqlconfig.ApiConfig) sqlconfig.AuthedHandler {
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

		databaseFeedFellow, err := a.DB.CreateFeedFollow(ctx, createFeedFollowParams)
		if err != nil {
			internal.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		feedFellow := internal.DatabaseFeedFollowToFeedFollow(databaseFeedFellow)

		internal.RespondWithJSON(w, http.StatusCreated, feedFellow)
	}
}
