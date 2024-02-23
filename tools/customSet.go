package tools

type Set struct {
	elements map[string]struct{}
}

func NewSet() *Set {
	return &Set{
		elements: make(map[string]struct{}),
	}
}

func (s *Set) Add(element string) {
	s.elements[element] = struct{}{}
}

func (s *Set) Contains(element string) bool {
	_, exist := s.elements[element]
	return exist
}
