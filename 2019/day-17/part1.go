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

type Tile rune

const TileEmpty Tile = '.'
const TileScaffold Tile = '#'

type Direction struct {
	Code   int
	Name   rune
	Di, Dj int
}

var North Direction = Direction{1, 'N', -1, 0}
var South Direction = Direction{2, 'S', 1, 0}
var West Direction = Direction{3, 'W', 0, -1}
var East Direction = Direction{4, 'E', 0, 1}
var Directions []Direction = []Direction{
	North, South, West, East,
}

func DirectionFromDiDj(di, dj int) Direction {
	if di == North.Di && dj == North.Dj {
		return North
	}
	if di == South.Di && dj == South.Dj {
		return South
	}
	if di == West.Di && dj == West.Dj {
		return West
	}
	if di == East.Di && dj == East.Dj {
		return East
	}
	panic("fuck")
}

type State struct {
	PrintingMap bool
	DroidStartI int
	DroidStartJ int
	DroidI      int
	DroidJ      int
	DroidDir    Direction
	Map         [][]Tile
}

func NewState() *State {
	s := State{}
	s.PrintingMap = true
	s.DroidStartI = -1
	s.DroidStartJ = -1
	s.DroidI = -1
	s.DroidJ = -1
	s.Map = make([][]Tile, 1)
	return &s
}


func (s *State) PrintMap() {
	fmt.Printf("---\n")
	for _, row := range s.Map {
		for _, tile := range row {
			fmt.Printf("%c", tile)
		}
		fmt.Printf("\n")
	}
	fmt.Println("---")
}

func (s *State) MapCopy() [][]Tile {
	c := make([][]Tile, len(s.Map))
	for i := range c {
		c[i] = make([]Tile, len(s.Map[0]))
		copy(c[i], s.Map[i])
	}
	return c
}

func Part1() *State {
	vm := NewVm(LoadFile("input.txt"))
	s := NewState()

	getInput := func() int {
		panic("Should never get input")
	}

	i, j := 0, 0
	sendOutput := func(c int) {
		if s.PrintingMap {

			if c == '\n' {
				j = 0
				i++
				s.Map = append(s.Map, nil)
				return
			}

			switch c {
			case '^':
				s.DroidDir = North
			case 'v':
				s.DroidDir = South
			case '<':
				s.DroidDir = West
			case '>':
				s.DroidDir = East
			}

			if c == '^' || c == 'v' || c == '<' || c == '>' {
				s.DroidStartI = i
				s.DroidStartJ = j
				s.DroidI = i
				s.DroidJ = j
				c = '#'
			}

			s.Map[len(s.Map)-1] = append(s.Map[len(s.Map)-1], Tile(c))
			j++
		}
	}

	vm.Run(getInput, sendOutput)

	s.PrintMap()

	sum := 0
	for i := range s.Map {
		for j, tile := range s.Map[i] {
			if tile != TileScaffold || i == 0 || i == len(s.Map)-1 || j == 0 || j == len(s.Map[i])-1 {
				continue
			}
			if s.Map[i+East.Di][j+East.Dj] != TileScaffold {
				continue
			}
			if s.Map[i+West.Di][j+West.Dj] != TileScaffold {
				continue
			}
			if s.Map[i+North.Di][j+North.Dj] != TileScaffold {
				continue
			}
			if s.Map[i+South.Di][j+South.Dj] != TileScaffold {
				continue
			}
			sum += i * j
		}
	}
}


func main() {

	s := Part1

	vm := NewVm(LoadFile("input.txt"))

	getInput := func() int {
		panic("Should never get input")
	}

	i, j := 0, 0
	sendOutput := func(c int) {
		if s.PrintingMap {

			if c == '\n' {
				j = 0
				i++
				s.Map = append(s.Map, nil)
				return
			}

			switch c {
			case '^':
				s.DroidDir = North
			case 'v':
				s.DroidDir = South
			case '<':
				s.DroidDir = West
			case '>':
				s.DroidDir = East
			}

			if c == '^' || c == 'v' || c == '<' || c == '>' {
				s.DroidStartI = i
				s.DroidStartJ = j
				s.DroidI = i
				s.DroidJ = j
				c = '#'
			}

			s.Map[len(s.Map)-1] = append(s.Map[len(s.Map)-1], Tile(c))
			j++
		}
	}

	vm.Run(getInput, sendOutput)
	s.PrintMap()
	fmt.Println(s.Part1())
}
