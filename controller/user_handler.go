package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/itishrishikesh/gofeed/constants"
	"github.com/itishrishikesh/gofeed/internal/database"
	"github.com/itishrishikesh/gofeed/models"
	"github.com/itishrishikesh/gofeed/utils"
	"golang.org/x/crypto/bcrypt"
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

	hash, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcrypt.MinCost)
	if err != nil {
		utils.RespondWithError(writer, constants.HTTP_BAD_REQUEST, "Error parsing password")
		log.Println("E#1PDKIV - Please check your password string")
		return
	}

	user, err := apiCfg.DB.CreateUser(request.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      params.Username,
		Password:  string(hash),
	})
	if err != nil {
		utils.RespondWithError(writer, constants.HTTP_ERROR, "Couldn't create user")
		log.Println(fmt.Sprintf("E#1ORI9D - Couldn't create user %v", err))
		return
	}

	utils.RespondWithJSON(writer, constants.HTTP_SUCCESS, models.DatabaseUserToUser(user))
}

func (apiCfg *ApiConfig) GetUserHandler(writer http.ResponseWriter, request *http.Request, user database.User) {
	utils.RespondWithJSON(writer, constants.HTTP_SUCCESS, user)
}

func (config *ApiConfig) GetPostsForUserHandler(writer http.ResponseWriter, request *http.Request, user database.User) {
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
