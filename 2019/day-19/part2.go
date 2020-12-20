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
	Paused bool
	RelativeBase int
}

func NewVm(mem []int) *Vm {
	vm := Vm{make([]int, 100*len(mem)), 0, false, false, 0}
	copy(vm.Mem, mem)
	return &vm
}

func (vm *Vm) Clone() *Vm {
	clone := &Vm{}
	clone.Mem = make([]int, len(vm.Mem))
	copy(clone.Mem, vm.Mem)
	clone.Ip = vm.Ip
	clone.Halted = vm.Halted
	clone.Paused = vm.Paused
	clone.RelativeBase = vm.RelativeBase
	return clone
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
	vm.Paused = false
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
			input := getInput()
			if vm.Paused {
				return
			}
			vm.SetVal(op.Params[0].Val, op.Params[0].Mode, input)
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

type Tile rune

const TileStationary Tile = '.'
const TileAffected Tile = '#'
const TileUnknown Tile = ' '

type State struct {
	DroneI int
	DroneJ int
	Map    [][]Tile
	I0, I1, J0, J1 int
}

const N = 1100

func NewState() *State {
	s := State{}
	mapHeight := N
	mapWidth := 2*N
	s.Map = make([][]Tile, mapHeight)
	for i := range s.Map {
		s.Map[i] = make([]Tile, mapWidth)
		for j := range s.Map[i] {
			s.Map[i][j] = TileUnknown
		}
	}
	return &s
}

func (s *State) Update(status, i, j int) {
	switch status {
	case 0: // stationary
		s.Map[i][j] = TileStationary
	case 1: // affected
		s.Map[i][j] = TileAffected
	default:
		panic("fuck")
	}
}

func (s *State) PrintMap() {
	fmt.Printf("---\n")
	for i, row := range s.Map {
		for j, tile := range row {
			if s.I0 <= i && i <= s.I1 && s.J0 <= j && j <= s.J1 {
				fmt.Printf("x")
			} else {
				fmt.Printf("%c", tile)
			}

		}
		fmt.Printf("\n")
	}
	fmt.Println("---")
}

func (s *State) PrintBox() {
	fmt.Printf("---\n")
	for i := s.I0-20; i < s.I1+2; i++ {
		for j := s.J0-20; j < s.J1+20; j++ {
			if s.I0 <= i && i <= s.I1 && s.J0 <= j && j <= s.J1 {
				fmt.Printf("x")
			} else {
				fmt.Printf("%c", s.Map[i][j])
			}

		}
		fmt.Printf("\n")
	}
	fmt.Println("---")
}

func main() {

	// Make a VM cached right before input
	cachedVm := NewVm(LoadFile("input.txt"))
	cachedVm.Run(func() int {
		cachedVm.Paused = true
		return 0
	}, nil)

	// Create state
	state := NewState()

	// Send drones
	d := 0
	for i := 5; i < len(state.Map); i++ {
		for j := 5; j < len(state.Map[0]); j++ {
			if float64(i)/float64(j) > 0.62  {
				continue
			}
			if float64(i)/float64(j) < 0.48  {
				continue
			}
			vm := cachedVm.Clone()
			coords := make(chan int, 2)
			// These should be swapped but not bothering to fix.
			coords <- i
			coords <- j
			vm.Run(func() int {
				return <-coords
			}, func(status int) {
				state.Update(status, i, j)
			})
		}
		// See if there is room for santa at some I
		i1 := i
		j0 := 0
		for state.Map[i1][j0] != TileAffected {
			j0++
		}
		d = 0
		i0 := i1
		j1 := j0
		for state.Map[i0-1][j1+1] == TileAffected {
			i0--
			j1++
			d++
		}
		state.I0 = i0
		state.I1 = i1
		state.J0 = j0
		state.J1 = j1
		fmt.Println(d)
		if d == 100-1 {
			break
		}
	}
	if d != 100-1 {
		log.Fatal("increase N")
	}

	//state.PrintBox()
	fmt.Println(state.I0*10000 + state.J0)
}
