package handlers

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// addBody ...
type addBody struct {
	Values []int `json:"values"`
}

// Add ...
func Add(w http.ResponseWriter, r *http.Request) {
	var err error
	var requestBody addBody

	// Obtain the value & position from POST Request Body
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&requestBody)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("failed parse request body")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	Set.AddElements(requestBody.Values)
	// if err != nil {
	// 	log.WithFields(log.Fields{"error": err}).Error("failed to add value")
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	return
	// }

	// DEBUG log in the case of success indicating
	// the new Set and the value added
	log.WithFields(log.Fields{
		"list":   Set.GetElements(),
		"values": requestBody.Values,
	}).Debug("successful set addition")

	// Return HTTP 200 OK in the case of success
	w.WriteHeader(http.StatusOK)
}
