package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	log "github.com/sirupsen/logrus"
)

// Reconcile ...
func Reconcile(w http.ResponseWriter, r *http.Request) {
	var err error
	var payload sync.Payload

	// Obtain the value & position from POST Request Body
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&payload)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("failed parse request body")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Reconcile the given value to our stored Set
	err := sync.Update(payload)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("failed to sync")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// DEBUG log in the case of success indicating
	// the new Set and the value added
	log.WithFields(log.Fields{
		"set":   Set,
		"value": payload.Value,
	}).Debug("successful set sync")

	// Return HTTP 200 OK in the case of success
	w.WriteHeader(http.StatusOK)
}
