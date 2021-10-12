package sync

import (
	"testing"

	"github.com/el10savio/goSetReconciliation/set"
	"github.com/stretchr/testify/assert"
)

func TestGetBFMissingElements(t *testing.T) {
	list := []uint32{1}
	set := set.Initialize()
	defer set.Clear()

	set.AddElement(2)
	set.AddElement(3)

	expectedElements := []uint32{1}
	actualElements := GetBFMissingElements(list, set.BF)

	assert.Equal(t, expectedElements, actualElements)
}

func TestGetBFMissingElements_EmptyBF(t *testing.T) {
	list := []uint32{1, 2, 3}
	set := set.Initialize()
	defer set.Clear()

	expectedElements := []uint32{1, 2, 3}
	actualElements := GetBFMissingElements(list, set.BF)

	assert.Equal(t, expectedElements, actualElements)
}

func TestGetBFMissingElements_EmptyList(t *testing.T) {
	list := []uint32{}
	set := set.Initialize()
	defer set.Clear()

	set.AddElement(2)
	set.AddElement(3)

	expectedElements := []uint32{}
	actualElements := GetBFMissingElements(list, set.BF)

	assert.Equal(t, expectedElements, actualElements)
}

func TestGetBFMissingElements_BothEmpty(t *testing.T) {
	list := []uint32{}
	set := set.Initialize()
	defer set.Clear()

	expectedElements := []uint32{}
	actualElements := GetBFMissingElements(list, set.BF)

	assert.Equal(t, expectedElements, actualElements)
}
