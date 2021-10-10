package sync

import (
	"github.com/bits-and-blooms/bloom/v3"

	"github.com/el10savio/goSetReconciliation/set"
)

func GetBFMissingElements(list []uint32, BF bloom.BloomFilter) []uint32 {
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
