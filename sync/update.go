package sync

import (
	"github.com/bits-and-blooms/bloom/v3"
	"github.com/el10savio/goSetReconciliation/set"
)

// Update ...
func Update(Set set.Set, payload Payload) (set.Set, []uint32) {
	if Set.GetHash() == payload.Hash {
		return Set, []uint32{}
	}

	missingElements := GetBFMissingElements(Set.GetElements(), payload.BF)
	Set = set.MergeElements(Set, payload.MissingElements)

	if len(missingElements) > 0 {
		Send(Set, missingElements)
	}

	return Set, missingElements
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
