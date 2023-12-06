package main

import (
	"log"
	"net/http"

	"github.com/itishrishikesh/gofeed/constants"
)

func errorHandler(writer http.ResponseWriter, request *http.Request) {
	log.Println("D#1ONS9C - Request received:", request)
	respondWithError(writer, constants.HTTP_BAD_REQUEST, "Something went wrong!")
}
