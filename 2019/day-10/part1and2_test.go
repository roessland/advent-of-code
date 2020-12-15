package main

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"math"
	"sort"
	"testing"
)

func TestGCD(t *testing.T) {
	require.Equal(t, 18, GCD(270, 198))
	require.Equal(t, 18, GCD(-270, 198))
	require.Equal(t, 18, GCD(270, -198))
	require.Equal(t, 18, GCD(-270, -198))
}

func TestAtan2AngleSorting(t *testing.T) {
	as := []Asteroid{
		{-1000, 0},
		{-1000, 1},
		{-1, 1000},
		{1, 1000},
		{-1000, -1},
	}

	for _, a := range as {
		fmt.Println(a, math.Atan2(float64(-a.J)-0.000001, float64(a.I)))
	}

	require.True(t, sort.SliceIsSorted(as, func(i,j int)bool{
		return math.Atan2(float64(-as[i].J)-0.000001, float64(as[i].I)) < math.Atan2(float64(-as[j].J)-0.000001, float64(as[j].I))
	}))
}