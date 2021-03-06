package sync

import (
	"bytes"
	"errors"
	"net/http"
	"os"
	"strings"
	"time"
)

// GetPeerList Obtains Peer List
// From Environment Variable
func GetPeerList() []string {
	return strings.Split(os.Getenv("PEERS"), ",")
}

// GetNetwork Obtains Network
// From Environment Variable
func GetNetwork() string {
	return os.Getenv("NETWORK") + ":8080"
}

// GetHost Obtains Host
// From Environment Variable
func GetHost() string {
	return os.Getenv("HOST")
}

// SendRequest handles sending of an HTTP POST Request
func SendRequest(url string, payload []byte) (int, error) {
	if url == "" {
		return 0, errors.New("empty url provided")
	}

	client := http.Client{
		Timeout: time.Duration(10 * time.Second),
	}

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return 0, err
	}
	request.Header.Set("Content-Type", "application/json")

	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	return response.StatusCode, nil
}
