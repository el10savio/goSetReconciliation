package handlers

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// List ...
func List(w http.ResponseWriter, r *http.Request) {
	// Get the values from the Set
	list := Set.GetElements()

	// DEBUG log in the case of success
	// indicating the new Set
	log.WithFields(log.Fields{
		"list": list,
	}).Debug("successful set list")

	// JSON encode response value
	json.NewEncoder(w).Encode(list)
}
