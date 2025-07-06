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

func (apiCfg *apiConfig) createFeedFollowHandler(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID string `json:"feed_id"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithErr(w, 400, fmt.Sprintf("Error Parsing JSON:%v", err))
		return
	}

	newFeedfollowID := uuid.NewString()

	err = apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        newFeedfollowID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})

	if err != nil {
		respondWithErr(w, 400, fmt.Sprintf("Couldn't create feed Follow: %v", err))
		return
	}

	dbFeedFollow, err := apiCfg.DB.GetFeedFollow(r.Context(), newFeedfollowID)
	if err != nil {
		respondWithErr(w, 400, fmt.Sprintf("Couldn't get created user: %v", err))
		return
	}

	feedFollow := dbFeedFollowToFeedFollow(dbFeedFollow)

	respondWithJSON(w, 201, feedFollow)

}

func (apiCfg *apiConfig) GetFeedFollowsHandler(w http.ResponseWriter, r *http.Request, user database.User) {
	dbFeedFollows, err := apiCfg.DB.GetFeedFollows(r.Context(), user.ID)
	if err != nil {
		respondWithErr(w, 400, "Couldn't get Feed's Follows!")
		return
	}

	respondWithJSON(w, 200, dbFFToFeedFollows(dbFeedFollows))
}

func (apiCfg *apiConfig) DeleteFeedFollowHandler(w http.ResponseWriter, r *http.Request, user database.User) {

	feedFollowIDString := chi.URLParam(r, "feedFollowID")
	// feedFollowID, err := uuid.Parse(feedFollowIDString)

	// if err != nil {
	// 	respondWithErr(w, 400, "Couldn't parse the ID!")
	// 	return
	// }

	err := apiCfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		UserID: user.ID,
		ID:     feedFollowIDString,
	})

	if err != nil {
		respondWithErr(w, 400, "Couldn't Unfollow feed follow!")
		return
	}

	respondWithJSON(w, 200, "Unfollow feed successfully")
}
