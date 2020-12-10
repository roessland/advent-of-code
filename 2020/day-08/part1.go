package main1

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type Vm struct {
	Mem []Inst
	Acc int
	Hits []int
	Ip int
}

type Inst struct {
	Op Op
	Arg int
}

func NewInst(line string) Inst {
	parts := strings.Split(line, " ")
	op := NewOp(parts[0])
	arg, err := strconv.Atoi(parts[1])
	if err != nil {
		log.Fatal(err)
	}
	return Inst{Op: op, Arg: arg}
}

func NewVm(file io.Reader) *Vm {
	var vm Vm
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		vm.Mem = append(vm.Mem, NewInst(line))
	}
	vm.Hits = make([]int, len(vm.Mem))
	return &vm
}

func NewVmFromFile(file string) *Vm {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	return NewVm(f)
}

func (vm *Vm) Run() int {
	for {
		inst := vm.Mem[vm.Ip]
		vm.Hits[vm.Ip]++
		if vm.Hits[vm.Ip] == 2 {
			return vm.Acc
		}
		switch inst.Op {
		case OpAcc:
			vm.Acc += inst.Arg
			vm.Ip += 1
		case OpJmp:
			vm.Ip += inst.Arg
		case OpNop:
			vm.Ip += 1
		}

	}
}


type Op int

const (
	OpAcc Op = iota
	OpJmp
	OpNop
)

func NewOp(name string) Op {
	switch name {
	case "acc":
		return OpAcc
	case "jmp":
		return OpJmp
	case "nop":
		return OpNop
	default:
		panic("Unknown op name " + name)
	}
}

func main() {
	vm := NewVmFromFile("input.txt")
	fmt.Println(vm.Run())
}