package handlers

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

// Clear is the internal util HTTP handler
// used to clear the set and reset its parameters
func Clear(w http.ResponseWriter, r *http.Request) {
	// Clear the values from the Set
	Set.Clear()

	// DEBUG log in the case of success
	// indicating the new Set
	log.WithFields(log.Fields{
		"set": Set,
	}).Debug("successful set clear")

	// Return HTTP 200 OK in the case of success
	w.WriteHeader(http.StatusOK)
}
