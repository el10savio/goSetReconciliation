package sync

import (
	"fmt"

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

	if len(missingElements) > 0 {
		fmt.Println("Phase 2", missingElements)
		if err := Send(Set, missingElements); err != nil {
			fmt.Println("Phase 2 error:", err)
		}
	}

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
