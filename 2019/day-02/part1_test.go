package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	var vm *Vm
	vm = NewVm([]int{1, 0, 0, 0, 99})
	assert.Equal(t, 2, vm.Run())
	vm = NewVm([]int{2, 3, 0, 3, 99})
	assert.Equal(t, 2, vm.Run())
	vm = NewVm([]int{2, 4, 4, 5, 99, 0})
	assert.Equal(t, 2, vm.Run())
	vm = NewVm([]int{1, 1, 1, 4, 99, 5, 6, 0, 99})
	assert.Equal(t, 30, vm.Run())
}
