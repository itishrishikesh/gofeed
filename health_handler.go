package main

import (
	"log"
	"net/http"

	"github.com/itishrishikesh/gofeed/constants"
)

func healthCheckHandler(writer http.ResponseWriter, request *http.Request) {
	log.Println("D#1ONRAR - Request received:", request)
	respondWithJSON(writer, constants.HTTP_SUCCESS, struct{}{})
}
