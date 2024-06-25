package feeds

import (
	"bootdev_blog_agg/internal"
	sqlconfig "bootdev_blog_agg/pkg/config"
	"net/http"
)

func GetFeedsHandler(a *sqlconfig.ApiConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		databaseFeeds, err := a.DB.GetFeeds(ctx)
		if err != nil {
			internal.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		feeds := internal.DataBaseFeedsToFeeds(databaseFeeds)

		internal.RespondWithJSON(w, http.StatusCreated, feeds)
	}
}
