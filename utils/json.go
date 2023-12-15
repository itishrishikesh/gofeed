package utils

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/itishrishikesh/gofeed/constants"
)

func RespondWithJSON(writer http.ResponseWriter, code int, payload interface{}) {
	log.Println("I#1ONRD6 - Trying to respond with json payload", payload)
	data, err := json.Marshal(payload)
	if err != nil {
		log.Fatalln("E#1OLY17 - failed to marshal JSON response", payload, err)
		writer.WriteHeader(constants.HTTP_ERROR)
	}
	writer.Header().Add(constants.CONTENT_TYPE, constants.APPLICATION_JSON)
	writer.WriteHeader(code)
	writer.Write(data)
}

func RespondWithError(writer http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Println("I#1ONRP0 - Responding with 5XX error", msg)
	}
	type errResponse struct {
		Error string `json:"error"`
	}

	RespondWithJSON(writer, code, errResponse{Error: msg})
}
