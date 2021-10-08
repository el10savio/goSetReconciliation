package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	log "github.com/sirupsen/logrus"
)

// addBody ...
type addBody struct {
	Value string `json:"value"`
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
	element, err := strconv.ParseUint(requestBody.Value, 10, 32)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("failed to parse value")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	Set.AddElement(uint32(element))
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
