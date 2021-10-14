package handlers

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/el10savio/goSetReconciliation/sync"
)

// Reconcile is the internal HTTP handler for a peer to
// receive the sync update from peer Sets in the cluster
func Reconcile(w http.ResponseWriter, r *http.Request) {
	var payload sync.Payload

	// Obtain the payload from POST Request Body
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&payload)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("failed parse request body")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	missingElements := []int{}

	// Reconcile the given value to our stored Set
	Set, missingElements = sync.Update(Set, payload)

	// Phase 2 corresponds with sending either the missing elements or the Bloom Filter
	// from our Set to other peers as we have found there to be a mismatch
	// between the parameters received and the new updated set
	if len(missingElements) > 0 || Set.Hash != payload.Hash {
		err := sync.Send(Set, missingElements)
		if err != nil {
			log.WithFields(log.Fields{"error": err}).Error("Phase 2 error")
		}
		log.WithFields(log.Fields{
			"set": Set, "missing elements": missingElements,
		}).Debug("completed sync phase 2")
	}

	// DEBUG log in the case of success indicating
	// the new Set and the values added
	log.WithFields(log.Fields{
		"set":              Set,
		"missing elements": missingElements,
	}).Debug("successful set sync")

	// Return HTTP 200 OK in the case of success
	w.WriteHeader(http.StatusOK)
}
