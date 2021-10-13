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

func TestUpdate(t *testing.T) {
	baseSet := set.Initialize()
	defer baseSet.Clear()

	payloadSet := set.Initialize()
	defer payloadSet.Clear()

	baseSet.AddElement(1)
	baseSet.AddElement(2)

	payloadSet.AddElement(1)
	payloadSet.AddElement(2)
	payloadSet.AddElement(3)

	payload := Payload{
		MissingElements: []uint32{3},
		BF:              payloadSet.GetBF(),
		Hash:            payloadSet.GetHash(),
	}

	expectedMissingElements := []uint32{}
	expectedSet := set.Initialize()
	defer expectedSet.Clear()

	expectedSet.AddElement(1)
	expectedSet.AddElement(2)
	expectedSet.AddElement(3)

	baseSet, actualMissingElements := Update(baseSet, payload)

	assert.Equal(t, expectedSet, baseSet)
	assert.Equal(t, expectedMissingElements, actualMissingElements)
}

func TestUpdate_EmptySet(t *testing.T) {
	baseSet := set.Initialize()
	defer baseSet.Clear()

	payloadSet := set.Initialize()
	defer payloadSet.Clear()

	payloadSet.AddElement(1)
	payloadSet.AddElement(2)
	payloadSet.AddElement(3)

	payload := Payload{
		MissingElements: []uint32{1, 2, 3},
		BF:              payloadSet.GetBF(),
		Hash:            payloadSet.GetHash(),
	}

	expectedMissingElements := []uint32{}
	expectedSet := set.Initialize()
	defer expectedSet.Clear()

	expectedSet.AddElement(1)
	expectedSet.AddElement(2)
	expectedSet.AddElement(3)

	baseSet, actualMissingElements := Update(baseSet, payload)

	assert.Equal(t, expectedSet, baseSet)
	assert.Equal(t, expectedMissingElements, actualMissingElements)
}

func TestUpdate_EmptyPayload(t *testing.T) {
	baseSet := set.Initialize()
	defer baseSet.Clear()

	baseSet.AddElement(1)
	baseSet.AddElement(2)

	payloadSet := set.Initialize()
	defer payloadSet.Clear()

	payload := Payload{
		MissingElements: []uint32{},
		BF:              payloadSet.GetBF(),
		Hash:            payloadSet.GetHash(),
	}

	expectedMissingElements := []uint32{1, 2}
	expectedSet := set.Initialize()
	defer expectedSet.Clear()

	expectedSet.AddElement(1)
	expectedSet.AddElement(2)

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
		MissingElements: []uint32{},
		BF:              payloadSet.GetBF(),
		Hash:            payloadSet.GetHash(),
	}

	expectedMissingElements := []uint32{}
	expectedSet := set.Initialize()
	defer expectedSet.Clear()

	baseSet, actualMissingElements := Update(baseSet, payload)

	assert.Equal(t, expectedSet, baseSet)
	assert.Equal(t, expectedMissingElements, actualMissingElements)
}

func TestUpdate_FullSync(t *testing.T) {
	baseSet := set.Initialize()
	defer baseSet.Clear()

	baseSet.AddElement(1)
	baseSet.AddElement(2)

	payloadSet := set.Initialize()
	defer payloadSet.Clear()

	payload := Payload{
		MissingElements: []uint32{},
		BF:              payloadSet.GetBF(),
		Hash:            payloadSet.GetHash(),
	}

	baseSet, actualMissingElements := Update(baseSet, payload)

	payload = Payload{
		MissingElements: actualMissingElements,
		BF:              baseSet.GetBF(),
		Hash:            baseSet.GetHash(),
	}

	expectedMissingElements := []uint32{}
	expectedSet := set.Initialize()
	defer expectedSet.Clear()

	expectedSet.AddElement(1)
	expectedSet.AddElement(2)

	payloadSet, actualMissingElements = Update(payloadSet, payload)

	assert.Equal(t, expectedSet, baseSet)
	assert.Equal(t, baseSet, payloadSet)
	assert.Equal(t, expectedMissingElements, actualMissingElements)
}

func TestUpdate_MixedSync(t *testing.T) {
	baseSet := set.Initialize()
	defer baseSet.Clear()

	baseSet.AddElement(1)
	baseSet.AddElement(2)
	baseSet.AddElement(6)

	payloadSet := set.Initialize()
	defer payloadSet.Clear()

	baseSet.AddElement(1)
	baseSet.AddElement(2)
	baseSet.AddElement(3)
	baseSet.AddElement(4)
	baseSet.AddElement(5)

	payload := Payload{
		MissingElements: []uint32{},
		BF:              payloadSet.GetBF(),
		Hash:            payloadSet.GetHash(),
	}

	expectedMissingElements := []uint32{1, 2, 6, 3, 4, 5}
	expectedSet := set.Initialize()
	defer expectedSet.Clear()

	expectedSet.AddElement(1)
	expectedSet.AddElement(2)
	expectedSet.AddElement(6)
	expectedSet.AddElement(3)
	expectedSet.AddElement(4)
	expectedSet.AddElement(5)

	baseSet, actualMissingElements := Update(baseSet, payload)

	assert.Equal(t, expectedSet, baseSet)
	assert.Equal(t, expectedMissingElements, actualMissingElements)

	payload = Payload{
		MissingElements: actualMissingElements,
		BF:              baseSet.GetBF(),
		Hash:            baseSet.GetHash(),
	}

	expectedMissingElements = []uint32{}

	payloadSet, actualMissingElements = Update(payloadSet, payload)

	assert.Equal(t, expectedSet, payloadSet)
	assert.Equal(t, expectedMissingElements, actualMissingElements)
}
