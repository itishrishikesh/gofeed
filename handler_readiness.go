package main

import (
	"net/http"

	"github.com/itishrishikesh/gofeed/constants"
)

func handlerReadiness(writer http.ResponseWriter, request *http.Request) {
	respondWithJSON(writer, constants.HTTP_SUCCESS, struct{}{})
}
