package aocutil

import (
	"cmp"
	"fmt"
	"sort"
)

type ImmutableSet[T cmp.Ordered] struct {
	m map[T]struct{}
}

func NewImmutableSet[T cmp.Ordered](vs ...T) *ImmutableSet[T] {
	s := &ImmutableSet[T]{
		m: make(map[T]struct{}),
	}
	for _, v := range vs {
		s.m[v] = struct{}{}
	}
	return s
}

func (s *ImmutableSet[T]) copy() *ImmutableSet[T] {
	c := NewImmutableSet[T]()
	for k, v := range s.m {
		c.m[k] = v
	}
	return c
}

func (s *ImmutableSet[T]) With(vs ...T) *ImmutableSet[T] {
	c := s.copy()
	for _, v := range vs {
		c.m[v] = struct{}{}
	}
	return c
}

func (s *ImmutableSet[T]) Without(vs ...T) *ImmutableSet[T] {
	c := s.copy()
	for _, v := range vs {
		delete(c.m, v)
	}
	return c
}

func (s *ImmutableSet[T]) Has(v T) bool {
	_, ok := s.m[v]
	return ok
}

func (s *ImmutableSet[T]) Size() int {
	return len(s.m)
}

func (s *ImmutableSet[T]) Values() []T {
	vals := make([]T, 0, len(s.m))
	for k := range s.m {
		vals = append(vals, k)
	}
	return vals
}

func (s *ImmutableSet[T]) Union(other *ImmutableSet[T]) *ImmutableSet[T] {
	c := s.copy()
	for k := range other.m {
		c.m[k] = struct{}{}
	}
	return c
}

func (s *ImmutableSet[T]) Intersection(other *ImmutableSet[T]) *ImmutableSet[T] {
	c := NewImmutableSet[T]()
	for k := range s.m {
		if other.Has(k) {
			c.m[k] = struct{}{}
		}
	}
	return c
}

func (s *ImmutableSet[T]) WithoutElements(other *ImmutableSet[T]) *ImmutableSet[T] {
	c := NewImmutableSet[T]()
	for k := range s.m {
		if !other.Has(k) {
			c.m[k] = struct{}{}
		}
	}
	return c
}

func (s *ImmutableSet[T]) Xor(other *ImmutableSet[T]) *ImmutableSet[T] {
	return s.WithoutElements(other).Union(other.WithoutElements(s))
}

func (s *ImmutableSet[T]) Equals(other *ImmutableSet[T]) bool {
	return len(s.m) == len(other.m) && s.WithoutElements(other).Size() == 0
}

func (s *ImmutableSet[T]) String() string {
	vs := s.Values()
	sort.Slice(vs, func(i, j int) bool {
		return vs[i] < vs[j]
	})
	return fmt.Sprintf("%v", vs)
}

type ImmutableMultiSet[T cmp.Ordered] struct {
	m map[T]int
}
