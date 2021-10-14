package set

import (
	"encoding/binary"
	"testing"

	"github.com/bits-and-blooms/bloom/v3"
	"github.com/stretchr/testify/assert"
)

func TestInitialize(t *testing.T) {
	expectedSet := Set{
		List: []int{},
		BF:   bloom.NewWithEstimates(setSize, falsePositiveRate),
		Hash: uint64(0),
	}
	actualSet := Initialize()
	defer actualSet.Clear()

	assert.Equal(t, expectedSet, actualSet)
}

func TestClear(t *testing.T) {
	expectedSet := Set{
		List: []int{},
		BF:   bloom.NewWithEstimates(setSize, falsePositiveRate),
		Hash: uint64(0),
	}
	actualSet := Initialize()
	actualSet.Clear()

	assert.Equal(t, expectedSet, actualSet)
}

func TestAddElements(t *testing.T) {
	elementsToAdd := []int{1}

	expectedSet := Set{List: elementsToAdd}
	actualSet := Initialize()
	defer actualSet.Clear()

	actualSet.AddElements(elementsToAdd)

	assert.Equal(t, expectedSet.List, actualSet.List)
}

func TestAddElements_Duplicate(t *testing.T) {
	elementsToAdd := []int{1, 1}

	expectedSet := Set{List: []int{1}}
	actualSet := Initialize()
	defer actualSet.Clear()

	actualSet.AddElements(elementsToAdd)

	assert.Equal(t, expectedSet.List, actualSet.List)
}

func TestGetElements(t *testing.T) {
	set := Initialize()
	defer set.Clear()

	set.AddElements([]int{1, 2, 3})

	expectedElements := []int{1, 2, 3}
	actualElements := set.GetElements()

	assert.Equal(t, expectedElements, actualElements)
}

func TestGetElements_Empty(t *testing.T) {
	set := Initialize()
	defer set.Clear()

	expectedElements := make([]int, 0)
	actualElements := set.GetElements()

	assert.Equal(t, expectedElements, actualElements)
}

func TestIsElementInBF(t *testing.T) {
	element := 1

	BF := bloom.NewWithEstimates(setSize, falsePositiveRate)
	BF = addElementToBF(element, BF)
	defer BF.ClearAll()

	expectedCondition := true
	actualCondition := IsElementInBF(element, BF)

	assert.Equal(t, expectedCondition, actualCondition)
}

func TestIsElementInBF_NotPresent(t *testing.T) {
	element := 2

	BF := bloom.NewWithEstimates(setSize, falsePositiveRate)
	BF = addElementToBF(int(3), BF)
	defer BF.ClearAll()

	expectedCondition := false
	actualCondition := IsElementInBF(element, BF)

	assert.Equal(t, expectedCondition, actualCondition)
}

func TestIsElementInBF_EmptyBF(t *testing.T) {
	element := 1

	BF := bloom.NewWithEstimates(setSize, falsePositiveRate)
	defer BF.ClearAll()

	expectedCondition := false
	actualCondition := IsElementInBF(element, BF)

	assert.Equal(t, expectedCondition, actualCondition)
}

func TestMergeElements(t *testing.T) {
	elements, elementsToMerge := []int{1, 2}, []int{3, 4, 5}
	elementsMerged := []int{1, 2, 3, 4, 5}

	expectedSet := Set{List: elementsMerged}
	actualSet := Initialize()
	defer actualSet.Clear()

	actualSet.AddElements(elements)
	actualSet = MergeElements(actualSet, elementsToMerge)

	assert.Equal(t, expectedSet.List, actualSet.List)
}

func TestMergeElements_Empty(t *testing.T) {
	elements, elementsToMerge := []int{1, 2}, []int{}
	elementsMerged := []int{1, 2}

	expectedSet := Set{List: elementsMerged}
	actualSet := Initialize()
	defer actualSet.Clear()

	actualSet.AddElements(elements)
	actualSet = MergeElements(actualSet, elementsToMerge)

	assert.Equal(t, expectedSet.List, actualSet.List)
}

func TestMergeElements_BothEmpty(t *testing.T) {
	elements, elementsToMerge := []int{}, []int{}
	elementsMerged := []int{}

	expectedSet := Set{List: elementsMerged}
	actualSet := Initialize()
	defer actualSet.Clear()

	actualSet.AddElements(elements)
	actualSet = MergeElements(actualSet, elementsToMerge)

	assert.Equal(t, expectedSet.List, actualSet.List)
}

func TestMergeElements_Duplicate(t *testing.T) {
	elements, elementsToMerge := []int{1, 2}, []int{1}
	elementsMerged := []int{1, 2}

	expectedSet := Set{List: elementsMerged}
	actualSet := Initialize()
	defer actualSet.Clear()

	actualSet.AddElements(elements)
	actualSet = MergeElements(actualSet, elementsToMerge)

	assert.Equal(t, expectedSet.List, actualSet.List)
}

func addElementToBF(element int, BF *bloom.BloomFilter) *bloom.BloomFilter {
	array := make([]byte, 4)
	binary.BigEndian.PutUint32(array, uint32(element))
	return BF.Add(array)
}
