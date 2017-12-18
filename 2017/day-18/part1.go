package main

import "fmt"
import "log"
import "strconv"
import "unicode"
import "strings"
import "os"
import "bufio"

func Atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return n
}

type Op struct {
	Name, X, Y string
}

type Vm struct {
	Ip         int
	Ops        []Op
	Reg        map[string]int
	Terminated bool
	LastFreq   int
}

func (vm *Vm) Run() {
	for !vm.Terminated && 0 <= vm.Ip && vm.Ip < len(vm.Ops) {
		op := vm.Ops[vm.Ip]
		vm.Do(op)
	}
}

func (vm *Vm) GetRegOrParseInt(Y string) int {
	if unicode.IsLetter(rune(Y[0])) {
		return vm.Reg[Y]
	} else {
		return Atoi(Y)
	}
}

func (vm *Vm) Do(op Op) {
	fmt.Println(op, vm.Reg, vm.LastFreq)
	switch op.Name {
	case "snd":
		vm.LastFreq = vm.Reg[op.X]
		fmt.Println("Playing", vm.LastFreq)
		vm.Ip++
	case "set":
		vm.Reg[op.X] = vm.GetRegOrParseInt(op.Y)
		vm.Ip++
	case "add":
		vm.Reg[op.X] += vm.GetRegOrParseInt(op.Y)
		vm.Ip++
	case "mul":
		vm.Reg[op.X] *= vm.GetRegOrParseInt(op.Y)
		vm.Ip++
	case "mod":
		vm.Reg[op.X] %= vm.GetRegOrParseInt(op.Y)
		vm.Ip++
	case "rcv":
		if vm.Reg[op.X] != 0 {
			fmt.Println("Recovered frequency", vm.LastFreq)
			vm.Terminated = true
		}
		vm.Ip++
	case "jgz":
		if vm.Reg[op.X] > 0 {
			vm.Ip += vm.GetRegOrParseInt(op.Y)
		} else {
			vm.Ip++
		}
	default:
		log.Fatal("Unknown instruction", op.Name)
	}
}

func ParseInput() *Vm {
	vm := Vm{}
	vm.Reg = make(map[string]int)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), " ")
		if len(fields) == 2 {
			vm.Ops = append(vm.Ops, Op{fields[0], fields[1], ""})
		} else if len(fields) == 3 {
			vm.Ops = append(vm.Ops, Op{fields[0], fields[1], fields[2]})
		}
	}
	return &vm
}

func main() {
	vm := ParseInput()
	vm.Run()
	fmt.Println("Terminated.")
}
