package sync

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/bits-and-blooms/bloom/v3"

	"github.com/el10savio/goSetReconciliation/set"
)

func Send(Set set.Set) error {
	// Send BF & Hash
	payload := Payload{
		MissingElements: []uint16{},
		BF:              Set.GetBF(),
		Hash:            Set.GetHash(),
	}

	fmt.Println(payload)
	return nil
}

// SendSyncRequest sends the HTTP Sync POST request to a given peer
func SendSyncRequest(peer string, payload Payload) (int, error) {
	if peer == "" {
		return 0, errors.New("empty peer provided")
	}

	url := fmt.Sprintf("http://%s.%s/sync/reconcile", peer, GetNetwork())
	json.Marshal(payload)

	return SendRequest(url, payload)
}

func GetBFMissingElements(list []uint32, BF *bloom.BloomFilter) []uint32 {
	if list == nil {
		return []uint32{}
	}

	missing := make([]uint32, 0)

	for _, element := range list {
		if set.IsElementInBF(element, BF) == false {
			missing = append(missing, element)
		}
	}

	return missing
}
