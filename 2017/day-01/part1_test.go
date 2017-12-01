package main

import "testing"
import "github.com/stretchr/testify/assert"

func TestSolveCaptcha1(t *testing.T) {
	inputs := [][]int{
		[]int{1, 1, 2, 2},
		[]int{1, 1, 1, 1},
		[]int{1, 2, 3, 4},
		[]int{9, 1, 2, 1, 2, 1, 2, 9},
	}
	outputs := []int{
		3, 4, 0, 9,
	}
	for i := range inputs {
		assert.Equal(t, outputs[i], SolveCaptcha1(inputs[i]))
	}
}

func TestSolveCaptcha2(t *testing.T) {
	inputs := [][]int{
		[]int{1, 2, 1, 2},
		[]int{1, 2, 2, 1},
		[]int{1, 2, 3, 4, 2, 5},
		[]int{1, 2, 3, 1, 2, 3},
		[]int{1, 2, 1, 3, 1, 4, 1, 5},
	}
	outputs := []int{
		6, 0, 4, 12, 4,
	}
	for i := range inputs {
		assert.Equal(t, outputs[i], SolveCaptcha2(inputs[i]))
	}
}
