package sync

import (
	"github.com/bits-and-blooms/bloom/v3"
	"github.com/el10savio/goSetReconciliation/set"
)

// Update takes in the sync payload and updates the Set
// if there are any missing elements and gets
// the Set's missing elements based on
// the payload's Bloom Filter
func Update(Set set.Set, payload Payload) (set.Set, []int) {
	if Set.Hash == payload.Hash {
		return Set, []int{}
	}

	missingElements := GetBFMissingElements(Set.GetElements(), payload.BF)
	Set.MergeElements(payload.MissingElements)

	return Set, missingElements
}

// GetBFMissingElements iterates over the Set
// and gets any elements that are not present
// in the provided Bloom Filter
func GetBFMissingElements(list []int, BF *bloom.BloomFilter) []int {
	if len(list) == 0 {
		return []int{}
	}

	missing := make([]int, 0)

	for _, element := range list {
		if set.IsElementInBF(element, BF) == false {
			missing = append(missing, element)
		}
	}

	return missing
}
