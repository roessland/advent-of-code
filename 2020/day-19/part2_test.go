package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAddRuleConst(t *testing.T) {
	var is Set
	rules := make(Rules)

	rules.AddRule(`0: "a"`)
	is = rules["0"]("a", 0)
	require.True(t, SetEqual(is, Set{1: true}))

	// Reference
	rules.AddRule(`1: 0 0`)
	is = rules["1"]("aa", 0)
	require.True(t, SetEqual(is, Set{2: true}))

	// Or
	rules.AddRule(`1: 0 0 | 0`)
	is = rules["1"]("a", 0)
	require.True(t, SetEqual(is, Set{1: true}))
	is = rules["1"]("aa", 0)
	require.True(t, SetEqual(is, Set{1: true, 2: true}))

	// Recursion
	rules.AddRule(`2: 0 2 | 0`)
	is = rules["2"]("a", 0)
	require.True(t, SetEqual(is, Set{1: true}))
	is = rules["2"]("aa", 0)
	require.True(t, SetEqual(is, Set{1: true, 2: true}))
	is = rules["2"]("aaa", 0)
	require.True(t, SetEqual(is, Set{1: true, 2: true, 3: true}), is)

	/*
		// Or TODO FIX
		rules.AddRule(`3: "b"`)
		rules.AddRule(`4: "c"`)
		rules.AddRule(`5: "d"`)
		rules.AddRule(`6: 3 | 4 | 5 | 1 1`)
		require.Equal(t, 0, rules["6"]("a"))

		// And
		rules.AddRule(`7: 0 3 0`)
		require.Equal(t, 3, rules["7"]("aba"))
		require.Equal(t, 0, rules["7"]("aaa"))
		require.Equal(t, 0, rules["7"]("ab"))

		// Recursion
		rules.AddRule( `8: 3 4 | 3 8 4`)
		require.Equal(t, 2, rules["8"]("bc"))
		require.Equal(t, 4, rules["8"]("bbcc"))
		require.Equal(t, 6, rules["8"]("bbbccc"))

		// Circular references
		rules.AddRule(`9: 0 10`)
		rules.AddRule(`10: 3 11`)
		rules.AddRule(`11: 4 9 | 4 5`)
		require.Equal(t, 4, rules["9"]("abcd", 0))
		require.Equal(t, 4+3, rules["9"]("abcabcd",0))
		require.Equal(t, 4+6, rules["9"]("abcabcabcd",0))

		// OR where first doesn't match doesn't match entire length
		rules.AddRule(`12: 13 | 14`)
		rules.AddRule(`13: 0 0 0`)
		rules.AddRule(`14: 0 0 0 0`)
		require.Equal(t, 4, rules["12"]("aaaa",0))
	*/
}
