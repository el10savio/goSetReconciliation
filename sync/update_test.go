package sync

import (
	"testing"

	"github.com/el10savio/goSetReconciliation/set"
	"github.com/stretchr/testify/assert"
)

// TestGetBFMissingElements tests the basic functionality
// of GetBFMissingElements()
func TestGetBFMissingElements(t *testing.T) {
	list := []int{1}
	set := set.Initialize()
	defer set.Clear()

	set.AddElements([]int{2, 3})

	expectedElements := []int{1}
	actualElements := GetBFMissingElements(list, set.BF)

	assert.Equal(t, expectedElements, actualElements)
}

// TestGetBFMissingElements_EmptyBF tests the functionality of GetBFMissingElements()
// When the Set's Bloom Filter is Empty
func TestGetBFMissingElements_EmptyBF(t *testing.T) {
	list := []int{1, 2, 3}
	set := set.Initialize()
	defer set.Clear()

	expectedElements := []int{1, 2, 3}
	actualElements := GetBFMissingElements(list, set.BF)

	assert.Equal(t, expectedElements, actualElements)
}

// TestGetBFMissingElements_EmptyList tests the functionality of GetBFMissingElements()
// When the List is Empty
func TestGetBFMissingElements_EmptyList(t *testing.T) {
	list := []int{}
	set := set.Initialize()
	defer set.Clear()

	set.AddElements([]int{2, 3})

	expectedElements := []int{}
	actualElements := GetBFMissingElements(list, set.BF)

	assert.Equal(t, expectedElements, actualElements)
}

// TestGetBFMissingElements_BothEmpty tests the functionality of GetBFMissingElements()
// When both the Set's Bloom Filter & the List are Empty
func TestGetBFMissingElements_BothEmpty(t *testing.T) {
	list := []int{}
	set := set.Initialize()
	defer set.Clear()

	expectedElements := []int{}
	actualElements := GetBFMissingElements(list, set.BF)

	assert.Equal(t, expectedElements, actualElements)
}

// TestUpdate tests the basic functionality
// of Update()
func TestUpdate(t *testing.T) {
	baseSet := set.Initialize()
	defer baseSet.Clear()

	payloadSet := set.Initialize()
	defer payloadSet.Clear()

	baseSet.AddElements([]int{1, 2})
	payloadSet.AddElements([]int{1, 2, 3})

	payload := Payload{
		MissingElements: []int{3},
		BF:              payloadSet.BF,
		Hash:            payloadSet.Hash,
	}

	expectedMissingElements := []int{}
	baseSet, actualMissingElements := Update(baseSet, payload)

	expectedSet := set.Initialize()
	defer expectedSet.Clear()

	expectedSet.AddElements([]int{1, 2, 3})

	assert.Equal(t, expectedSet, baseSet)
	assert.Equal(t, expectedMissingElements, actualMissingElements)
}

// TestUpdate_BothEqual tests the functionality of Update()
// When both Sets are equal
func TestUpdate_BothEqual(t *testing.T) {
	baseSet := set.Initialize()
	defer baseSet.Clear()

	payloadSet := set.Initialize()
	defer payloadSet.Clear()

	baseSet.AddElements([]int{1, 2, 3})
	payloadSet.AddElements([]int{1, 2, 3})

	payload := Payload{
		MissingElements: []int{},
		BF:              payloadSet.BF,
		Hash:            payloadSet.Hash,
	}

	expectedMissingElements := []int{}
	baseSet, actualMissingElements := Update(baseSet, payload)

	expectedSet := set.Initialize()
	defer expectedSet.Clear()

	expectedSet.AddElements([]int{1, 2, 3})

	assert.Equal(t, expectedSet, baseSet)
	assert.Equal(t, expectedMissingElements, actualMissingElements)
}

// TestUpdate_EmptySet tests the functionality of Update()
// When the base Set is empty
func TestUpdate_EmptySet(t *testing.T) {
	baseSet := set.Initialize()
	defer baseSet.Clear()

	payloadSet := set.Initialize()
	defer payloadSet.Clear()

	payloadSet.AddElements([]int{1, 2, 3})

	payload := Payload{
		MissingElements: []int{1, 2, 3},
		BF:              payloadSet.BF,
		Hash:            payloadSet.Hash,
	}

	expectedMissingElements := []int{}
	baseSet, actualMissingElements := Update(baseSet, payload)

	expectedSet := set.Initialize()
	defer expectedSet.Clear()

	expectedSet.AddElements([]int{1, 2, 3})

	assert.Equal(t, expectedSet, baseSet)
	assert.Equal(t, expectedMissingElements, actualMissingElements)
}

// TestUpdate_EmptyPayload tests the functionality of Update()
// When the payload Set is empty
func TestUpdate_EmptyPayload(t *testing.T) {
	baseSet := set.Initialize()
	defer baseSet.Clear()

	baseSet.AddElements([]int{1, 2})

	payloadSet := set.Initialize()
	defer payloadSet.Clear()

	payload := Payload{
		MissingElements: []int{},
		BF:              payloadSet.BF,
		Hash:            payloadSet.Hash,
	}

	expectedMissingElements := []int{1, 2}
	baseSet, actualMissingElements := Update(baseSet, payload)

	expectedSet := set.Initialize()
	defer expectedSet.Clear()

	expectedSet.AddElements([]int{1, 2})

	assert.Equal(t, expectedSet, baseSet)
	assert.Equal(t, expectedMissingElements, actualMissingElements)
}

// TestUpdate_BothEmpty tests the functionality of Update()
// When both the base Set & payload Set are empty
func TestUpdate_BothEmpty(t *testing.T) {
	baseSet := set.Initialize()
	defer baseSet.Clear()

	payloadSet := set.Initialize()
	defer payloadSet.Clear()

	payload := Payload{
		MissingElements: []int{},
		BF:              payloadSet.BF,
		Hash:            payloadSet.Hash,
	}

	expectedMissingElements := []int{}
	baseSet, actualMissingElements := Update(baseSet, payload)

	expectedSet := set.Initialize()
	defer expectedSet.Clear()

	assert.Equal(t, expectedSet, baseSet)
	assert.Equal(t, expectedMissingElements, actualMissingElements)
}

// TestUpdate_FullSync tests the functionality of Update()
// When the base Set in one node has elements & the Set in the other node is empty
// and the sync is initiated from the node that has elements
func TestUpdate_FullSync(t *testing.T) {
	baseSet := set.Initialize()
	defer baseSet.Clear()

	baseSet.AddElements([]int{1, 2})

	payloadSet := set.Initialize()
	defer payloadSet.Clear()

	payload := Payload{
		MissingElements: []int{},
		BF:              payloadSet.BF,
		Hash:            payloadSet.Hash,
	}

	baseSet, actualMissingElements := Update(baseSet, payload)

	payload = Payload{
		MissingElements: actualMissingElements,
		BF:              baseSet.BF,
		Hash:            baseSet.Hash,
	}

	expectedMissingElements := []int{}
	payloadSet, actualMissingElements = Update(payloadSet, payload)

	expectedSet := set.Initialize()
	defer expectedSet.Clear()

	expectedSet.AddElements([]int{1, 2})

	assert.Equal(t, expectedSet, baseSet)
	assert.Equal(t, baseSet, payloadSet)
	assert.Equal(t, expectedMissingElements, actualMissingElements)
}

// TestUpdate_FullSyncOtherNode tests the functionality of Update()
// When the base Set in one node has elements & the Set in the other node is empty
// and the sync is initiated from the node that is empty
func TestUpdate_FullSyncOtherNode(t *testing.T) {
	baseSet := set.Initialize()
	defer baseSet.Clear()

	payloadSet := set.Initialize()
	defer payloadSet.Clear()

	payloadSet.AddElements([]int{1, 2})

	payload := Payload{
		MissingElements: []int{},
		BF:              payloadSet.BF,
		Hash:            payloadSet.Hash,
	}

	expectedMissingElements := []int{}
	baseSet, actualMissingElements := Update(baseSet, payload)

	assert.Equal(t, expectedMissingElements, actualMissingElements)

	// baseSet Unchanged

	payload = Payload{
		MissingElements: actualMissingElements,
		BF:              baseSet.BF,
		Hash:            baseSet.Hash,
	}

	expectedMissingElements = []int{1, 2}
	payloadSet, actualMissingElements = Update(payloadSet, payload)

	assert.Equal(t, expectedMissingElements, actualMissingElements)

	// payloadSet Unchanged

	payload = Payload{
		MissingElements: actualMissingElements,
		BF:              payloadSet.BF,
		Hash:            payloadSet.Hash,
	}

	expectedMissingElements = []int{}
	baseSet, actualMissingElements = Update(baseSet, payload)

	// baseSet Changed

	assert.Equal(t, expectedMissingElements, actualMissingElements)
	assert.Equal(t, payloadSet, baseSet)
}

// TestUpdate_MixedSync tests the functionality of Update()
// When the base Set in one node has elements
// different from the Set in the other node
func TestUpdate_MixedSync(t *testing.T) {
	baseSet := set.Initialize()
	defer baseSet.Clear()

	baseSet.AddElements([]int{1, 2, 6})

	payloadSet := set.Initialize()
	defer payloadSet.Clear()

	baseSet.AddElements([]int{1, 2, 3, 4, 5})

	payload := Payload{
		MissingElements: []int{},
		BF:              payloadSet.BF,
		Hash:            payloadSet.Hash,
	}

	expectedMissingElements := []int{1, 2, 6, 3, 4, 5}
	baseSet, actualMissingElements := Update(baseSet, payload)

	expectedSet := set.Initialize()
	defer expectedSet.Clear()

	expectedSet.AddElements([]int{1, 2, 6, 3, 4, 5})

	assert.Equal(t, expectedSet, baseSet)
	assert.Equal(t, expectedMissingElements, actualMissingElements)

	payload = Payload{
		MissingElements: actualMissingElements,
		BF:              baseSet.BF,
		Hash:            baseSet.Hash,
	}

	expectedMissingElements = []int{}
	payloadSet, actualMissingElements = Update(payloadSet, payload)

	assert.Equal(t, expectedSet, payloadSet)
	assert.Equal(t, expectedMissingElements, actualMissingElements)
}
