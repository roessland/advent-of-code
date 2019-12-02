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

func (vm *Vm) Run() int {
	for !vm.Halted {
		opCode := vm.Mem[vm.Ip]
		if opCode == 1 {
			srcA := vm.Mem[vm.Ip+1]
			srcB := vm.Mem[vm.Ip+2]
			dst := vm.Mem[vm.Ip+3]
			vm.Mem[dst] = vm.Mem[srcA] + vm.Mem[srcB]
			vm.Ip += 4
		} else if opCode == 2 {
			srcA := vm.Mem[vm.Ip+1]
			srcB := vm.Mem[vm.Ip+2]
			dst := vm.Mem[vm.Ip+3]
			vm.Mem[dst] = vm.Mem[srcA] * vm.Mem[srcB]
			vm.Ip += 4
		} else if opCode == 99 {
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

	for in1 := 0; in1 <= 99; in1++ {
		for in2 := 0; in2 <= 99; in2++ {
			vm := NewVm(nums)
			vm.Mem[1] = in1
			vm.Mem[2] = in2
			output := vm.Run()
			if output == 19690720 {
				fmt.Println(100*in1 + in2)
				log.Fatal("Success")
			}
		}
	}

}
