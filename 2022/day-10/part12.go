package main

import (
	"bufio"
	"embed"
	_ "embed"
	"fmt"
	"log"
	"strconv"
	"strings"
)

//go:embed input.txt input-ex*.txt
var inputFiles embed.FS

func main() {
	part1() // 1:04 argh. Made stuff execute in parallel...
	part2() //
}

func part2() {
	instructions := ReadInput()
	// fmt.Println(instructions)

	vm := Vm{
		IP:                  0,
		X:                   1,
		CurrInstr:           nil,
		CurrInstrCyclesLeft: 0,
	}

	cycle := 1
	fmt.Print("#")
	for vm.CurrInstr != nil || vm.IP < len(instructions) {
		if vm.CurrInstr == nil && vm.IP < len(instructions) {
			// No instruction, more instructions to read.
			// Read instruction and start processing it.
			instr := instructions[vm.IP]
			vm.CurrInstr = &instr
			vm.CurrInstrCyclesLeft = instr.Duration
			vm.IP++
		}

		vm.CurrInstrCyclesLeft--

		// Done processing instruction, add effects to register and
		// remove current instruction.
		if vm.CurrInstr != nil && vm.CurrInstrCyclesLeft == 0 {
			switch vm.CurrInstr.Type {
			case InstrTypeAddx:
				vm.X += vm.CurrInstr.Val
			}
			vm.CurrInstr = nil
		}
		if cycle%40 == vm.X || cycle%40 == vm.X-1 || cycle%40 == vm.X+1 {
			fmt.Print("#")
		} else {
			fmt.Print(".")
		}

		if (cycle+1)%40 == 0 {
			fmt.Println("")
		}
		cycle++
	}
	// fmt.Println(cycle, vm.X)
	// fmt.Println(ans1)
}

func part1() {
	instructions := ReadInput()
	// fmt.Println(instructions)

	vm := Vm{
		IP:                  0,
		X:                   1,
		CurrInstr:           nil,
		CurrInstrCyclesLeft: 0,
	}

	ans1 := 0
	cycle := 1
	for vm.CurrInstr != nil || vm.IP < len(instructions) {
		if vm.CurrInstr == nil && vm.IP < len(instructions) {
			// No instruction, more instructions to read.
			// Read instruction and start processing it.
			instr := instructions[vm.IP]
			vm.CurrInstr = &instr
			vm.CurrInstrCyclesLeft = instr.Duration
			vm.IP++
		}

		if cycle == 20 || cycle == 60 || cycle == 100 || cycle == 140 || cycle == 180 || cycle == 220 {
			ans1 += vm.X * cycle
			// fmt.Println(cycle, vm.X, vm.X*cycle)
		}

		vm.CurrInstrCyclesLeft--

		// Done processing instruction, add effects to register and
		// remove current instruction.
		if vm.CurrInstr != nil && vm.CurrInstrCyclesLeft == 0 {
			switch vm.CurrInstr.Type {
			case InstrTypeAddx:
				vm.X += vm.CurrInstr.Val
			}
			vm.CurrInstr = nil
		}

		cycle++
	}
	// fmt.Println(cycle, vm.X)
	fmt.Println(ans1)
}

func nextQueue(queue map[int][]Instr) ([]Instr, map[int][]Instr) {
	next := make(map[int][]Instr)
	for delay, instrs := range queue {
		if delay > 1 {
			next[delay-1] = instrs
		}
	}
	return queue[1], next
}

type Vm struct {
	IP                  int
	Queue               map[int][]Instr
	X                   int
	CurrInstr           *Instr
	CurrInstrCyclesLeft int
}

type Instr struct {
	Type     InstrType
	Val      int
	Duration int
}

type InstrType string

const (
	InstrTypeNoop InstrType = "noop"
	InstrTypeAddx InstrType = "addx"
)

func ReadInput() []Instr {
	f, err := inputFiles.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var instrs []Instr
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		switch InstrType(parts[0]) {
		case InstrTypeNoop:
			instrs = append(instrs, Instr{
				Type:     InstrTypeNoop,
				Duration: 1,
			})
		case InstrTypeAddx:
			num, err := strconv.Atoi(parts[1])
			if err != nil {
				log.Fatal(err)
			}
			instrs = append(instrs, Instr{
				Type:     InstrTypeAddx,
				Val:      num,
				Duration: 2,
			})
		}

	}
	return instrs
}
