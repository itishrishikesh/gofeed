package middleware

import (
	"fmt"
	"log"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/itishrishikesh/gofeed/auth"
	"github.com/itishrishikesh/gofeed/constants"
	"github.com/itishrishikesh/gofeed/internal/database"
	"github.com/itishrishikesh/gofeed/models"
	"github.com/itishrishikesh/gofeed/utils"
)

type authHeader func(http.ResponseWriter, *http.Request, models.User)

type ApiConfig struct {
	DB *database.Queries
}

func (config *ApiConfig) AuthMiddleware(handler authHeader) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		apiKeyOrToken, err := auth.GetAPIKeyOrToken(request.Header)
		if err != nil {
			utils.RespondWithError(writer, constants.HTTP_FORBIDDEN, fmt.Sprintf("E#1OV41H - Authentication Error: %v", err))
			return
		}
		var user database.User
		token := auth.VerifyToken(apiKeyOrToken)
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			log.Println("I#1PK9UW - Successfully verified token")
			user, err = config.DB.GetUserByUsername(request.Context(), claims["user"].(string))
			if err != nil {
				utils.RespondWithError(writer, constants.HTTP_BAD_REQUEST, "E#1PKA5F - Couldn't authenticate User!")
				log.Println("E#1PKA5F - Failed to authenticate user", err)
				return
			}
		} else {
			// This check will happen if user hasn't provided a valid token.
			// And if the user has provided an ApiKey
			user, err = config.DB.GetUserByAPIKey(request.Context(), apiKeyOrToken)
			if err != nil {
				utils.RespondWithError(writer, constants.HTTP_BAD_REQUEST, "E#1OV44M - Couldn't authenticate User!")
				log.Println("E#1OV44M - Failed to authenticate user", err)
				return
			}
		}
		handler(writer, request, models.User{Name: user.Name, ID: user.ID})
	}
}
