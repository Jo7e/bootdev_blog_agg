package feedfollows

import (
	"bootdev_blog_agg/internal"
	"bootdev_blog_agg/internal/database"
	sqlconfig "bootdev_blog_agg/pkg/config"
	"net/http"
)

func GetFeedFollowsHandler(a *sqlconfig.ApiConfig) sqlconfig.AuthedHandler {
	return func(w http.ResponseWriter, r *http.Request, user database.User) {
		ctx := r.Context()

		databaseFeedFollows, err := a.DB.GetFeedFollows(ctx, user.ID)
		if err != nil {
			internal.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		feedFollows := internal.DatabaseFeedFollowsToFeedFollows(databaseFeedFollows)

		internal.RespondWithJSON(w, http.StatusCreated, feedFollows)
	}
}
