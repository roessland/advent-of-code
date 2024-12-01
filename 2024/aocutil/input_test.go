package aocutil_test

import (
	"testing"

	"github.com/roessland/advent-of-code/2024/aocutil"
	"github.com/stretchr/testify/require"
)

func TestGetNums(t *testing.T) {
	require.EqualValues(t, []int{5}, aocutil.GetIntsInString("fdsaf 5,hey"))
	require.EqualValues(t, []int{-5}, aocutil.GetIntsInString("fdsaf,-5,hey"))
	require.EqualValues(t, []int{5, 3}, aocutil.GetIntsInString("banana-5-cloak-3-hey"))
	require.EqualValues(t, []int{773, 332, 773, 332}, aocutil.GetIntsInString("773,332,773,332"))
	require.EqualValues(t, []int{123, 234, 456}, aocutil.GetIntsInString("123+234-456"))
	require.EqualValues(t, []int{123, 234, -456}, aocutil.GetIntsInString("123 +234 -456"))
}
