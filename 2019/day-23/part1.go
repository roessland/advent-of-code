package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
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

type Packet struct {
	Dst, X, Y int
}

type Router struct {
	Mux chan Packet
	Machines map[int]chan<- Packet
}

func NewRouter() *Router {
	return &Router{
		Mux: make(chan Packet, 2000000),
		Machines: make(map[int]chan<- Packet),
	}
}

func (r *Router) Register(machine int, input chan<-Packet, output <-chan Packet) {
	// Register machine input channel
	r.Machines[machine] = input

	// Forward all packets from this machine to router
	go func() {
		for {
			r.Mux<- <-output
		}
	}()
}

func (r *Router) Run() {
	for {
		// Get a packet
		p := <-r.Mux
		if p.Dst == 255 {
			fmt.Println("Part 1:", p.Y)
			os.Exit(0)
		}
		dstMachine, ok := r.Machines[p.Dst]
		if !ok {
			fmt.Println(p)
			panic("Got packet to unknown machine")
		}
		// Forward contents to correct machine
		dstMachine <- p
	}

}

func main() {
	baseVm := NewVm(LoadFile("input.txt"))

	router := NewRouter()

	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(1)

		// Networking
		packetsIn := make(chan Packet)
		packetsOut := make(chan Packet)
		router.Register(i, packetsIn, packetsOut)

		// I/O
		output := make(chan int)

		// Packets in to input, infinitely buffered.
		var queue []Packet
		var queueMutex sync.Mutex
		go func(i int) {
			for {
				p := <-packetsIn
				queueMutex.Lock()
				queue = append(queue, p)
				queueMutex.Unlock()
			}
		}(i)

		// Packets out to output
		go func(i int) {
			for {
				dst, x, y := <-output, <-output, <-output
				packetsOut <- Packet{dst, x, y}
			}
		}(i)

		// Run computer
		go func(i int) {
			hasAddr := false
			hasX := false
			vm := baseVm.Clone()
			vm.Run(func() int {
				// First input, just the address
				if !hasAddr {
					hasAddr = true
					return i
				}

				// Subsequent inputs, send packet data
				queueMutex.Lock()
				defer queueMutex.Unlock()

				// Empty packet queue, send -1
				if len(queue) == 0 {
					return -1
				}

				// Else, send X or Y from the available packet.
				if !hasX {
					// Send X of first packet
					x := queue[0].X
					hasX = true
					return x
				} else {
					// Send Y of first packet, then delete packet from queue
					hasX = false
					y := queue[0].Y
					queue = queue[1:]
					return y
				}
			}, func(outputVal int) {
				output <- outputVal
			})
			wg.Done()
		}(i)
	}
	go router.Run()
	wg.Wait()
}
