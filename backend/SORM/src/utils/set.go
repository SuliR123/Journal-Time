package utils

type Set[T comparable] map[T]struct{}

// NewSet initializes and returns a new empty set.
func NewSet[T comparable]() Set[T] {
	return make(Set[T])
}

// Add inserts an element into the set.
func (s Set[T]) Add(value T) {
	s[value] = struct{}{}
}

// Remove deletes an element from the set.
func (s Set[T]) Remove(value T) {
	delete(s, value)
}

// Has checks if an element exists in the set.
func (s Set[T]) Has(value T) bool {
	_, exists := s[value]
	return exists
}
