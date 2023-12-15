package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/itishrishikesh/gofeed/constants"
	"github.com/itishrishikesh/gofeed/internal/database"
	"github.com/itishrishikesh/gofeed/models"
	"github.com/itishrishikesh/gofeed/utils"
)

func (apiCfg *ApiConfig) CreateFeedHandler(writer http.ResponseWriter, request *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
	params := parameters{}
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&params)
	if err != nil {
		utils.RespondWithError(writer, constants.HTTP_BAD_REQUEST, "Error parsing JSON")
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
		utils.RespondWithError(writer, constants.HTTP_BAD_REQUEST, fmt.Sprintf("E#1OV5DR - Couldn't create feed %v", err))
		return
	}

	utils.RespondWithJSON(writer, constants.HTTP_CREATED, models.DatabaseFeedToFeed(feed))
}

func (config *ApiConfig) GetFeedsHandler(writer http.ResponseWriter, request *http.Request) {
	feeds, err := config.DB.GetFeed(request.Context())
	if err != nil {
		utils.RespondWithError(writer, constants.HTTP_BAD_REQUEST, fmt.Sprintf("E#1OV63D - Couldn't get feed %v", err))
		return
	}
	utils.RespondWithJSON(writer, constants.HTTP_CREATED, models.DatabaseFeedToFeeds(feeds))
}

func (config *ApiConfig) CreateFeedFollowHandler(writer http.ResponseWriter, request *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}
	params := parameters{}
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&params)
	if err != nil {
		utils.RespondWithError(writer, constants.HTTP_BAD_REQUEST, "Error parsing JSON")
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
		utils.RespondWithError(writer, constants.HTTP_BAD_REQUEST, fmt.Sprintf("E#1OV7DV - Couldn't create feed follow %v", err))
		return
	}
	utils.RespondWithJSON(writer, constants.HTTP_CREATED, models.DatabaseFeedFollowToFeedFollow(feedFollow))
}

func (config *ApiConfig) GetFeedFollowsHandler(writer http.ResponseWriter, request *http.Request, user database.User) {
	feedFollows, err := config.DB.GetFeedFollows(request.Context(), user.ID)
	if err != nil {
		utils.RespondWithError(writer, constants.HTTP_BAD_REQUEST, fmt.Sprintf("E#1OWRVV - Couldn't get feed follows %v", err))
		return
	}
	utils.RespondWithJSON(writer, constants.HTTP_CREATED, models.DatabaseFeedFollowToFeedFollows(feedFollows))
}

func (config *ApiConfig) DeleteFeedHandler(writer http.ResponseWriter, request *http.Request, user database.User) {
	feedFollowId := chi.URLParam(request, "feedFollowId")
	feedFollowUUID, err := uuid.Parse(feedFollowId)
	if err != nil {
		utils.RespondWithError(writer, constants.HTTP_BAD_REQUEST, fmt.Sprintf("E#1OWSD0 - Feed follow ID isn't correct %v", err))
		return
	}
	err = config.DB.DeleteFeedFollow(request.Context(), database.DeleteFeedFollowParams{
		ID:     feedFollowUUID,
		UserID: user.ID,
	})
	if err != nil {
		utils.RespondWithError(writer, constants.HTTP_BAD_REQUEST, fmt.Sprintf("E#1OWSG3 - Couldn't delete feed follow %v", err))
		return
	}
	utils.RespondWithJSON(writer, constants.HTTP_SUCCESS, struct{}{})
}
