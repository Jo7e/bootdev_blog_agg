package routes

import (
	sqlconfig "bootdev_blog_agg/pkg/config"
	"bootdev_blog_agg/pkg/handlers/feedfollows"
	"bootdev_blog_agg/pkg/handlers/feeds"
	"bootdev_blog_agg/pkg/handlers/users"
	"net/http"
)

func BuildRoutes(m *http.ServeMux) {
	apiCfg := sqlconfig.GetApiConfig()

	m.HandleFunc("POST /v1/users", users.CreateUserHandler(apiCfg))
	m.HandleFunc("GET /v1/users", apiCfg.MiddlewareAuth(users.GetUserByApikeyHandler(apiCfg)))

	m.HandleFunc("POST /v1/feeds", apiCfg.MiddlewareAuth(feeds.CreateFeedHandler(apiCfg)))
	m.HandleFunc("GET /v1/feeds", feeds.GetFeedsHandler(apiCfg))

	m.HandleFunc("POST /v1/feed_follows", apiCfg.MiddlewareAuth(feedfollows.CreateFeedFollowHandler(apiCfg)))
	m.HandleFunc("GET /v1/feed_follows", apiCfg.MiddlewareAuth(feedfollows.GetFeedFollowsHandler(apiCfg)))
	m.HandleFunc("DELETE /v1/feed_follows/{feedFollowId}", feedfollows.DeleteFeedFollowHandler(apiCfg))

	m.HandleFunc("GET /v1/healthz", HealthzHandler)
	m.HandleFunc("GET /v1/err", ErrHandler)
}
