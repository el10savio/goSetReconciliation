package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/el10savio/goSetReconciliation/sync"
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

	missingElements := []int{}
	oldHash := Set.Hash

	// Reconcile the given value to our stored Set
	Set, missingElements = sync.Update(Set, payload)

	if len(missingElements) > 0 || oldHash != Set.Hash {
		fmt.Println("Phase 2", missingElements)
		if err := sync.Send(Set, missingElements); err != nil {
			log.WithFields(log.Fields{"error": err}).Error("Phase 2 error")
		}
	}

	// DEBUG log in the case of success indicating
	// the new Set and the value added
	log.WithFields(log.Fields{
		"set":              Set,
		"missing elements": missingElements,
	}).Debug("successful set sync")

	// Return HTTP 200 OK in the case of success
	w.WriteHeader(http.StatusOK)
}
