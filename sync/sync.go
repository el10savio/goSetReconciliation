package sync

import (
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
