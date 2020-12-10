package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)


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

type Vm struct {
	Mem []Inst
	Acc int
	Hits []int
	Ip int
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

func (vm *Vm) Clone() *Vm {
	clone := Vm{
		Mem:  make([]Inst, len(vm.Mem)),
		Acc:  vm.Acc,
		Hits: make([]int, len(vm.Hits)),
		Ip:   vm.Ip,
	}
	copy(clone.Mem, vm.Mem)
	copy(clone.Hits, vm.Hits)
	return &clone
}

func (vm *Vm) Run(ctx context.Context) (int, error) {
	for {
		if vm.Ip == len(vm.Mem) {
			return vm.Acc, nil
		}
		inst := vm.Mem[vm.Ip]
		vm.Hits[vm.Ip]++
		switch inst.Op {
		case OpAcc:
			vm.Acc += inst.Arg
			vm.Ip += 1
		case OpJmp:
			vm.Ip += inst.Arg
		case OpNop:
			vm.Ip += 1
		}

		select {
		case <-ctx.Done():
			return vm.Acc, ctx.Err()
		default:
			continue
		}
	}
}



func main() {
	vm0 := NewVmFromFile("input.txt")

	for i := 0; i < len(vm0.Mem); i++ {
		vm := vm0.Clone()
		inst := vm.Mem[i]
		if inst.Op == OpJmp {
			inst.Op = OpNop
		} else if inst.Op == OpNop {
			inst.Op = OpJmp
		} else {
			continue
		}
		vm.Mem[i] = inst
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Millisecond)
		acc, err := vm.Run(ctx)
		if err == nil {
			fmt.Println(acc)
		}
	}

}