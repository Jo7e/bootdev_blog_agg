package handlers

import (
	"bootdev_blog_agg/internal"
	"net/http"
)

func HealthzHandler(w http.ResponseWriter, r *http.Request) {
	internal.RespondWithJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func ErrHandler(w http.ResponseWriter, r *http.Request) {
	internal.RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
}
