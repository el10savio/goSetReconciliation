package handlers

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// GetHash ...
func GetHash(w http.ResponseWriter, r *http.Request) {
	// Get the Hash from the Set
	hash := Set.GetHash()

	// DEBUG log in the case of success
	// indicating the Hash
	log.WithFields(log.Fields{
		"list": Set.GetElements(),
		"hash": hash,
	}).Debug("successful set get hash")

	// JSON encode response value
	json.NewEncoder(w).Encode(hash)
}
