package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/itishrishikesh/gofeed/constants"
)

func respondWithJSON(writer http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Fatalln("E#1OLY17 - failed to marshal JSON response", payload, err)
		writer.WriteHeader(constants.HTTP_ERROR)
	}
	writer.Header().Add(constants.CONTENT_TYPE, constants.APPLICATION_JSON)
	writer.WriteHeader(code)
	writer.Write(data)
}
