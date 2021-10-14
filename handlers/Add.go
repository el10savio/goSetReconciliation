package handlers

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// addBody is the format of the
// input JSON body to the
// Add Handler
type addBody struct {
	Values []int `json:"values"`
}

// Add is the HTTP handler used to add in
// values to the Set node in the server
func Add(w http.ResponseWriter, r *http.Request) {
	var requestBody addBody

	// Obtain the values from POST Request Body
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestBody)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("failed parse request body")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Add the values into the Set
	Set.AddElements(requestBody.Values)

	// DEBUG log in the case of success indicating
	// the new Set and the values added
	log.WithFields(log.Fields{
		"list":   Set.GetElements(),
		"values": requestBody.Values,
	}).Debug("successful set addition")

	// Return HTTP 200 OK
	// in the case of success
	w.WriteHeader(http.StatusOK)
}
