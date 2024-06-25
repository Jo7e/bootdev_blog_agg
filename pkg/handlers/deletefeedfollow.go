package handlers

import (
	"bootdev_blog_agg/internal"
	"net/http"

	"github.com/google/uuid"
)

func DeleteFeedFollowHandler(a *ApiConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		feedFollowIdString := r.PathValue("feedFollowId")
		feedFollowId, err := uuid.Parse(feedFollowIdString)
		if err != nil {
			internal.RespondWithError(w, http.StatusBadRequest, "Invalid FeedFollowId")
			return
		}

		err = a.DB.DeleteFeedFollow(ctx, feedFollowId)
		if err != nil {
			internal.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		internal.RespondWithJSON(w, http.StatusAccepted, nil)
	}
}
