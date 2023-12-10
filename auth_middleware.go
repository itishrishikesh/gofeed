package main

import (
	"fmt"
	"net/http"

	"github.com/itishrishikesh/gofeed/auth"
	"github.com/itishrishikesh/gofeed/constants"
	"github.com/itishrishikesh/gofeed/internal/database"
)

type authHeader func(http.ResponseWriter, *http.Request, database.User)

func (config *apiConfig) authMiddleware(handler authHeader) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		apiKey, err := auth.GetAPIKey(request.Header)
		if err != nil {
			respondWithError(writer, constants.HTTP_FORBIDDEN, fmt.Sprintf("E#1OV41H - Authentication Error: %v", err))
			return
		}
		user, err := config.DB.GetUserByAPIKey(request.Context(), apiKey)
		if err != nil {
			respondWithError(writer, constants.HTTP_BAD_REQUEST, fmt.Sprintf("E#1OV44M - Couldn't get user: %v", err))
			return
		}
		handler(writer, request, user)
	}
}
