package sync

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/bits-and-blooms/bloom/v3"

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

	fmt.Println(payload)
	fmt.Println(GetPeerList())

	for _, peer := range GetPeerList() {
		_, err := SendSyncRequest(peer, payload)
		if err != nil {
			return err
		}
	}

	return nil
}

// Update ...
func Update(Set set.Set, payload Payload) set.Set {
	if Set.GetHash() == payload.Hash {
		return Set
	}

	Set = *Set.MergeElements(payload.MissingElements)
	missingElements := GetBFMissingElements(Set.GetElements(), payload.BF)

	fmt.Println(Set, payload.BF, payload.Hash, missingElements)
	_ = Send(Set, missingElements)

	return Set
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

// GetBFMissingElements ...
func GetBFMissingElements(list []uint32, BF *bloom.BloomFilter) []uint32 {
	missing := make([]uint32, 0)

	for _, element := range list {
		if set.IsElementInBF(element, BF) == false {
			missing = append(missing, element)
		}
	}

	return missing
}
