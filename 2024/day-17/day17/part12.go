package day17

import (
	"embed"
	"fmt"
	"math/rand"
	"time"

	"github.com/roessland/advent-of-code/2024/aocutil"
)

//go:embed input*.txt
var Input embed.FS

type Vm struct {
	Program []Instruction
	Output  []byte
	State   State
}

type Instruction = byte

const (
	Adv = Instruction(0)
	Bxl = Instruction(1)
	Bst = Instruction(2)
	Jnz = Instruction(3)
	Bxc = Instruction(4)
	Out = Instruction(5)
	Bdv = Instruction(6)
	Cdv = Instruction(7)
)

type State struct {
	IP   byte
	Regs [3]uint
}

const (
	RegOffset = 4
	IdxA      = 0
	IdxB      = 1
	IdxC      = 2
)

func (vm *Vm) Combo(addr byte) uint {
	combo := vm.Program[addr]
	if combo < RegOffset {
		return uint(combo)
	}
	return vm.State.Regs[combo-RegOffset]
}

func (vm *Vm) NextState() (s State, halt bool) {
	time.Sleep(0 * time.Second)
	s = vm.State
	if s.IP >= byte(len(vm.Program)) {
		return s, true
	}

	op := vm.Program[s.IP]
	switch op {
	case Adv: // 0
		num := uint(s.Regs[IdxA])
		denom := uint(1) << uint(vm.Combo(s.IP+1))
		result := num / denom
		s.Regs[IdxA] = result
	case Bxl: // 1
		lit := uint(vm.Program[s.IP+1])
		result := uint(s.Regs[IdxB]) ^ uint(lit)
		s.Regs[IdxB] = result
	case Bst: // 2
		combo := vm.Combo(s.IP + 1)
		result := combo % 8
		s.Regs[IdxB] = result
	case Jnz: // 3
		if s.Regs[IdxA] == 0 {
			break
		}
		lit := vm.Program[s.IP+1]
		s.IP = byte(lit)
		s.IP -= 2
	case Bxc: // 4
		result := s.Regs[IdxB] ^ s.Regs[IdxC]
		s.Regs[IdxB] = result
	case Out: // 5
		combo := vm.Combo(s.IP + 1)
		result := combo % 8
		vm.Output = append(vm.Output, byte(result))
	case Bdv:
		num := uint(s.Regs[IdxA])
		denom := uint(1) << uint(vm.Combo(s.IP+1))
		result := num / denom
		s.Regs[IdxB] = result
	case Cdv:
		num := uint(s.Regs[IdxA])
		denom := uint(1) << uint(vm.Combo(s.IP+1))
		result := num / denom
		s.Regs[IdxC] = result
	default:
		panic("unknown instruction")
	}

	s.IP += 2
	return s, false
}

func (vm *Vm) Run() State {
	s := vm.State
	var halted bool
	for ; !halted; s, halted = vm.NextState() {
		vm.State = s
	}
	return s
}

// 4,6,3,5,6,3,5,2,1,0 not the right answer
func Part12(inputName string) (string, int) {
	input := aocutil.FSGetIntsInStringLines(Input, inputName)
	A, B, C := input[0][0], input[1][0], input[2][0]
	var program []byte
	for _, num := range input[4] {
		program = append(program, Instruction(num))
	}

	vm := Vm{
		Program: program,
		State:   State{Regs: [3]uint{uint(A), uint(B), uint(C)}},
	}
	vm.Run()
	var p1 string
	for _, b := range vm.Output {
		p1 = p1 + fmt.Sprintf("%d,", b)
	}
	p1 = p1[:len(p1)-1]

	{

		// Random bit flip fuckery.
		// Not proud of this code.
		// Returns correctd answer 50% of the time.
		// Sometimes too high.
		var origA uint = 0b111_111_111_111_111_111_111_111_111_111_111_111_111_111_111_111
		var A uint = 0b111_111_111_111_111_111_111_111_111_111_111_111_111_111_111_111
		numCorrect := 0
		for {
			// Random restarts
			if rand.Intn(10000) < 10 {
				numCorrect = 0
				A = origA
			}
			altA := A
			for i := 0; i < 10; i++ {
				if rand.Int()%10 < 5 {
					k := rand.Intn(48)
					altA = A ^ (1 << k)
				}
			}
			altOut := Run(altA, 0, 0)
			altCorrect := NumCorrect(altOut)
			if altCorrect >= numCorrect {
				if numCorrect == 16 {
					return p1, int(altA)
				}
				numCorrect = altCorrect
				A = altA
			}
		}
	}
}

func Run(A, B, C uint) []byte {
	vm := Vm{
		Program: []Instruction{2, 4, 1, 3, 7, 5, 1, 5, 0, 3, 4, 3, 5, 5, 3, 0},
		State:   State{Regs: [3]uint{uint(A), uint(B), uint(C)}},
	}
	vm.Run()
	return vm.Output
}

func NumCorrect(output []byte) int {
	desired := []uint{2, 4, 1, 3, 7, 5, 1, 5, 0, 3, 4, 3, 5, 5, 3, 0}
	numCorrect := 0
	for i, b := range output {
		if b == byte(desired[i]) {
			numCorrect++
		}
	}
	return numCorrect
}
