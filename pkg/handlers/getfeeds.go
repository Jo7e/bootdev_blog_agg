package handlers

import (
	"bootdev_blog_agg/internal"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type GetFeedsResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	UserID    uuid.UUID `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func GetFeedsHandler(a *ApiConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		feeds, err := a.DB.GetFeeds(ctx)
		if err != nil {
			internal.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		feedsResponse := make([]GetFeedsResponse, 0, len(feeds))

		for _, feed := range feeds {
			feedResponse := GetFeedsResponse{
				ID:        feed.ID,
				Name:      feed.Name,
				Url:       feed.Url,
				UserID:    feed.UserID,
				CreatedAt: feed.CreatedAt,
				UpdatedAt: feed.UpdatedAt,
			}
			feedsResponse = append(feedsResponse, feedResponse)
		}

		internal.RespondWithJSON(w, http.StatusCreated, feedsResponse)
	}
}
