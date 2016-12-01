package main

import "io/ioutil"
import "fmt"
import "strings"
import "strconv"

type Instruction struct {
	Name     string
	Register string
	Offset   int
}

type VM struct {
	Register map[string]int
	Memory   []Instruction
	IP       int
}

func NewVM() *VM {
	return &VM{
		Register: map[string]int{},
		Memory:   []Instruction{},
		IP:       0,
	}
}

func (vm *VM) Execute() {
	for {
		if vm.IP >= len(vm.Memory) {
			return
		}
		instr := vm.Memory[vm.IP]
		switch instr.Name {
		case "hlf":
			vm.Register[instr.Register] /= 2
			vm.IP++
		case "tpl":
			vm.Register[instr.Register] *= 3
			vm.IP++
		case "inc":
			vm.Register[instr.Register]++
			vm.IP++
		case "jmp":
			vm.IP += instr.Offset
		case "jie":
			if vm.Register[instr.Register]%2 == 0 {
				vm.IP += instr.Offset
			} else {
				vm.IP++
			}
		case "jio":
			if vm.Register[instr.Register] == 1 {
				vm.IP += instr.Offset
			} else {
				vm.IP++
			}
		default:
			panic("Unknown instruction")
		}
	}
}

func LoadProgram() *VM {
	vm := NewVM()
	buf, _ := ioutil.ReadFile("input.txt")
	for _, line := range strings.Split(string(buf), "\n") {
		if len(line) == 0 {
			break
		}
		words := strings.Split(line, " ")
		instr := Instruction{Name: words[0]}
		switch instr.Name {
		case "hlf":
			instr.Register = words[1]
		case "tpl":
			instr.Register = words[1]
		case "inc":
			instr.Register = words[1]
		case "jmp":
			instr.Offset, _ = strconv.Atoi(words[1])
		case "jie":
			instr.Register = strings.TrimRight(words[1], ",")
			instr.Offset, _ = strconv.Atoi(words[2])
		case "jio":
			instr.Register = strings.TrimRight(words[1], ",")
			instr.Offset, _ = strconv.Atoi(words[2])
		default:
			panic("Unknown instruction")
		}
		vm.Memory = append(vm.Memory, instr)
	}
	return vm
}

func main() {
	vm1 := LoadProgram()
	vm1.Execute()
	fmt.Printf("Part 1: %v\n", vm1.Register)

	vm2 := LoadProgram()
	vm2.Register["a"] = 1
	vm2.Execute()
	fmt.Printf("Part 2: %v\n", vm2.Register)

}
