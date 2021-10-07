package handlers

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// addBody ...
type addBody struct {
	Value uint32 `json:"value"`
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

	// Add the given value to our stored Set
	Set.AddElement(requestBody.Value)
	// if err != nil {
	// 	log.WithFields(log.Fields{"error": err}).Error("failed to add value")
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	return
	// }

	// DEBUG log in the case of success indicating
	// the new Set and the value added
	log.WithFields(log.Fields{
		"list":  Set.GetElements(),
		"value": requestBody.Value,
	}).Debug("successful wstring addition")

	// Return HTTP 200 OK in the case of success
	w.WriteHeader(http.StatusOK)
}
