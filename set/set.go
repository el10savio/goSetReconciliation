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

func (set *Set) GetElements() []int {
	return set.List
}

func (set *Set) AddElement(element int) {
	set.List = append(set.List, element)
}
