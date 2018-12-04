package main

import "fmt"
import "os"
import "strconv"
import "bufio"
import "strings"

type Op struct {
	Name string
	A    string
	B    string
}

type Vm struct {
	Ip    int
	Mem   []Op
	Reg   map[string]int
	Stats map[string]int
}

func NewVm() *Vm {
	vm := Vm{}
	vm.Ip = 0
	vm.Mem = []Op{}
	vm.Reg = make(map[string]int)
	vm.Stats = make(map[string]int)
	return &vm
}

func (vm *Vm) Get(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		return vm.Reg[s]
	}
	return n
}

func (vm *Vm) Run() {
	for {
		if vm.Ip < 0 || vm.Ip >= len(vm.Mem) {
			fmt.Println("Terminated.")
			break
		}
		op := vm.Mem[vm.Ip]
		switch op.Name {
		case "set":
			vm.Reg[op.A] = vm.Get(op.B)
			vm.Ip++
		case "sub":
			vm.Reg[op.A] -= vm.Get(op.B)
			vm.Ip++
		case "mul":
			vm.Stats["mul"]++
			vm.Reg[op.A] *= vm.Get(op.B)
			vm.Ip++
		case "jnz":
			X := vm.Get(op.A)
			if X != 0 {
				vm.Ip += vm.Get(op.B)
			} else {
				vm.Ip++
			}
		}
	}
}

func main() {
	ops := []Op{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), " ")
		ops = append(ops, Op{fields[0], fields[1], fields[2]})
	}
	vm := NewVm()
	vm.Mem = ops
	vm.Run()
	fmt.Println(vm.Stats["mul"])
}
