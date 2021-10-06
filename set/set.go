package set

import (
	"encoding/binary"

	"github.com/bits-and-blooms/bloom/v3"
)

const (
	setSize           = 1000
	falsePositiveRate = 0.01
)

type Set struct {
	List []uint32
	BF   *bloom.BloomFilter
	// Hash
}

func Initialize() Set {
	return Set{
		List: []uint32{},
		BF:   bloom.NewWithEstimates(setSize, falsePositiveRate),
	}
}

func (set *Set) GetElements() []uint32 {
	return set.List
}

func (set *Set) AddElement(element uint32) {
	set.List = append(set.List, element)
	set.AddElementToBF(element)
}

func (set *Set) AddElementToBF(element uint32) {
	array := make([]byte, 4)
	binary.BigEndian.PutUint32(array, element)
	set.BF.Add(array)
}
