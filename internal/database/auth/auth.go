package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetAPIkey extract api key from the HTTP request
// Example: Authorization: APIKey {user's apikey}

func GetAPIkey(h http.Header) (string, error) {
	val := h.Get("Authorization")
	if val == "" {
		return "", errors.New("authentication Faild")
	}
	vals := strings.Split(val, " ")
	if vals[0] != "APIKey" {
		return "", errors.New("malformed authentication request")
	}
	if len(vals) != 2 {
		return "", errors.New("malformed authentication request")
	}
	return vals[1], nil
}
