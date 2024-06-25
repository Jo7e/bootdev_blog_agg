package feeds

import (
	"bootdev_blog_agg/internal"
	"bootdev_blog_agg/internal/database"
	sqlconfig "bootdev_blog_agg/pkg/config"
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
	Feed       internal.Feed       `json:"feed"`
	FeedFollow internal.FeedFollow `json:"feed_follow"`
}

func CreateFeedHandler(a *sqlconfig.ApiConfig) sqlconfig.AuthedHandler {
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

		databaseFeed, err := a.DB.CreateFeed(ctx, createFeedParams)
		if err != nil {
			internal.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		feed := internal.DatabaseFeedToFeed(databaseFeed)

		createFeedFollowParams := database.CreateFeedFollowParams{
			ID:        uuid.New(),
			FeedID:    feed.ID,
			UserID:    user.ID,
			CreatedAt: now,
			UpdatedAt: now,
		}

		databaseFeedFollow, err := a.DB.CreateFeedFollow(ctx, createFeedFollowParams)
		if err != nil {
			internal.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		feedFollow := internal.DatabaseFeedFollowToFeedFollow(databaseFeedFollow)

		createFeedResponse := CreateFeedResponse{
			Feed:       feed,
			FeedFollow: feedFollow,
		}

		internal.RespondWithJSON(w, http.StatusCreated, createFeedResponse)
	}
}
