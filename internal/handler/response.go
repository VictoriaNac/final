package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"
)

type messageResponse struct {
	Message string `json:"message"`
}

type listResponse struct {
	CityNames []string `json:"cityNames"`
}


func newListResponse(w http.ResponseWriter, statusCode int, response []string) {
	w.WriteHeader(statusCode)

	message, err := json.Marshal(listResponse{response})
	if err != nil {
		log.Err(err).Msg("JSON Serialization failed")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("json.Marshal error"))
		return
	}

	log.Info().Msg(fmt.Sprintf("Sending: %v", string(message)))
	w.Write(message)
}


func newMessageResponse(w http.ResponseWriter, statusCode int, response string) {
	w.WriteHeader(statusCode)

	message, err := json.Marshal(messageResponse{response})
	if err != nil {
		log.Warn().Msg("JSON Serialization failed")
		w.Write([]byte(response + " ::json.Marshal error"))
		return
	}

	log.Info().Msg(fmt.Sprintf("Sending: %v", string(message)))
	w.Write(message)
}