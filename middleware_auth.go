package main

import (
	"net/http"

	"github.com/a-rezk/rssagg/internal/database"
	"github.com/a-rezk/rssagg/internal/database/auth"
)

type authHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIkey(r.Header)
		if err != nil {
			respondWithErr(w, 403, "Authentication Faild.")
			return
		}

		usr, err := apiCfg.DB.GetUserByAPIkey(r.Context(), apiKey)
		if err != nil {
			respondWithErr(w, 400, "Couldn't get user")
			return
		}
		handler(w, r, usr)
	}
}
