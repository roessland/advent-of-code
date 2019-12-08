package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"strconv"
	"strings"
)

type Vm struct {
	OriginalMem []int
	Mem         []int
	Ip          int
	Halted      bool
	InputCh     <-chan int
	OutputCh    chan<- int
}

func NewVm(mem []int) *Vm {
	vm := Vm{}
	vm.OriginalMem = make([]int, len(mem))
	copy(vm.OriginalMem, mem)
	vm.Mem = make([]int, len(mem))
	return &vm
}

type ParameterMode int

const (
	PositionMode  int = 0
	ImmediateMode int = 1
)

type Param struct {
	Val  int
	Mode int
}

type Op struct {
	Code     int
	Params   []Param
	Length   int
	FullCode int
}

func (vm *Vm) ReadOp() Op {
	op := Op{}
	val := vm.Mem[vm.Ip]
	op.FullCode = val
	op.Code = val % 100
	switch op.Code {
	case 1:
		op.Length = 4
	case 2:
		op.Length = 4
	case 3:
		op.Length = 2
	case 4:
		op.Length = 2
	case 5:
		op.Length = 3
	case 6:
		op.Length = 3
	case 7:
		op.Length = 4
	case 8:
		op.Length = 4
	case 99:
		op.Length = 1
	default:
		log.Fatal("ReadOp: unknown opcode", op.Code)
	}
	if op.Length >= 2 {
		op.Params = append(op.Params, Param{Val: vm.Mem[vm.Ip+1], Mode: (val / 100) % 10})
	}
	if op.Length >= 3 {
		op.Params = append(op.Params, Param{Val: vm.Mem[vm.Ip+2], Mode: (val / 1000) % 10})
	}
	if op.Length >= 4 {
		op.Params = append(op.Params, Param{Val: vm.Mem[vm.Ip+3], Mode: (val / 10000) % 10})
	}
	if op.Length >= 5 {
		op.Params = append(op.Params, Param{Val: vm.Mem[vm.Ip+4], Mode: (val / 100000) % 10})
	}
	if op.Length >= 6 {
		log.Fatal("ReadOp: instruction of length", op.Length, "not supported")
	}
	return op
}

func (vm *Vm) GetVal(val int, mode int) int {
	if mode == PositionMode {
		return vm.Mem[val]
	} else if mode == ImmediateMode {
		return val
	} else {
		log.Fatal("GetVal: unknown position mode", mode)
		return -1337
	}
}

func (vm *Vm) Run(inputCh <-chan int, outputCh chan<- int, signalCh chan<- int) {
	copy(vm.Mem, vm.OriginalMem)
	vm.Ip = 0
	vm.Halted = false
	lastOutput := 0
	for !vm.Halted {
		//fmt.Println("MEM", vm.Mem)
		//fmt.Println("IP", vm.Ip)
		op := vm.ReadOp()
		//fmt.Printf("OP: %#v\n", op)
		if op.Code == 1 {
			// Add
			vm.Mem[op.Params[2].Val] = vm.GetVal(op.Params[0].Val, op.Params[0].Mode) + vm.GetVal(op.Params[1].Val, op.Params[1].Mode)
			vm.Ip += op.Length
		} else if op.Code == 2 {
			// Multiply
			vm.Mem[op.Params[2].Val] = vm.GetVal(op.Params[0].Val, op.Params[0].Mode) * vm.GetVal(op.Params[1].Val, op.Params[1].Mode)
			vm.Ip += op.Length
		} else if op.Code == 3 {
			// Input
			vm.Mem[op.Params[0].Val] = <-inputCh
			vm.Ip += op.Length
		} else if op.Code == 4 {
			// Output
			lastOutput = vm.GetVal(op.Params[0].Val, op.Params[0].Mode)
			outputCh <- lastOutput
			vm.Ip += op.Length
		} else if op.Code == 5 {
			// jump-if-true
			if vm.GetVal(op.Params[0].Val, op.Params[0].Mode) != 0 {
				vm.Ip = vm.GetVal(op.Params[1].Val, op.Params[1].Mode)
			} else {
				vm.Ip += op.Length
			}
		} else if op.Code == 6 {
			// jump-if-false
			if vm.GetVal(op.Params[0].Val, op.Params[0].Mode) == 0 {
				vm.Ip = vm.GetVal(op.Params[1].Val, op.Params[1].Mode)
			} else {
				vm.Ip += op.Length
			}
		} else if op.Code == 7 {
			// less than
			if vm.GetVal(op.Params[0].Val, op.Params[0].Mode) < vm.GetVal(op.Params[1].Val, op.Params[1].Mode) {
				vm.Mem[op.Params[2].Val] = 1
			} else {
				vm.Mem[op.Params[2].Val] = 0
			}
			vm.Ip += op.Length
		} else if op.Code == 8 {
			// equals
			if vm.GetVal(op.Params[0].Val, op.Params[0].Mode) == vm.GetVal(op.Params[1].Val, op.Params[1].Mode) {
				vm.Mem[op.Params[2].Val] = 1
			} else {
				vm.Mem[op.Params[2].Val] = 0
			}
			vm.Ip += op.Length
		} else if op.Code == 99 {
			// Halt
			vm.Halted = true
			if signalCh != nil {
				signalCh <- lastOutput
			}
			//fmt.Println("Halt")
		} else {
			fmt.Printf("Unknown opcode %d\n", op.Code)
			vm.Halted = true
		}
	}
}

func main() {
	buf, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	words := strings.Split(string(buf), ",")
	nums := []int{}
	for _, word := range words {
		num, err := strconv.Atoi(strings.Trim(word, "\n"))
		if err != nil {
			log.Fatal(err)
		}
		nums = append(nums, num)
	}

	fmt.Println(FindMaxPhaseSetting(nums))

}

func FindMaxPhaseSetting(nums []int) (int, []int) {
	A := NewVm(nums)
	B := NewVm(nums)
	C := NewVm(nums)
	D := NewVm(nums)
	E := NewVm(nums)

	maxSignal := 0
	phases := []int{5, 6, 7, 8, 9}
	maxPhaseSetting := make([]int, len(phases))
	for i := int64(0); i < 6*6*6*6*6*6*6; i++ {
		// Didn't bother to loop through these properly
		rand.Shuffle(len(phases), func(i, j int) {
			phases[i], phases[j] = phases[j], phases[i]
		})

		// Wiring
		eaCh := make(chan int, 2)
		eaCh <- phases[0]
		eaCh <- 0 // Initial input
		abCh := make(chan int, 1)
		abCh <- phases[1]
		bcCh := make(chan int, 1)
		bcCh <- phases[2]
		cdCh := make(chan int, 1)
		cdCh <- phases[3]
		deCh := make(chan int, 1)
		deCh <- phases[4]
		signalCh := make(chan int, 0)

		go A.Run(eaCh, abCh, nil)
		go B.Run(abCh, bcCh, nil)
		go C.Run(bcCh, cdCh, nil)
		go D.Run(cdCh, deCh, nil)
		go E.Run(deCh, eaCh, signalCh)

		signal := <-signalCh

		if signal > maxSignal {
			copy(maxPhaseSetting, phases)
			maxSignal = signal
			fmt.Println(maxSignal, maxPhaseSetting)
		}
	}
	return maxSignal, maxPhaseSetting
}
