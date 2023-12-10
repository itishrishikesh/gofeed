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

func (apiCfg *apiConfig) createFeedHandler(writer http.ResponseWriter, request *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
	params := parameters{}
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(writer, constants.HTTP_BAD_REQUEST, "Error parsing JSON")
		log.Println("D#1OV4W5 - User passed incorrect JSON")
		return
	}
	feed, err := apiCfg.DB.CreateFeed(request.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      params.Name,
		Url:       params.URL,
		UserID:    user.ID,
	})
	if err != nil {
		respondWithError(writer, constants.HTTP_BAD_REQUEST, fmt.Sprintf("E#1OV5DR - Couldn't create feed %v", err))
		return
	}

	respondWithJSON(writer, constants.HTTP_CREATED, databaseFeedToFeed(feed))
}

func (config *apiConfig) getFeedsHandler(writer http.ResponseWriter, request *http.Request) {
	feeds, err := config.DB.GetFeed(request.Context())
	if err != nil {
		respondWithError(writer, constants.HTTP_BAD_REQUEST, fmt.Sprintf("E#1OV63D - Couldn't get feed %v", err))
		return
	}
	respondWithJSON(writer, constants.HTTP_CREATED, databaseFeedToFeeds(feeds))
}

func (config *apiConfig) createFeedFollowHandler(writer http.ResponseWriter, request *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}
	params := parameters{}
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(writer, constants.HTTP_BAD_REQUEST, "Error parsing JSON")
		log.Println("D#1OV7DL - User passed incorrect JSON")
		return
	}
	feedFollow, err := config.DB.CreateFeedFollow(request.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})
	if err != nil {
		respondWithError(writer, constants.HTTP_BAD_REQUEST, fmt.Sprintf("E#1OV7DV - Couldn't create feed follow %v", err))
		return
	}
	respondWithJSON(writer, constants.HTTP_CREATED, databaseFeedFollowToFeedFollow(feedFollow))
}
