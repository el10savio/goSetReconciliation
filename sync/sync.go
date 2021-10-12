package sync

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/el10savio/goSetReconciliation/set"
)

// Send ...
func Send(Set set.Set, missingElements []uint32) error {
	// Send BF & Hash
	payload := Payload{
		MissingElements: missingElements,
		BF:              Set.GetBF(),
		Hash:            Set.GetHash(),
	}

	for _, peer := range GetPeerList() {
		if peer == GetHost() {
			continue
		}
		_, err := SendSyncRequest(peer, payload)
		if err != nil {
			return err
		}
	}

	return nil
}

// SendSyncRequest sends the HTTP Sync POST request to a given peer
func SendSyncRequest(peer string, payload Payload) (int, error) {
	if peer == "" {
		return 0, errors.New("empty peer provided")
	}

	url := fmt.Sprintf("http://%s.%s/set/sync/reconcile", peer, GetNetwork())
	JSONPayload, err := json.Marshal(payload)
	if err != nil {
		return 0, err
	}

	return SendRequest(url, JSONPayload)
}
