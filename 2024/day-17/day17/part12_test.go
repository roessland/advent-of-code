package day17

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExamples(t *testing.T) {
	{
		vm := Vm{
			Program: []Instruction{0, 1, 5, 4, 3, 0},
			State:   State{Regs: [3]uint{2024, 0, 0}},
		}
		vm.Run()
		assert.EqualValues(t,
			[]byte{4, 2, 5, 6, 7, 7, 7, 7, 3, 1, 0},
			vm.Output,
		)
		assert.EqualValues(t, 0, vm.State.Regs[IdxA])
	}
	{
		vm := Vm{
			Program: []Instruction{2, 6},
			State:   State{Regs: [3]uint{0, 0, 9}},
		}
		s, _ := vm.NextState()
		assert.EqualValues(t, 1, s.Regs[IdxB])
	}
	{
		vm := Vm{
			Program: []Instruction{5, 0, 5, 1, 5, 4},
			State:   State{Regs: [3]uint{10, 0, 0}},
		}
		vm.Run()
		assert.EqualValues(t, []byte{0, 1, 2}, vm.Output)
	}
	{
		vm := Vm{
			Program: []Instruction{1, 7},
			State:   State{Regs: [3]uint{0, 29, 0}},
		}
		vm.Run()
		assert.EqualValues(t, 26, vm.State.Regs[IdxB])
	}
	{
		vm := Vm{
			Program: []Instruction{4, 0},
			State:   State{Regs: [3]uint{0, 2024, 43690}},
		}
		vm.Run()
		assert.EqualValues(t, 44354, vm.State.Regs[IdxB])
	}
	{
		vm := Vm{
			Program: []Instruction{0, 1, 5, 4, 3, 0},
			State:   State{Regs: [3]uint{729, 0, 0}},
		}
		vm.Run()
		assert.EqualValues(t, []byte{4, 6, 3, 5, 6, 3, 5, 2, 1, 0}, vm.Output)
	}
	{
		vm := Vm{
			Program: []Instruction{2, 4, 1, 3, 7, 5, 1, 5, 0, 3, 4, 3, 5, 5, 3, 0},
			State:   State{Regs: [3]uint{47006051, 0, 0}},
		}
		vm.Run()
		assert.EqualValues(t, []byte{6, 2, 7, 2, 3, 1, 6, 0, 5}, vm.Output)
	}

	// Part12("input.txt")
}

// 5,2,3,6,2,4,5,0,4 wrong
