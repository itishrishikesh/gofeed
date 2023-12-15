package controller

import (
	"log"
	"net/http"

	"github.com/itishrishikesh/gofeed/constants"
	"github.com/itishrishikesh/gofeed/utils"
)

func HealthCheckHandler(writer http.ResponseWriter, request *http.Request) {
	log.Println("D#1ONRAR - Request received:", request)
	utils.RespondWithJSON(writer, constants.HTTP_SUCCESS, struct{}{})
}
