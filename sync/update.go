package sync

import (
	"github.com/bits-and-blooms/bloom/v3"
	"github.com/el10savio/goSetReconciliation/set"
)

// Update ...
func Update(Set set.Set, payload Payload) (set.Set, []int) {
	if Set.GetHash() == payload.Hash {
		return Set, []int{}
	}

	missingElements := GetBFMissingElements(Set.GetElements(), payload.BF)
	Set = set.MergeElements(Set, payload.MissingElements)

	return Set, missingElements
}

// GetBFMissingElements ...
func GetBFMissingElements(list []int, BF *bloom.BloomFilter) []int {
	missing := make([]int, 0)

	for _, element := range list {
		if set.IsElementInBF(element, BF) == false {
			missing = append(missing, element)
		}
	}

	return missing
}
