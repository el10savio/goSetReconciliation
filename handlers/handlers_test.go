package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestIndex tests the basic behaviour
// of the Index Handler
func TestIndex(t *testing.T) {
	request, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Index)
	handler.ServeHTTP(rr, request)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "Hello World Set Node\n", rr.Body.String())
}

// TestAdd tests the basic behaviour
// of the Add Handler
func TestAdd(t *testing.T) {
	defer testClearUtil(t)
	testAddUtil(t, "[1]", http.StatusOK)
}

// TestGet tests the basic behaviour
// of the Get Handler
func TestGet(t *testing.T) {
	defer testClearUtil(t)
	testAddUtil(t, "[1]", http.StatusOK)
	testGetUtil(t, "[1]\n", http.StatusOK)
}

// TestGet_EmptyElement tests the behaviour of the List Handler
// When there are no elements present
func TestGet_Empty(t *testing.T) {
	defer testClearUtil(t)
	testGetUtil(t, "[]\n", http.StatusOK)
}

// TestGet_DuplicateElement tests the behaviour of the List Handler
// When there are duplicate elements trying to get added in
func TestGet_DuplicateElement(t *testing.T) {
	defer testClearUtil(t)
	testAddUtil(t, "[1,1]", http.StatusOK)
	testGetUtil(t, "[1]\n", http.StatusOK)
}

// TestGet_MultipleElements tests the behaviour of the List Handler
// When there are multiple elements trying to get added in
func TestGet_MultipleElements(t *testing.T) {
	defer testClearUtil(t)
	testAddUtil(t, "[1,2,3]", http.StatusOK)
	testGetUtil(t, "[1,2,3]\n", http.StatusOK)
}

// testAddUtil is a test utility function
// used to send requests to the Add Handler
func testAddUtil(t *testing.T, elements string, expectedStatus int) {
	payload := fmt.Sprintf(`{"values":%s}`, elements)

	request, err := http.NewRequest("POST", "/set/add", strings.NewReader(payload))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Add)
	handler.ServeHTTP(rr, request)

	assert.Equal(t, expectedStatus, rr.Code)
}

// testGetUtil is a test utility function
// used to send requests to the List Handler
func testGetUtil(t *testing.T, expectedList string, expectedStatus int) {
	request, err := http.NewRequest("GET", "/set/list", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(List)
	handler.ServeHTTP(rr, request)

	assert.Equal(t, expectedStatus, rr.Code)
	assert.Equal(t, expectedList, rr.Body.String())
}

// testClearUtil is a test utility function
// used to send requests to the Clear Handler
func testClearUtil(t *testing.T) {
	request, err := http.NewRequest("GET", "/set/debug/clear", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Clear)
	handler.ServeHTTP(rr, request)

	assert.Equal(t, http.StatusOK, rr.Code)
	testGetUtil(t, "[]\n", http.StatusOK)
}
