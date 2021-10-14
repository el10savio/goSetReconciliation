package set

import (
	"encoding/binary"

	"github.com/bits-and-blooms/bloom/v3"
	"github.com/mitchellh/hashstructure"
)

const (
	setSize           = 1000
	falsePositiveRate = 0.01
)

// Set ...
type Set struct {
	List []int              `hash:"set"`
	BF   *bloom.BloomFilter `hash:"ignore"`
	Hash uint64             `hash:"ignore"`
}

// Initialize ...
func Initialize() Set {
	return Set{
		List: []int{},
		BF:   bloom.NewWithEstimates(setSize, falsePositiveRate),
		Hash: uint64(0),
	}
}

// GetElements ...
func (set *Set) GetElements() []int {
	return set.List
}

// AddElements ...
func (set *Set) AddElements(elements []int) {
	for _, element := range elements {
		if !set.AddElementToBF(element) {
			set.List = append(set.List, element)
			set.Hash, _ = hashstructure.Hash(set, nil)
		}
	}
}

// AddElementToBF ...
func (set *Set) AddElementToBF(element int) bool {
	array := make([]byte, 4)
	binary.BigEndian.PutUint32(array, uint32(element))
	return set.BF.TestAndAdd(array)
}

// IsElementInBF ...
func IsElementInBF(element int, BF *bloom.BloomFilter) bool {
	array := make([]byte, 4)
	binary.BigEndian.PutUint32(array, uint32(element))
	return BF.Test(array)
}

// MergeElements ...
func MergeElements(set Set, elements []int) Set {
	set.AddElements(elements)
	return set
}

// Clear ...
func (set *Set) Clear() {
	set.List = []int{}
	set.BF.ClearAll()
	set.Hash = uint64(0)
}
