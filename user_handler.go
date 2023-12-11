package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/itishrishikesh/gofeed/constants"
	"github.com/itishrishikesh/gofeed/internal/database"
)

func (apiCfg *apiConfig) createUserHandler(writer http.ResponseWriter, request *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}

	params := parameters{}
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(writer, constants.HTTP_BAD_REQUEST, "Error parsing JSON")
		log.Println("D#1ORHXO - User passed incorrect JSON")
		return
	}

	user, err := apiCfg.DB.CreateUser(request.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      params.Name,
	})
	if err != nil {
		respondWithError(writer, constants.HTTP_ERROR, "Couldn't create user")
		log.Println(fmt.Sprintf("E#1ORI9D - Couldn't create user %v", err))
		return
	}

	respondWithJSON(writer, constants.HTTP_SUCCESS, databaseUserToUser(user))
}

func (apiCfg *apiConfig) getUserHandler(writer http.ResponseWriter, request *http.Request, user database.User) {
	respondWithJSON(writer, constants.HTTP_SUCCESS, user)
}

func (config *apiConfig) getPostsForUserHandler(writer http.ResponseWriter, request *http.Request, user database.User) {
	posts, err := config.DB.GetPostsForUser(request.Context(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  10,
	})
	if err != nil {
		respondWithError(writer, constants.HTTP_ERROR, "Couldn't get posts")
		log.Println(fmt.Sprintf("E#1OWWJW - Couldn't get posts %v", err))
		return
	}

	respondWithJSON(writer, constants.HTTP_SUCCESS, databasePostsToPosts(posts))
}
