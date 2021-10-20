package set

var testAddElementsTestSuite = []struct {
	name             string
	elementsToAdd    []int
	expectedElements []int
}{
	{"BasicFuntionality", []int{1, 2, 3}, []int{1, 2, 3}},
	{"Duplicate", []int{1, 1}, []int{1}},
}

var testGetElementsTestSuite = []struct {
	name             string
	inputElements    []int
	expectedElements []int
}{
	{"BasicFuntionality", []int{1, 2, 3}, []int{1, 2, 3}},
	{"Empty", []int{}, []int{}},
}

var testIsElementInBFTestSuite = []struct {
	name              string
	elementToCheck    int
	elementsToBeAdded []int
	expectedCondition bool
}{
	{"BasicFuntionality", 1, []int{1}, true},
	{"NotPresent", 2, []int{1}, false},
	{"EmptyBloomFilter", 1, []int{}, false},
	{"MultipleDuplicateElements", 1, []int{1, 1, 2, 3}, true},
}

var testMergeElementsTestSuite = []struct {
	name                   string
	elementsBase           []int
	elementsToMerge        []int
	expectedElementsMerged []int
}{
	{"BasicFuntionality", []int{1, 2}, []int{3, 4, 5}, []int{1, 2, 3, 4, 5}},
	{"Empty", []int{1, 2}, []int{}, []int{1, 2}},
	{"BothEmpty", []int{}, []int{}, []int{}},
	{"Duplicate", []int{1, 2}, []int{1}, []int{1, 2}},
}
