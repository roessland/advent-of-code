package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type Vm struct {
	Mem          []int
	Ip           int
	Halted       bool
	RelativeBase int
}

func NewVm(mem []int) *Vm {
	vm := Vm{make([]int, 100*len(mem)), 0, false, 0}
	copy(vm.Mem, mem)
	return &vm
}

const (
	PositionMode  int = 0
	ImmediateMode int = 1
	RelativeMode  int = 2
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
	case 9:
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
	} else if mode == RelativeMode {
		return vm.Mem[vm.RelativeBase+val]
	} else {
		log.Fatal("GetVal: unknown position mode", mode)
		return -1337
	}
}

func (vm *Vm) SetVal(pos int, mode int, val int) {
	if mode == PositionMode {
		vm.Mem[pos] = val
	} else if mode == ImmediateMode {
		panic("Impossible to set a value in immediate mode!")
	} else if mode == RelativeMode {
		vm.Mem[vm.RelativeBase+pos] = val
	} else {
		log.Fatal("GetVal: unknown position mode", mode)
	}
}

func (vm *Vm) Run(getInput func() int, sendOutput func(int)) {
	for !vm.Halted {
		//fmt.Println("MEM", vm.Mem)
		//fmt.Println("IP", vm.Ip)
		op := vm.ReadOp()
		//fmt.Printf("OP: %#v\n", op)
		if op.Code == 1 {
			// Add
			val := vm.GetVal(op.Params[0].Val, op.Params[0].Mode) + vm.GetVal(op.Params[1].Val, op.Params[1].Mode)
			vm.SetVal(op.Params[2].Val, op.Params[2].Mode, val)
			vm.Ip += op.Length
		} else if op.Code == 2 {
			// Multiply
			val := vm.GetVal(op.Params[0].Val, op.Params[0].Mode) * vm.GetVal(op.Params[1].Val, op.Params[1].Mode)
			vm.SetVal(op.Params[2].Val, op.Params[2].Mode, val)
			vm.Ip += op.Length
		} else if op.Code == 3 {
			// Input
			vm.SetVal(op.Params[0].Val, op.Params[0].Mode, getInput())
			// fmt.Println("Provided input", input)
			vm.Ip += op.Length
		} else if op.Code == 4 {
			// Output
			sendOutput(vm.GetVal(op.Params[0].Val, op.Params[0].Mode))
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
				vm.SetVal(op.Params[2].Val, op.Params[2].Mode, 1)
			} else {
				vm.SetVal(op.Params[2].Val, op.Params[2].Mode, 0)
			}
			vm.Ip += op.Length
		} else if op.Code == 8 {
			// equals
			if vm.GetVal(op.Params[0].Val, op.Params[0].Mode) == vm.GetVal(op.Params[1].Val, op.Params[1].Mode) {
				vm.SetVal(op.Params[2].Val, op.Params[2].Mode, 1)

			} else {
				vm.SetVal(op.Params[2].Val, op.Params[2].Mode, 0)
			}
			vm.Ip += op.Length
		} else if op.Code == 9 {
			// adjust relative base
			vm.RelativeBase += vm.GetVal(op.Params[0].Val, op.Params[0].Mode)
			vm.Ip += op.Length
		} else if op.Code == 99 {
			// Halt
			vm.Halted = true
			//fmt.Println("Halt")
		} else {
			fmt.Printf("Unknown opcode %d\n", op.Code)
			vm.Halted = true
		}
	}
}

func LoadString(numsStr string) []int {
	words := strings.Split(numsStr, ",")
	nums := []int{}
	for _, word := range words {
		num, err := strconv.Atoi(strings.Trim(word, "\n"))
		if err != nil {
			log.Fatal(err)
		}
		nums = append(nums, num)
	}
	return nums
}

func LoadFile(fileName string) []int {
	buf, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	return LoadString(string(buf))
}

type Pos struct {
	I, J int
}

type Robot struct {
	Vm          *Vm
	Dir         string
	Pos         Pos
	OutputState int
	Colors map[Pos]int
}

func (robot *Robot) TurnLeft() {
	switch robot.Dir {
	case "^":
		robot.Dir = "<"
	case "<":
		robot.Dir = "v"
	case "v":
		robot.Dir = ">"
	case ">":
		robot.Dir = "^"
	default:
		panic("unknown dir")
	}
}

func (robot *Robot) TurnRight() {
	switch robot.Dir {
	case "^":
		robot.Dir = ">"
	case "<":
		robot.Dir = "^"
	case "v":
		robot.Dir = "<"
	case ">":
		robot.Dir = "v"
	default:
		panic("unknown dir")
	}
}

func (robot *Robot) Turn(code int) {
	switch code {
	case 0:
		robot.TurnLeft()
	case 1:
		robot.TurnRight()
	default:
		panic("unknown turncode")
	}
}

func (robot *Robot) Drive() {
	switch robot.Dir {
	case "^":
		robot.Pos.I--
	case "<":
		robot.Pos.J--
	case "v":
		robot.Pos.I++
	case ">":
		robot.Pos.J++
	default:
		panic("unknown drive dir")
	}
}

func (robot *Robot) Run(startColor int) int {
	robot.Colors = map[Pos]int{}
	robot.Colors[Pos{0,0}] = startColor

	onVmInput := func() int {
		return robot.Colors[robot.Pos]
	}

	onVmOutput := func(c int) {
		switch robot.OutputState {
		case 0:
			// First, it will output a value indicating the color to paint the panel the robot is over:
			// 0 means to paint the panel black, and 1 means to paint the panel white.
			robot.Colors[robot.Pos] = c
		case 1:
			// Second, it will output a value indicating the direction the robot should turn:
			// 0 means it should turn left 90 degrees, and 1 means it should turn right 90 degrees.
			robot.Turn(c)
			robot.Drive()
		}
		robot.OutputState = (robot.OutputState + 1) % 2
	}

	robot.Vm.Run(onVmInput, onVmOutput)
	return len(robot.Colors)
}

func (robot *Robot) PrintPanels() {
	for i := 0; i < 7; i++ {
		for j := 0; j < 80; j++ {
			switch robot.Colors[Pos{i,j}] {
			case 0:
				fmt.Printf(" ")
			case 1:
				fmt.Printf("█")
			default:
				panic("unknown color")
			}
		}
		fmt.Println()
	}
}

func Part1() {
	robot := &Robot{}
	robot.Dir = "^"
	robot.Vm = NewVm(LoadFile("input.txt"))
	panelsColored := robot.Run(0)
	fmt.Println("Part 1:", panelsColored)
}

func Part2() {
	robot := &Robot{}
	robot.Dir = "^"
	robot.Vm = NewVm(LoadFile("input.txt"))
	robot.Run(1)
	fmt.Println("Part 2:")
	robot.PrintPanels()
}

func main() {
	Part1()
	Part2()
}
