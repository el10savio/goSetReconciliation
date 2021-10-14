package sync

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/el10savio/goSetReconciliation/set"
)

// Send generates the sync payload and sends
// it to all the peers in the cluster
func Send(Set set.Set, missingElements []int) error {
	// Send across missing elements (no elements in Phase 1)
	// and the Set's Bloom Filer & Hash
	payload := Payload{
		MissingElements: missingElements,
		BF:              Set.BF,
		Hash:            Set.Hash,
	}

	peers := GetPeerList()
	if len(peers) == 0 {
		return errors.New("No Peers Provided")
	}

	for _, peer := range peers {
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

// SendSyncRequest sends the HTTP Sync Reconcile
// POST request to a given peer
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
