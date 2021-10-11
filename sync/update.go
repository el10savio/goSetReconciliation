package sync

import (
	"fmt"

	"github.com/bits-and-blooms/bloom/v3"
	"github.com/el10savio/goSetReconciliation/set"
)

// Update ...
func Update(Set set.Set, payload Payload) set.Set {
	if Set.GetHash() == payload.Hash {
		return Set
	}

	Set = *Set.MergeElements(payload.MissingElements)
	missingElements := GetBFMissingElements(Set.GetElements(), payload.BF)

	fmt.Println(missingElements)
	// _ = Send(Set, missingElements)

	return Set
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
