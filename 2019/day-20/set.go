package main// SET INTERFACE

type Set interface {
	Add(string)
	Contains(string) bool
}

func NewSet() Set {
	return simpleSet{}
}


// SET IMPLEMENTATION

type simpleSet map[string]struct{}

func (s simpleSet) Add(id string) {
	s[id] = struct{}{}
}

func (s simpleSet) Contains(id string) bool {
	_, ok := s[id]
	return ok
}

var _ Set = &simpleSet{}
