package aocutil

import (
	"cmp"
	"fmt"
	"sort"
)

func NewImmutableMultiSet[T cmp.Ordered](vs ...T) *ImmutableMultiSet[T] {
	s := &ImmutableMultiSet[T]{
		m: make(map[T]int),
	}
	for _, v := range vs {
		s.m[v]++
	}
	return s
}

func (s *ImmutableMultiSet[T]) copy() *ImmutableMultiSet[T] {
	c := NewImmutableMultiSet[T]()
	for k, v := range s.m {
		c.m[k] = v
	}
	return c
}

func (s *ImmutableMultiSet[T]) With(vs ...T) *ImmutableMultiSet[T] {
	c := s.copy()
	for _, v := range vs {
		c.m[v]++
	}
	return c
}

func (s *ImmutableMultiSet[T]) Without(vs ...T) *ImmutableMultiSet[T] {
	c := s.copy()
	for _, v := range vs {
		if c.m[v] == 0 {
			panic(fmt.Sprintf("Cannot remove %v from set %v", v, s))
		}
		c.m[v]--
	}
	return c
}

func (s *ImmutableMultiSet[T]) Has(v T) bool {
	return s.m[v] > 0
}

func (s *ImmutableMultiSet[T]) Dimension() int {
	return len(s.m)
}

func (s *ImmutableMultiSet[T]) Supp() *ImmutableSet[T] {
	c := NewImmutableSet[T]()
	for k, v := range s.m {
		if v > 0 {
			c.m[k] = struct{}{}
		}
	}
	return c
}

func (s *ImmutableMultiSet[T]) Multiplicity(v T) int {
	return s.m[v]
}

func (s *ImmutableMultiSet[T]) Multiplicities() []int {
	multiplicities := make([]int, 0, len(s.m))
	for _, v := range s.m {
		multiplicities = append(multiplicities, v)
	}
	return multiplicities
}

func (s *ImmutableMultiSet[T]) Cardinality() int {
	size := 0
	for _, v := range s.m {
		size += v
	}
	return size
}

type Pair[T cmp.Ordered] struct {
	Val          T
	Multiplicity int
}

func (s *ImmutableMultiSet[T]) Values() []Pair[T] {
	vals := make([]Pair[T], 0, len(s.m))
	for k, v := range s.m {
		vals = append(vals, Pair[T]{k, v})
	}
	return vals
}

func (s *ImmutableMultiSet[T]) Union(other *ImmutableMultiSet[T]) *ImmutableMultiSet[T] {
	c := s.copy()
	for k, v := range other.m {
		c.m[k] += v
	}
	return c
}

func (s *ImmutableMultiSet[T]) Intersection(other *ImmutableMultiSet[T]) *ImmutableMultiSet[T] {
	c := NewImmutableMultiSet[T]()
	for k, v := range s.m {
		if other.m[k] > 0 {
			c.m[k] = min(v, other.m[k])
		}
	}
	return c
}

func (s *ImmutableMultiSet[T]) WithoutElements(other *ImmutableMultiSet[T]) *ImmutableMultiSet[T] {
	c := NewImmutableMultiSet[T]()
	for k, v := range s.m {
		multiplicity := v - other.m[k]
		if multiplicity > 0 {
			c.m[k] = multiplicity
		}
	}
	return c
}

func (s *ImmutableMultiSet[T]) Xor(other *ImmutableMultiSet[T]) *ImmutableMultiSet[T] {
	return s.WithoutElements(other).Union(other.WithoutElements(s))
}

func (s *ImmutableMultiSet[T]) Equals(other *ImmutableMultiSet[T]) bool {
	return len(s.m) == len(other.m) && s.WithoutElements(other).Cardinality() == 0
}

func (s *ImmutableMultiSet[T]) String() string {
	vs := s.Values()
	sort.Slice(vs, func(i, j int) bool {
		return vs[i].Multiplicity < vs[j].Multiplicity
	})
	return fmt.Sprintf("%v", vs)
}
