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
	testSuite := []struct {
		name             string
		elementsToAdd    []int
		expectedElements []int
	}{
		{name: "BasicFuntionality", elementsToAdd: []int{1, 2, 3}, expectedElements: []int{1, 2, 3}},
		{name: "Duplicate", elementsToAdd: []int{1, 1}, expectedElements: []int{1}},
	}

	for _, testCase := range testSuite {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			actualSet := Initialize()
			defer actualSet.Clear()

			actualSet.AddElements(testCase.elementsToAdd)

			expectedSet := Set{List: testCase.expectedElements}

			assert.Equal(t, expectedSet.List, actualSet.List)
		})
	}
}

// TestGetElements checks the functionality
// of Set GetElements()
func TestGetElements(t *testing.T) {
	testSuite := []struct {
		name             string
		inputElements    []int
		expectedElements []int
	}{
		{name: "BasicFuntionality", inputElements: []int{1, 2, 3}, expectedElements: []int{1, 2, 3}},
		{name: "Empty", inputElements: []int{}, expectedElements: []int{}},
	}

	for _, testCase := range testSuite {
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
	testSuite := []struct {
		name              string
		elementToCheck    int
		elementsToBeAdded []int
		expectedCondition bool
	}{
		{name: "BasicFuntionality", elementToCheck: 1, elementsToBeAdded: []int{1}, expectedCondition: true},
		{name: "NotPresent", elementToCheck: 2, elementsToBeAdded: []int{1}, expectedCondition: false},
		{name: "EmptyBloomFilter", elementToCheck: 1, elementsToBeAdded: []int{}, expectedCondition: false},
		{name: "MultipleDuplicateElements", elementToCheck: 1, elementsToBeAdded: []int{1, 1, 2, 3}, expectedCondition: true},
	}

	for _, testCase := range testSuite {
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
	testSuite := []struct {
		name                   string
		elementsBase           []int
		elementsToMerge        []int
		expectedElementsMerged []int
	}{
		{name: "BasicFuntionality", elementsBase: []int{1, 2}, elementsToMerge: []int{3, 4, 5}, expectedElementsMerged: []int{1, 2, 3, 4, 5}},
		{name: "Empty", elementsBase: []int{1, 2}, elementsToMerge: []int{}, expectedElementsMerged: []int{1, 2}},
		{name: "BothEmpty", elementsBase: []int{}, elementsToMerge: []int{}, expectedElementsMerged: []int{}},
		{name: "Duplicate", elementsBase: []int{1, 2}, elementsToMerge: []int{1}, expectedElementsMerged: []int{1, 2}},
	}

	for _, testCase := range testSuite {
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
