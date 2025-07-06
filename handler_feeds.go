package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/a-rezk/rssagg/internal/database"
	"github.com/go-chi/chi/v5"

	"github.com/google/uuid"
)

func (apiCfg *apiConfig) createFeedHandler(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithErr(w, 400, fmt.Sprintf("Error Parsing JSON:%v", err))
		return
	}

	newFeedID := uuid.NewString()

	_, err = apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        newFeedID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.URL,
		UserID:    user.ID,
	})

	if err != nil {
		respondWithErr(w, 400, fmt.Sprintf("Couldn't create feed: %v", err))
		return
	}

	dbFeed, err := apiCfg.DB.GetFeedByID(r.Context(), newFeedID)
	if err != nil {
		respondWithErr(w, 400, fmt.Sprintf("Couldn't get created user: %v", err))
		return
	}

	feed := dbFeedtoFeed(dbFeed)

	respondWithJSON(w, 201, feed)

}

func (apicfg *apiConfig) GetFeedsHandler(w http.ResponseWriter, r *http.Request) {
	feeds, err := apicfg.DB.GetFeeds(r.Context())
	if err != nil {
		respondWithErr(w, 400, "Couldn't get feeds")
		return
	}

	respondWithJSON(w, 200, dbFeedstoFeeds(feeds))
}

func (apiCfg *apiConfig) DeleteFeedHandler(w http.ResponseWriter, r *http.Request, user database.User) {
	feedIDString := chi.URLParam(r, "feedID")

	err := apiCfg.DB.DeleteFeedByID(r.Context(), database.DeleteFeedByIDParams{
		ID:     feedIDString,
		UserID: user.ID,
	})

	if err != nil {
		respondWithErr(w, 400, "Couldn't Delete feed!")
		return
	}

	respondWithJSON(w, 200, "Feed Deleted successfully")
}
