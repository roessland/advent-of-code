package main

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCut(t *testing.T) {
	{
		d := NewDeck(10, 0)
		d = d.Cut(1)
		require.Equal(t, 9, d.CardPos)
	}
	{
		d := NewDeck(10, 1)
		d = d.Cut(1)
		require.Equal(t, 0, d.CardPos)
	}
	{
		d := NewDeck(10, 9)
		d = d.Cut(1)
		require.Equal(t, 8, d.CardPos)
	}
}


func TestIncrement(t *testing.T) {
	{
		d := NewDeck(10, 0)
		d = d.Increment(3)
		require.Equal(t, 0, d.CardPos)
	}
	{
		d := NewDeck(10, 1)
		d = d.Increment(3)
		require.Equal(t, 3, d.CardPos)
	}
	{
		d := NewDeck(10, 9)
		d = d.Increment(3)
		require.Equal(t, 7, d.CardPos)
	}
}

func TestDeal(t *testing.T) {
	{
		d := NewDeck(10, 0)
		d = d.Deal()
		require.Equal(t, 9, d.CardPos)
	}
	{
		d := NewDeck(10, 9)
		d = d.Deal()
		require.Equal(t, 0, d.CardPos)
	}
	{
		d := NewDeck(10, 3)
		d = d.Deal()
		require.Equal(t, 6, d.CardPos)
	}
}

func TestCase1(t *testing.T) {
	d := NewDeck(10, 1)
	d = d.Increment(7)
	d = d.Deal()
	d = d.Deal()
	require.Equal(t, 7, d.CardPos)
}

func TestCase2(t *testing.T) {
	expectedCards := []int{
		9,2,5,8,1,4,7,0,3,6,
	}
	expectedPos := make([]int, 10)
	for i := range expectedCards {
		expectedPos[expectedCards[i]] = i
	}
	fmt.Println(expectedPos)
	for c := 0; c < 10; c++ {
		fmt.Println("c is", c)
		d := NewDeck(10, c)
		d = d.Deal()
		d = d.Cut(-2)
		d = d.Increment(7)
		d = d.Cut(8)
		d = d.Cut(-4)
		d = d.Increment(7)
		d = d.Cut(3)
		d = d.Increment(9)
		d = d.Increment(3)
		d = d.Cut(-1)
		require.Equal(t, expectedPos[c], d.CardPos)
		d = d.RevCut(-1)
		d = d.RevIncrement(3)
		d = d.RevIncrement(9)
		d = d.RevCut(3)
		d = d.RevIncrement(7)
		d = d.RevCut(-4)
		d = d.RevCut(8)
		d = d.RevIncrement(7)
		d = d.RevCut(-2)
		d = d.RevDeal()
		require.Equal(t, c, d.CardPos)
	}

}