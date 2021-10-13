package handlers

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/el10savio/goSetReconciliation/sync"
)

// StartSync ...
func StartSync(w http.ResponseWriter, r *http.Request) {
	// Get the values from the Set
	err := sync.Send(Set, []int{})
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("failed to send sync start")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// DEBUG log in the case of success
	// indicating the new Set
	log.WithFields(log.Fields{
		"set": Set,
	}).Debug("successful set sync")

	// Return HTTP 200 OK in the case of success
	w.WriteHeader(http.StatusOK)

	// JSON encode response value
	// json.NewEncoder(w).Encode(list)
}
