package sync

import "github.com/bits-and-blooms/bloom/v3"

// Payload is the representation of the message
// passed between nodes to sync with each other
type Payload struct {
	MissingElements []int              `json:"missingElements"`
	BF              *bloom.BloomFilter `json:"bf"`
	Hash            uint64             `json:"hash"`
}
