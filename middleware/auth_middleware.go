package middleware

import (
	"fmt"
	"net/http"

	"github.com/itishrishikesh/gofeed/auth"
	"github.com/itishrishikesh/gofeed/constants"
	"github.com/itishrishikesh/gofeed/internal/database"
	"github.com/itishrishikesh/gofeed/utils"
)

type authHeader func(http.ResponseWriter, *http.Request, database.User)

type ApiConfig struct {
	DB *database.Queries
}

func (config *ApiConfig) AuthMiddleware(handler authHeader) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		apiKey, err := auth.GetAPIKey(request.Header)
		if err != nil {
			utils.RespondWithError(writer, constants.HTTP_FORBIDDEN, fmt.Sprintf("E#1OV41H - Authentication Error: %v", err))
			return
		}
		user, err := config.DB.GetUserByAPIKey(request.Context(), apiKey)
		if err != nil {
			utils.RespondWithError(writer, constants.HTTP_BAD_REQUEST, fmt.Sprintf("E#1OV44M - Couldn't get user: %v", err))
			return
		}
		handler(writer, request, user)
	}
}
