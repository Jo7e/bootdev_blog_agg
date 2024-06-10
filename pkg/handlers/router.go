package handlers

import (
	"bootdev_blog_agg/internal"
	"bootdev_blog_agg/internal/auth"
	"bootdev_blog_agg/internal/database"
	"net/http"
)

func BuildRoutes(m *http.ServeMux) {
	apiCfg := GetApiConfig()

	m.HandleFunc("POST /v1/users", CreateUserHandler(apiCfg))
	m.HandleFunc("GET /v1/users", apiCfg.middlewareAuth(GetUserByApikeyHandler(apiCfg)))

	m.HandleFunc("POST /v1/feeds", apiCfg.middlewareAuth(CreateFeedHandler(apiCfg)))

	m.HandleFunc("GET /v1/healthz", HealthzHandler)
	m.HandleFunc("GET /v1/err", ErrHandler)
}

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *ApiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		auth, err := auth.GetAuthApikey(r)
		if err != nil {
			internal.RespondWithError(w, http.StatusUnauthorized, err.Error())
			return
		}

		user, err := cfg.DB.GetUserByApikey(ctx, auth)
		if err != nil {
			internal.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		handler(w, r, user)
	}
}
