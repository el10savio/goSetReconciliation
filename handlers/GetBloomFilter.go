package handlers

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// GetBloomFilter ...
func GetBloomFilter(w http.ResponseWriter, r *http.Request) {
	// Get the Bloom Filter from the Set
	bloomFilter := Set.BF

	// DEBUG log in the case of success
	// indicating the Bloom Filter
	log.WithFields(log.Fields{
		"list": Set.GetElements(),
		"bf":   bloomFilter,
	}).Debug("successful set get bf")

	// JSON encode response value
	json.NewEncoder(w).Encode(bloomFilter)
}
