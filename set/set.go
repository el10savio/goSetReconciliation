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

type Set struct {
	List []uint32           `hash:"set"`
	BF   *bloom.BloomFilter `hash:"ignore"`
	Hash uint64             `hash:"ignore"`
}

func Initialize() Set {
	return Set{
		List: []uint32{},
		BF:   bloom.NewWithEstimates(setSize, falsePositiveRate),
		Hash: uint64(0),
	}
}

func (set *Set) GetElements() []uint32 {
	return set.List
}

// TODO: Handle Hash Error
func (set *Set) AddElement(element uint32) {
	if !set.AddElementToBF(element) {
		set.List = append(set.List, element)
		set.Hash, _ = hashstructure.Hash(set, nil)
	}
}

func (set *Set) AddElementToBF(element uint32) bool {
	array := make([]byte, 4)
	binary.BigEndian.PutUint32(array, element)
	return set.BF.TestAndAdd(array)
}

func (set *Set) GetBF() *bloom.BloomFilter {
	return set.BF
}

func (set *Set) GetHash() uint64 {
	return set.Hash
}

func (set *Set) Clear() {
	set.List = []uint32{}
	set.BF.ClearAll()
	set.Hash = uint64(0)
}
