package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/a-rezk/rssagg/internal/database"

	"github.com/google/uuid"
)

func (apiCfg *apiConfig) createUserHandler(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithErr(w, 400, fmt.Sprintf("Error Parsing JSON:%v", err))
		return
	}

	newUserID := uuid.NewString()

	_, err = apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        newUserID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})

	if err != nil {
		respondWithErr(w, 400, fmt.Sprintf("Couldn't create user: %v", err))
		return
	}

	user, err := apiCfg.DB.GetUserByID(r.Context(), newUserID)
	if err != nil {
		respondWithErr(w, 400, fmt.Sprintf("Couldn't get created user: %v", err))
		return
	}

	dbUser := dbUsertoUser(user)

	respondWithJSON(w, 201, dbUser)

}

func (apiCfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request, usr database.User) {

	respondWithJSON(w, 200, dbUsertoUser(usr))

}

func (apiCfg *apiConfig) handlerGetPostForUser(w http.ResponseWriter, r *http.Request, user database.User) {
	posts, err := apiCfg.DB.GetPostsForUser(r.Context(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  10,
	})

	if err != nil {
		respondWithErr(w, 400, fmt.Sprintf("Couldn't get posts: %v", err))
		return
	}

	respondWithJSON(w, 200, dbPoststoPosts(posts))

}
