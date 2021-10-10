package set

import (
	"encoding/binary"
	"testing"

	"github.com/bits-and-blooms/bloom/v3"
	"github.com/stretchr/testify/assert"
)

func TestInitialize(t *testing.T) {
	expectedSet := Set{
		List: []uint32{},
		BF:   bloom.NewWithEstimates(setSize, falsePositiveRate),
		Hash: uint64(0),
	}
	actualSet := Initialize()
	defer actualSet.Clear()

	assert.Equal(t, expectedSet, actualSet)
}

func TestClear(t *testing.T) {
	expectedSet := Set{
		List: []uint32{},
		BF:   bloom.NewWithEstimates(setSize, falsePositiveRate),
		Hash: uint64(0),
	}
	actualSet := Initialize()
	actualSet.Clear()

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

func TestIsElementInBF(t *testing.T) {
	element := uint32(1)
	BF := bloom.NewWithEstimates(setSize, falsePositiveRate)
	BF = addElementToBF(element, BF)
	defer BF.ClearAll()

	expectedCondition := true
	actualCondition := IsElementInBF(element, BF)

	assert.Equal(t, expectedCondition, actualCondition)
}

func TestIsElementInBF_NotPresent(t *testing.T) {
	element := uint32(2)
	BF := bloom.NewWithEstimates(setSize, falsePositiveRate)
	BF = addElementToBF(uint32(3), BF)
	defer BF.ClearAll()

	expectedCondition := false
	actualCondition := IsElementInBF(element, BF)

	assert.Equal(t, expectedCondition, actualCondition)
}

func TestIsElementInBF_EmptyBF(t *testing.T) {
	element := uint32(1)
	BF := bloom.NewWithEstimates(setSize, falsePositiveRate)
	defer BF.ClearAll()

	expectedCondition := false
	actualCondition := IsElementInBF(element, BF)

	assert.Equal(t, expectedCondition, actualCondition)
}

func TestMergeElements(t *testing.T) {
	elements, elementsToMerge := []uint32{1, 2}, []uint32{3, 4, 5}
	elementsMerged := []uint32{1, 2, 3, 4, 5}

	expectedSet := Set{List: elementsMerged}
	actualSet := Initialize()
	defer actualSet.Clear()

	for _, element := range elements {
		actualSet.AddElement(element)
	}
	actualSet.MergeElements(elementsToMerge)

	assert.Equal(t, expectedSet.List, actualSet.List)
}

func TestMergeElements_Empty(t *testing.T) {
	elements, elementsToMerge := []uint32{1, 2}, []uint32{}
	elementsMerged := []uint32{1, 2}

	expectedSet := Set{List: elementsMerged}
	actualSet := Initialize()
	defer actualSet.Clear()

	for _, element := range elements {
		actualSet.AddElement(element)
	}
	actualSet.MergeElements(elementsToMerge)

	assert.Equal(t, expectedSet.List, actualSet.List)
}

func TestMergeElements_BothEmpty(t *testing.T) {
	elements, elementsToMerge := []uint32{}, []uint32{}
	elementsMerged := []uint32{}

	expectedSet := Set{List: elementsMerged}
	actualSet := Initialize()
	defer actualSet.Clear()

	for _, element := range elements {
		actualSet.AddElement(element)
	}
	actualSet.MergeElements(elementsToMerge)

	assert.Equal(t, expectedSet.List, actualSet.List)
}

func TestMergeElements_Duplicate(t *testing.T) {
	elements := []uint32{1, 2}
	elementsToMerge := []uint32{1}
	elementsMerged := []uint32{1, 2}

	expectedSet := Set{List: elementsMerged}
	actualSet := Initialize()
	defer actualSet.Clear()

	for _, element := range elements {
		actualSet.AddElement(element)
	}
	actualSet.MergeElements(elementsToMerge)

	assert.Equal(t, expectedSet.List, actualSet.List)
}

func addElementToBF(element uint32, BF *bloom.BloomFilter) *bloom.BloomFilter {
	array := make([]byte, 4)
	binary.BigEndian.PutUint32(array, element)
	return BF.Add(array)
}
