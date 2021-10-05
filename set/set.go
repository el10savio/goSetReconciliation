package set

type Set struct {
	List []int
	// BloomFilter
	// Hash
}

func Initialize() Set {
	return Set{
		List: []int{},
	}
}
