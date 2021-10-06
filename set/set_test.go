package set

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitialize(t *testing.T) {
	expectedSet := Set{List: []int{}}
	actualSet := Initialize()
	assert.Equal(t, expectedSet, actualSet)
}

func TestAddElement(t *testing.T) {
	elementToAdd := 1

	expectedSet := Set{List: []int{elementToAdd}}
	actualset := Initialize()
	actualset.AddElement(elementToAdd)

	assert.Equal(t, expectedSet, actualset)
}

// TODO: Change To Skip Duplicates
func TestAddElement_Duplicate(t *testing.T) {
	elementToAdd := 1

	expectedSet := Set{List: []int{elementToAdd, elementToAdd}}
	actualset := Initialize()
	actualset.AddElement(elementToAdd)
	actualset.AddElement(elementToAdd)

	assert.Equal(t, expectedSet, actualset)
}

func TestGetElements(t *testing.T) {
	set := Initialize()
	set.AddElement(1)
	set.AddElement(2)
	set.AddElement(3)

	expectedElements := []int{1, 2, 3}
	actualElements := set.GetElements()

	assert.Equal(t, expectedElements, actualElements)
}

func TestGetElements_Empty(t *testing.T) {
	set := Initialize()

	expectedElements := make([]int, 0)
	actualElements := set.GetElements()

	assert.Equal(t, expectedElements, actualElements)
}
