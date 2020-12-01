package utils

type IntSet struct {
	m map[int]bool
}

func NewIntSet() *IntSet {
	s := &IntSet{}
	s.m = make(map[int]bool)
	return s
}

func (s *IntSet) Add(value int) {
	s.m[value] = true
}

func (s *IntSet) Remove(value int) {
	delete(s.m, value)
}

func (s *IntSet) Contains(value int) bool {
	_, c := s.m[value]
	return c
}

func (s *IntSet) Find(f func(int) bool) (int, bool) {
	for item := range s.m {
		if f(item) {
			return item, true
		}
	}

	return 0, false
}
