package set

import (
	"encoding/binary"

	"github.com/bits-and-blooms/bloom/v3"
	"github.com/mitchellh/hashstructure"
)

const (
	// Bloom Filter set size
	setSize = 1000

	// Bloom Filter False Positive Rate
	falsePositiveRate = 0.01
)

// Set is used to define our list of unique numbers
// We can only add elements to it & clear it for tests
// Every time a new element is added the Bloom Filter
// and corresponding Hash gets updated
type Set struct {
	List []int              `hash:"set"`    // Collection of unique integers
	BF   *bloom.BloomFilter `hash:"ignore"` // State of elements of the list
	Hash uint64             `hash:"ignore"` // Unique Hash for the list state
}

// Initialize creates a new set
// with an empty list, Bloom Filter
// and a default Hash of 0
func Initialize() Set {
	return Set{
		List: []int{},
		BF:   initBloomFilter(),
		Hash: uint64(0),
	}
}

// initBloomFilter is a utility function to return
// an empty Bloom Filter with our specified defaults
func initBloomFilter() *bloom.BloomFilter {
	return bloom.NewWithEstimates(setSize, falsePositiveRate)
}

// addElementToBF is a utitily function to
// add an element into the Bloom Filter
func addElementToBF(element int, BF *bloom.BloomFilter) *bloom.BloomFilter {
	array := make([]byte, 4)
	binary.BigEndian.PutUint32(array, uint32(element))
	return BF.Add(array)
}

// GetElements returns the elements
// present in the Set
func (set *Set) GetElements() []int {
	return set.List
}

// AddElements inserts new elements to the list
// only if it was  never present before &
// updates the Bloom Filter & Hash
func (set *Set) AddElements(elements []int) {
	for _, element := range elements {
		if !set.AddElementToBF(element) {
			set.List = append(set.List, element)
			set.Hash, _ = hashstructure.Hash(set, nil)
		}
	}
}

// AddElementToBF updates the Bloom Filter with the
// new element by converting it into binary
// and adds it in if it is unique
func (set *Set) AddElementToBF(element int) bool {
	array := make([]byte, 4)
	binary.BigEndian.PutUint32(array, uint32(element))
	return set.BF.TestAndAdd(array)
}

// Clear resets the set
// and its parameters
func (set *Set) Clear() {
	set.List = []int{}
	set.BF.ClearAll()
	set.Hash = uint64(0)
}

// IsElementInBF checks if the given element
// is present in the bloom filter
func IsElementInBF(element int, BF *bloom.BloomFilter) bool {
	array := make([]byte, 4)
	binary.BigEndian.PutUint32(array, uint32(element))
	return BF.Test(array)
}

// MergeElements takes a list of elements
// to be added into & a Set and returns the
// Set merged with the given elements
func (set *Set) MergeElements(elements []int) {
	set.AddElements(elements)
}
