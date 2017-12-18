package main

import "fmt"
import "log"
import "sync"
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

type Coordinator struct {
	NumSends int
	NumRecvs int
	sync.Mutex
}

type Vm struct {
	Ip         int
	Ops        []Op
	Reg        map[string]int
	LastFreq   int
	Done       chan (bool)
	Snd        chan (int)
	Rcv        chan (int)
	SentValues int
	Name       int
}

func (vm *Vm) Clone() *Vm {
	cpy := Vm{}
	cpy.Ip = 0
	cpy.Ops = make([]Op, len(vm.Ops))
	cpy.Reg = make(map[string]int)
	cpy.LastFreq = 0
	copy(cpy.Ops, vm.Ops)
	return &cpy
}

func (vm *Vm) Run() {
	for 0 <= vm.Ip && vm.Ip < len(vm.Ops) {
		op := vm.Ops[vm.Ip]
		vm.Do(op)
	}
	vm.Done <- true
}

func (vm *Vm) GetRegOrParseInt(Y string) int {
	if unicode.IsLetter(rune(Y[0])) {
		return vm.Reg[Y]
	} else {
		return Atoi(Y)
	}
}

func (vm *Vm) Do(op Op) {
	switch op.Name {
	case "snd":
		vm.Snd <- vm.GetRegOrParseInt(op.X)
		vm.SentValues++
		if vm.Name == 1 {
			fmt.Println(vm.SentValues)
		}
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
		vm.Reg[op.X] = <-vm.Rcv
		vm.Ip++
	case "jgz":
		if vm.GetRegOrParseInt(op.X) > 0 {
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

	vm0 := vm.Clone()
	vm0.Done = make(chan (bool))
	vm0.Snd = make(chan (int), 100000)
	vm0.Rcv = make(chan (int), 100000)
	vm0.Name = 0
	go vm0.Run()

	vm1 := vm.Clone()
	vm1.Done = make(chan (bool))
	vm1.Snd = vm0.Rcv
	vm1.Rcv = vm0.Snd
	vm1.Reg["p"] = 1
	vm1.Name = 1
	go vm1.Run()

	<-vm0.Done
	<-vm1.Done
}
