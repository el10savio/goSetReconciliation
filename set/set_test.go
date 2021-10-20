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

// TestAddElements checks the functionality
// of Set AddElements()
func TestAddElements(t *testing.T) {
	for _, testCase := range testAddElementsTestSuite {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			actualSet := Initialize()
			defer actualSet.Clear()

			expectedSet := Set{List: testCase.expectedElements}
			actualSet.AddElements(testCase.elementsToAdd)

			assert.Equal(t, expectedSet.List, actualSet.List)
		})
	}
}

// TestGetElements checks the functionality
// of Set GetElements()
func TestGetElements(t *testing.T) {
	for _, testCase := range testGetElementsTestSuite {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			set := Initialize()
			defer set.Clear()

			set.AddElements(testCase.inputElements)

			expectedElements := testCase.expectedElements
			actualElements := set.GetElements()

			assert.Equal(t, expectedElements, actualElements)
		})
	}
}

// TestIsElementInBF checks the functionality
// of IsElementInBF()
func TestIsElementInBF(t *testing.T) {
	for _, testCase := range testIsElementInBFTestSuite {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			BF := initBloomFilter()
			defer BF.ClearAll()

			for _, element := range testCase.elementsToBeAdded {
				BF = addElementToBF(element, BF)
			}

			expectedCondition := testCase.expectedCondition
			actualCondition := IsElementInBF(testCase.elementToCheck, BF)

			assert.Equal(t, expectedCondition, actualCondition)
		})
	}
}

// TestMergeElements checks the functionality
// of Set MergeElements()
func TestMergeElements(t *testing.T) {
	for _, testCase := range testMergeElementsTestSuite {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			set := Initialize()
			defer set.Clear()

			set.AddElements(testCase.elementsBase)
			set.MergeElements(testCase.elementsToMerge)

			expectedSet := Set{List: testCase.expectedElementsMerged}

			assert.Equal(t, expectedSet.List, set.List)
		})
	}
}
