package auth

import (
	"errors"
	"net/http"
	"strings"
)

var ErrNoAuthHeader = errors.New("missing Authorization header")
var ErrInvalidAuthHeader = errors.New("invalid Authorization header")

func GetAuthApikey(r *http.Request) (string, error) {
	auth := r.Header.Get("Authorization")
	if auth == "" {
		return "", ErrNoAuthHeader
	}

	splitAuth := strings.Split(auth, " ")
	if len(splitAuth) < 2 || splitAuth[0] != "Apikey" {
		return "", ErrInvalidAuthHeader
	}

	return splitAuth[1], nil
}
