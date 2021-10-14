package handlers

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/el10savio/goSetReconciliation/sync"
)

// StartSync is the HTTP handler used to initiate
// the sync procedure between the Sets
// in the cluster
func StartSync(w http.ResponseWriter, r *http.Request) {
	// Send our Set to peers
	err := sync.Send(Set, []int{})
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("failed to send sync start")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// DEBUG log in the case of success
	log.WithFields(log.Fields{
		"set": Set,
	}).Debug("successful set sync start")

	// Return HTTP 200 OK
	// in the case of success
	w.WriteHeader(http.StatusOK)
}
