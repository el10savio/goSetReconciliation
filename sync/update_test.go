package sync

import (
	"testing"

	"github.com/el10savio/goSetReconciliation/set"
	"github.com/stretchr/testify/assert"
)

func TestGetBFMissingElements(t *testing.T) {
	list := []int{1}
	set := set.Initialize()
	defer set.Clear()

	set.AddElements([]int{2, 3})

	expectedElements := []int{1}
	actualElements := GetBFMissingElements(list, set.BF)

	assert.Equal(t, expectedElements, actualElements)
}

func TestGetBFMissingElements_EmptyBF(t *testing.T) {
	list := []int{1, 2, 3}
	set := set.Initialize()
	defer set.Clear()

	expectedElements := []int{1, 2, 3}
	actualElements := GetBFMissingElements(list, set.BF)

	assert.Equal(t, expectedElements, actualElements)
}

func TestGetBFMissingElements_EmptyList(t *testing.T) {
	list := []int{}
	set := set.Initialize()
	defer set.Clear()

	set.AddElements([]int{2, 3})

	expectedElements := []int{}
	actualElements := GetBFMissingElements(list, set.BF)

	assert.Equal(t, expectedElements, actualElements)
}

func TestGetBFMissingElements_BothEmpty(t *testing.T) {
	list := []int{}
	set := set.Initialize()
	defer set.Clear()

	expectedElements := []int{}
	actualElements := GetBFMissingElements(list, set.BF)

	assert.Equal(t, expectedElements, actualElements)
}

func TestUpdate(t *testing.T) {
	baseSet := set.Initialize()
	defer baseSet.Clear()

	payloadSet := set.Initialize()
	defer payloadSet.Clear()

	baseSet.AddElements([]int{1, 2})
	payloadSet.AddElements([]int{1, 2, 3})

	payload := Payload{
		MissingElements: []int{3},
		BF:              payloadSet.GetBF(),
		Hash:            payloadSet.GetHash(),
	}

	expectedMissingElements := []int{}
	expectedSet := set.Initialize()
	defer expectedSet.Clear()

	expectedSet.AddElements([]int{1, 2, 3})

	baseSet, actualMissingElements := Update(baseSet, payload)

	assert.Equal(t, expectedSet, baseSet)
	assert.Equal(t, expectedMissingElements, actualMissingElements)
}

func TestUpdate_BothEqual(t *testing.T) {
	baseSet := set.Initialize()
	defer baseSet.Clear()

	payloadSet := set.Initialize()
	defer payloadSet.Clear()

	baseSet.AddElements([]int{1, 2, 3})
	payloadSet.AddElements([]int{1, 2, 3})

	payload := Payload{
		MissingElements: []int{},
		BF:              payloadSet.GetBF(),
		Hash:            payloadSet.GetHash(),
	}

	expectedMissingElements := []int{}
	expectedSet := set.Initialize()
	defer expectedSet.Clear()

	expectedSet.AddElements([]int{1, 2, 3})

	baseSet, actualMissingElements := Update(baseSet, payload)

	assert.Equal(t, expectedSet, baseSet)
	assert.Equal(t, expectedMissingElements, actualMissingElements)
}

func TestUpdate_EmptySet(t *testing.T) {
	baseSet := set.Initialize()
	defer baseSet.Clear()

	payloadSet := set.Initialize()
	defer payloadSet.Clear()

	payloadSet.AddElements([]int{1, 2, 3})

	payload := Payload{
		MissingElements: []int{1, 2, 3},
		BF:              payloadSet.GetBF(),
		Hash:            payloadSet.GetHash(),
	}

	expectedMissingElements := []int{}
	expectedSet := set.Initialize()
	defer expectedSet.Clear()

	expectedSet.AddElements([]int{1, 2, 3})

	baseSet, actualMissingElements := Update(baseSet, payload)

	assert.Equal(t, expectedSet, baseSet)
	assert.Equal(t, expectedMissingElements, actualMissingElements)
}

func TestUpdate_EmptyPayload(t *testing.T) {
	baseSet := set.Initialize()
	defer baseSet.Clear()

	baseSet.AddElements([]int{1, 2})

	payloadSet := set.Initialize()
	defer payloadSet.Clear()

	payload := Payload{
		MissingElements: []int{},
		BF:              payloadSet.GetBF(),
		Hash:            payloadSet.GetHash(),
	}

	expectedMissingElements := []int{1, 2}
	expectedSet := set.Initialize()
	defer expectedSet.Clear()

	expectedSet.AddElements([]int{1, 2})

	baseSet, actualMissingElements := Update(baseSet, payload)

	assert.Equal(t, expectedSet, baseSet)
	assert.Equal(t, expectedMissingElements, actualMissingElements)
}

func TestUpdate_BothEmpty(t *testing.T) {
	baseSet := set.Initialize()
	defer baseSet.Clear()

	payloadSet := set.Initialize()
	defer payloadSet.Clear()

	payload := Payload{
		MissingElements: []int{},
		BF:              payloadSet.GetBF(),
		Hash:            payloadSet.GetHash(),
	}

	expectedMissingElements := []int{}
	expectedSet := set.Initialize()
	defer expectedSet.Clear()

	baseSet, actualMissingElements := Update(baseSet, payload)

	assert.Equal(t, expectedSet, baseSet)
	assert.Equal(t, expectedMissingElements, actualMissingElements)
}

func TestUpdate_FullSync(t *testing.T) {
	baseSet := set.Initialize()
	defer baseSet.Clear()

	baseSet.AddElements([]int{1, 2})

	payloadSet := set.Initialize()
	defer payloadSet.Clear()

	payload := Payload{
		MissingElements: []int{},
		BF:              payloadSet.GetBF(),
		Hash:            payloadSet.GetHash(),
	}

	baseSet, actualMissingElements := Update(baseSet, payload)

	payload = Payload{
		MissingElements: actualMissingElements,
		BF:              baseSet.GetBF(),
		Hash:            baseSet.GetHash(),
	}

	expectedMissingElements := []int{}
	expectedSet := set.Initialize()
	defer expectedSet.Clear()

	expectedSet.AddElements([]int{1, 2})

	payloadSet, actualMissingElements = Update(payloadSet, payload)

	assert.Equal(t, expectedSet, baseSet)
	assert.Equal(t, baseSet, payloadSet)
	assert.Equal(t, expectedMissingElements, actualMissingElements)
}

func TestUpdate_MixedSync(t *testing.T) {
	baseSet := set.Initialize()
	defer baseSet.Clear()

	baseSet.AddElements([]int{1, 2, 6})

	payloadSet := set.Initialize()
	defer payloadSet.Clear()

	baseSet.AddElements([]int{1, 2, 3, 4, 5})

	payload := Payload{
		MissingElements: []int{},
		BF:              payloadSet.GetBF(),
		Hash:            payloadSet.GetHash(),
	}

	expectedMissingElements := []int{1, 2, 6, 3, 4, 5}
	expectedSet := set.Initialize()
	defer expectedSet.Clear()

	expectedSet.AddElements([]int{1, 2, 6, 3, 4, 5})

	baseSet, actualMissingElements := Update(baseSet, payload)

	assert.Equal(t, expectedSet, baseSet)
	assert.Equal(t, expectedMissingElements, actualMissingElements)

	payload = Payload{
		MissingElements: actualMissingElements,
		BF:              baseSet.GetBF(),
		Hash:            baseSet.GetHash(),
	}

	expectedMissingElements = []int{}

	payloadSet, actualMissingElements = Update(payloadSet, payload)

	assert.Equal(t, expectedSet, payloadSet)
	assert.Equal(t, expectedMissingElements, actualMissingElements)
}
