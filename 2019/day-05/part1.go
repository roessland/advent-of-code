package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type Vm struct {
	Mem    []int
	Ip     int
	Halted bool
}

func NewVm(mem []int) *Vm {
	vm := Vm{make([]int, len(mem)), 0, false}
	copy(vm.Mem, mem)
	return &vm
}

type ParameterMode int

const (
	PositionMode int = 0
	ImmediateMode int = 1
)

type Param struct {
	Val int
	Mode int
}

type Op struct {
	Code int
	Params []Param
	Length int
}

func (vm *Vm) ReadOp() Op {
	op := Op{}
	val := vm.Mem[vm.Ip]
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

func (vm *Vm) Run() int {
	for !vm.Halted {
		opCode := vm.Mem[vm.Ip]
		op := vm.ReadOp()
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
			vm.Mem[op.Params[0].Val] = 1
			fmt.Println("Provided input 1")
			vm.Ip += op.Length
		} else if op.Code == 4 {
			// Output
			fmt.Println("OUTPUT: ", vm.GetVal(op.Params[0].Val, op.Params[0].Mode), strconv.Itoa(vm.GetVal(op.Params[0].Val, op.Params[0].Mode)))
			vm.Ip += op.Length
		} else if op.Code == 99 {
			// Halt
			vm.Halted = true
			fmt.Println("Halt")
		} else {
			fmt.Printf("Unknown opcode %d\n", opCode)
			vm.Halted = true
		}
	}
	return vm.Mem[0]
}

func main() {
	buf, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	words := strings.Split(string(buf), ",")
	nums := []int{}
	for _, word := range words {
		num, _ := strconv.Atoi(word)
		nums = append(nums, num)
	}

	vm := NewVm(nums)
	vm.Run()
}
