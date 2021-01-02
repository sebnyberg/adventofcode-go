package structx

type IntSet map[int]struct{}

func NewIntSet(ns ...int) IntSet {
	s := make(IntSet, len(ns))
	s.Add(ns...)
	return s
}

func (s *IntSet) GetAll() []int {
	res := make([]int, 0, len((*s)))
	for n := range *s {
		res = append(res, n)
	}
	return res
}

func (s *IntSet) Has(n int) bool {
	if _, exists := (*s)[n]; exists {
		return true
	}
	return false
}

func (s *IntSet) Add(ns ...int) {
	for _, n := range ns {
		(*s)[n] = struct{}{}
	}
}

func (s *IntSet) Remove(ns ...int) {
	for _, n := range ns {
		delete((*s), n)
	}
}
