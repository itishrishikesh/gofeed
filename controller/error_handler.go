package controller

import (
	"log"
	"net/http"

	"github.com/itishrishikesh/gofeed/constants"
	"github.com/itishrishikesh/gofeed/utils"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request) {
	log.Println("D#1ONS9C - Request received:", request)
	utils.RespondWithError(writer, constants.HTTP_BAD_REQUEST, "Something went wrong!")
}
