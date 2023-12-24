package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/itishrishikesh/gofeed/auth"
	"github.com/itishrishikesh/gofeed/constants"
	"github.com/itishrishikesh/gofeed/internal/database"
	"github.com/itishrishikesh/gofeed/models"
	"github.com/itishrishikesh/gofeed/utils"
)

func (apiCfg *ApiConfig) CreateUserHandler(writer http.ResponseWriter, request *http.Request) {
	type parameters struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	params := parameters{}
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&params)
	if err != nil {
		utils.RespondWithError(writer, constants.HTTP_BAD_REQUEST, "Error parsing JSON")
		log.Println("E#1PDKJ2 - User passed incorrect JSON")
		return
	}

	user, err := apiCfg.DB.CreateUser(request.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      params.Username,
		Password:  utils.HashPassword(params.Password),
	})
	if err != nil {
		utils.RespondWithError(writer, constants.HTTP_ERROR, "Couldn't create user")
		log.Println(fmt.Sprintf("E#1ORI9D - Couldn't create user %v", err))
		return
	}

	utils.RespondWithJSON(writer, constants.HTTP_SUCCESS, models.DatabaseUserToUser(user))
}

func (apiCfg *ApiConfig) GetUserHandler(writer http.ResponseWriter, request *http.Request, user models.User) {
	utils.RespondWithJSON(writer, constants.HTTP_SUCCESS, user)
}

func (config *ApiConfig) GetPostsForUserHandler(writer http.ResponseWriter, request *http.Request, user models.User) {
	posts, err := config.DB.GetPostsForUser(request.Context(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  10,
	})
	if err != nil {
		utils.RespondWithError(writer, constants.HTTP_ERROR, "Couldn't get posts")
		log.Println(fmt.Sprintf("E#1OWWJW - Couldn't get posts %v", err))
		return
	}

	utils.RespondWithJSON(writer, constants.HTTP_SUCCESS, models.DatabasePostsToPosts(posts))
}

func (config *ApiConfig) TokenHandler(writer http.ResponseWriter, request *http.Request) {
	type parameters struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	params := parameters{}
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&params)
	if err != nil {
		utils.RespondWithError(writer, constants.HTTP_BAD_REQUEST, "Error parsing JSON")
		log.Println("E#1PFLY0 - Something's wrong with token", err)
		return
	}
	user, err := config.DB.GetUserByUsername(context.Background(), params.Username)
	log.Println("D#1PFMM0 - Received user is", user)
	if err != nil {
		utils.RespondWithError(writer, constants.HTTP_UNAUTHORIZED, "Invalid Username or Password")
		log.Println("E#1PFM6A - Invalid Username or Password. E:", err)
		return
	}
	if !utils.CompareHashAndPassword(user.Password, params.Password) {
		utils.RespondWithError(writer, constants.HTTP_UNAUTHORIZED, "Invalid Username or Password")
		log.Println("E#1PHB77 - Invalid Username or Password. E:", err)
		return
	}
	token, err := auth.GenerateJWT(user.Name)
	if err != nil {
		utils.RespondWithError(writer, constants.HTTP_ERROR, "Error while generating token")
		log.Println(fmt.Sprintf("E#1PFKYV - Error while generating JWT Token %v", err))
		return
	}
	type returnToken struct {
		Token string `json:"token"`
	}
	payload := returnToken{
		Token: token,
	}
	utils.RespondWithJSON(writer, constants.HTTP_SUCCESS, payload)
}
