package set

import (
	"testing"

	"github.com/bits-and-blooms/bloom/v3"
	"github.com/stretchr/testify/assert"
)

func TestInitialize(t *testing.T) {
	expectedSet := Set{
		List: []uint32{},
		BF:   bloom.NewWithEstimates(setSize, falsePositiveRate),
	}
	actualSet := Initialize()
	defer actualSet.Clear()

	assert.Equal(t, expectedSet, actualSet)
}

func TestAddElement(t *testing.T) {
	elementToAdd := uint32(1)

	expectedSet := Set{List: []uint32{elementToAdd}}
	actualSet := Initialize()
	defer actualSet.Clear()

	actualSet.AddElement(elementToAdd)

	assert.Equal(t, expectedSet.List, actualSet.List)
}

// TODO: Change To Skip Duplicates
func TestAddElement_Duplicate(t *testing.T) {
	elementToAdd := uint32(1)

	expectedSet := Set{List: []uint32{elementToAdd}}
	actualSet := Initialize()
	defer actualSet.Clear()

	actualSet.AddElement(elementToAdd)
	actualSet.AddElement(elementToAdd)

	assert.Equal(t, expectedSet.List, actualSet.List)
}

func TestGetElements(t *testing.T) {
	set := Initialize()
	defer set.Clear()

	set.AddElement(1)
	set.AddElement(2)
	set.AddElement(3)

	expectedElements := []uint32{1, 2, 3}
	actualElements := set.GetElements()

	assert.Equal(t, expectedElements, actualElements)
}

func TestGetElements_Empty(t *testing.T) {
	set := Initialize()
	defer set.Clear()

	expectedElements := make([]uint32, 0)
	actualElements := set.GetElements()

	assert.Equal(t, expectedElements, actualElements)
}
