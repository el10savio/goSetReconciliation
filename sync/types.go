package sync

import "github.com/bits-and-blooms/bloom/v3"

// Payload ...
type Payload struct {
	MissingElements []int              `json:"missingElements"`
	BF              *bloom.BloomFilter `json:"bf"`
	Hash            uint64             `json:"hash"`
}
