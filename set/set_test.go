package set

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestInitialize checks the basic functionality
// of Set Initialize()
func TestInitialize(t *testing.T) {
	expectedSet := Set{
		List: []int{},
		BF:   initBloomFilter(),
		Hash: uint64(0),
	}
	actualSet := Initialize()
	defer actualSet.Clear()

	assert.Equal(t, expectedSet, actualSet)
}

// TestClear checks the basic functionality
// of Set Clear()
func TestClear(t *testing.T) {
	expectedSet := Set{
		List: []int{},
		BF:   initBloomFilter(),
		Hash: uint64(0),
	}
	actualSet := Initialize()
	actualSet.Clear()

	assert.Equal(t, expectedSet, actualSet)
}

// TestAddElements checks the basic functionality
// of Set AddElements()
func TestAddElements(t *testing.T) {
	elementsToAdd := []int{1}

	expectedSet := Set{List: elementsToAdd}
	actualSet := Initialize()
	defer actualSet.Clear()

	actualSet.AddElements(elementsToAdd)

	assert.Equal(t, expectedSet.List, actualSet.List)
}

// TestAddElements_Duplicate checks the functionality of AddElements()
// When duplicate elements are provided to be added
func TestAddElements_Duplicate(t *testing.T) {
	elementsToAdd := []int{1, 1}

	expectedSet := Set{List: []int{1}}
	actualSet := Initialize()
	defer actualSet.Clear()

	actualSet.AddElements(elementsToAdd)

	assert.Equal(t, expectedSet.List, actualSet.List)
}

// TestGetElements checks the basic functionality
// of Set GetElements()
func TestGetElements(t *testing.T) {
	set := Initialize()
	defer set.Clear()

	set.AddElements([]int{1, 2, 3})

	expectedElements := []int{1, 2, 3}
	actualElements := set.GetElements()

	assert.Equal(t, expectedElements, actualElements)
}

// TestGetElements_Empty checks the functionality of GetElements()
// When the Set contains no elements
func TestGetElements_Empty(t *testing.T) {
	set := Initialize()
	defer set.Clear()

	expectedElements := make([]int, 0)
	actualElements := set.GetElements()

	assert.Equal(t, expectedElements, actualElements)
}

// TestIsElementInBF checks the basic functionality
// of IsElementInBF()
func TestIsElementInBF(t *testing.T) {
	element := 1

	BF := initBloomFilter()
	BF = addElementToBF(element, BF)
	defer BF.ClearAll()

	expectedCondition := true
	actualCondition := IsElementInBF(element, BF)

	assert.Equal(t, expectedCondition, actualCondition)
}

// TestIsElementInBF_NotPresent checks the functionality of IsElementInBF()
// When the element is not present in the Bloom Filter
func TestIsElementInBF_NotPresent(t *testing.T) {
	element := 2

	BF := initBloomFilter()
	BF = addElementToBF(int(3), BF)
	defer BF.ClearAll()

	expectedCondition := false
	actualCondition := IsElementInBF(element, BF)

	assert.Equal(t, expectedCondition, actualCondition)
}

// TestIsElementInBF_EmptyBF checks the functionality of IsElementInBF()
// When the Bloom Filter is empty
func TestIsElementInBF_EmptyBF(t *testing.T) {
	element := 1

	BF := initBloomFilter()
	defer BF.ClearAll()

	expectedCondition := false
	actualCondition := IsElementInBF(element, BF)

	assert.Equal(t, expectedCondition, actualCondition)
}

// TestMergeElements checks the basic functionality
// of Set MergeElements()
func TestMergeElements(t *testing.T) {
	elements, elementsToMerge := []int{1, 2}, []int{3, 4, 5}
	elementsMerged := []int{1, 2, 3, 4, 5}

	expectedSet := Set{List: elementsMerged}
	actualSet := Initialize()
	defer actualSet.Clear()

	actualSet.AddElements(elements)
	actualSet.MergeElements(elementsToMerge)

	assert.Equal(t, expectedSet.List, actualSet.List)
}

// TestMergeElements_Empty checks the functionality of MergeElements()
// When the list to merge is empty
func TestMergeElements_Empty(t *testing.T) {
	elements, elementsToMerge := []int{1, 2}, []int{}
	elementsMerged := []int{1, 2}

	expectedSet := Set{List: elementsMerged}
	actualSet := Initialize()
	defer actualSet.Clear()

	actualSet.AddElements(elements)
	actualSet.MergeElements(elementsToMerge)

	assert.Equal(t, expectedSet.List, actualSet.List)
}

// TestMergeElements_BothEmpty checks the functionality of MergeElements()
// When both the list to merge & the base Set are empty
func TestMergeElements_BothEmpty(t *testing.T) {
	elements, elementsToMerge := []int{}, []int{}
	elementsMerged := []int{}

	expectedSet := Set{List: elementsMerged}
	actualSet := Initialize()
	defer actualSet.Clear()

	actualSet.AddElements(elements)
	actualSet.MergeElements(elementsToMerge)

	assert.Equal(t, expectedSet.List, actualSet.List)
}

// TestMergeElements_Duplicate checks the functionality of MergeElements()
// When both the list to merge contains elements same as the base Set
func TestMergeElements_Duplicate(t *testing.T) {
	elements, elementsToMerge := []int{1, 2}, []int{1}
	elementsMerged := []int{1, 2}

	expectedSet := Set{List: elementsMerged}
	actualSet := Initialize()
	defer actualSet.Clear()

	actualSet.AddElements(elements)
	actualSet.MergeElements(elementsToMerge)

	assert.Equal(t, expectedSet.List, actualSet.List)
}
